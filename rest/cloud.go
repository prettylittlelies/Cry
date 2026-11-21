package rest

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/prettylittlelies/cry/internal/backoff"
	"github.com/prettylittlelies/cry/internal/mimetype"
	"github.com/prettylittlelies/cry/snowflake"
)

const (
	uploadURLTTL   = 12 * time.Hour
	uploadRetries  = 12
	uploadDeadline = time.Hour
)

var rangePattern = regexp.MustCompile(`^bytes=(\d+)-(\d+)(?:/\d+)?$`)

type CloudFile struct {
	Name        string
	ContentType string
	Data        []byte
	Description string
	Spoiler     bool
}

func (f CloudFile) mime() string {
	if f.ContentType != "" {
		return f.ContentType
	}
	return mimetype.ForFilename(f.Name)
}

type uploadSlot struct {
	Filename            string `json:"filename"`
	FileSize            int    `json:"file_size"`
	ID                  string `json:"id"`
	IsClip              bool   `json:"is_clip"`
	OriginalContentType string `json:"original_content_type,omitempty"`
}

type CloudAttachment struct {
	ID             string `json:"id"`
	UploadURL      string `json:"upload_url"`
	UploadFilename string `json:"upload_filename"`
}

type MessageAttachment struct {
	ID               string `json:"id"`
	Description      string `json:"description,omitempty"`
	Filename         string `json:"filename"`
	UploadedFilename string `json:"uploaded_filename"`
	IsSpoiler        bool   `json:"is_spoiler,omitempty"`
}

type Upload struct {
	File             CloudFile
	UploadedFilename string
}

func (u Upload) Attachment(index int) MessageAttachment {
	return MessageAttachment{
		ID:               strconv.Itoa(index),
		Description:      u.File.Description,
		Filename:         u.File.Name,
		UploadedFilename: u.UploadedFilename,
		IsSpoiler:        u.File.Spoiler,
	}
}

type UploadError struct {
	Filename string
	Attempts int
	Status   int
	Err      error
}

func (e *UploadError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("rest: upload %s failed after %d attempts: %v", e.Filename, e.Attempts, e.Err)
	}
	return fmt.Sprintf("rest: upload %s failed after %d attempts (status %d)", e.Filename, e.Attempts, e.Status)
}

func (e *UploadError) Unwrap() error {
	return e.Err
}

var ErrEmptyFile = errors.New("rest: cannot upload an empty file")

type uploadIDs struct {
	next atomic.Uint64
}

func newUploadIDs() *uploadIDs {
	ids := new(uploadIDs)
	ids.next.Store(uint64(100 + rand.N(1900)))
	return ids
}

func (u *uploadIDs) take() string {
	return strconv.FormatUint(u.next.Add(1), 10)
}

func (c *Client) requestUploadSlot(ctx context.Context, channel snowflake.ID, f CloudFile) (CloudAttachment, error) {
	body := struct {
		Files []uploadSlot `json:"files"`
	}{Files: []uploadSlot{{
		Filename:            f.Name,
		FileSize:            len(f.Data),
		ID:                  c.uploads.take(),
		OriginalContentType: f.mime(),
	}}}
	out, err := Call[struct {
		Attachments []CloudAttachment `json:"attachments"`
	}](ctx, c, &Request{
		Route: NewRoute(http.MethodPost, "/channels/{channel_id}/attachments", P("channel_id", channel)),
		JSON:  body,
	})
	if err != nil {
		return CloudAttachment{}, err
	}
	if len(out.Attachments) == 0 {
		return CloudAttachment{}, &DecodeError{Route: "POST /channels/{channel_id}/attachments", Err: errors.New("no attachments returned")}
	}
	return out.Attachments[0], nil
}

func (c *Client) UploadAttachment(ctx context.Context, channel snowflake.ID, f CloudFile) (*Upload, error) {
	if len(f.Data) == 0 {
		return nil, ErrEmptyFile
	}
	ctx, cancel := context.WithTimeout(ctx, uploadDeadline)
	defer cancel()
	slot, err := c.requestUploadSlot(ctx, channel, f)
	if err != nil {
		return nil, err
	}
	issued := c.limiter.now()
	wait := backoff.Default()
	var (
		offset   int
		lastErr  error
		lastCode int
	)
	for attempt := 1; attempt <= uploadRetries+1; attempt++ {
		switch {
		case c.limiter.now().Sub(issued) >= uploadURLTTL:
			if slot, err = c.requestUploadSlot(ctx, channel, f); err != nil {
				return nil, err
			}
			issued, offset = c.limiter.now(), 0
		case attempt > 1:
			if offset, err = c.resumePosition(ctx, slot.UploadURL, len(f.Data)); err != nil {
				offset = 0
			}
		}
		if offset >= len(f.Data) {
			return &Upload{File: f, UploadedFilename: slot.UploadFilename}, nil
		}
		resp, err := c.putChunk(ctx, slot.UploadURL, f.Data, offset)
		if err == nil && (resp.Status == http.StatusOK || resp.Status == http.StatusCreated) {
			return &Upload{File: f, UploadedFilename: slot.UploadFilename}, nil
		}
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		lastErr, lastCode = err, 0
		if resp != nil {
			lastCode = resp.Status
		}
		if err := sleep(ctx, wait.Fail(0)); err != nil {
			return nil, err
		}
	}
	return nil, &UploadError{Filename: f.Name, Attempts: uploadRetries + 1, Status: lastCode, Err: lastErr}
}

func (c *Client) putChunk(ctx context.Context, target string, data []byte, offset int) (*Response, error) {
	h := http.Header{}
	h.Set("Content-Type", "application/octet-stream")
	if offset > 0 {
		h.Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", offset, len(data)-1, len(data)))
	}
	return c.External(ctx, &External{Method: http.MethodPut, URL: target, Header: h, Body: data[offset:]})
}

func (c *Client) resumePosition(ctx context.Context, target string, size int) (int, error) {
	h := http.Header{}
	h.Set("Content-Range", "bytes */*")
	resp, err := c.External(ctx, &External{Method: http.MethodPut, URL: target, Header: h})
	if err != nil {
		return 0, err
	}
	switch resp.Status {
	case http.StatusOK, http.StatusCreated:
		return size, nil
	case http.StatusPermanentRedirect:
		m := rangePattern.FindStringSubmatch(resp.Header.Get("Range"))
		if m == nil {
			return 0, nil
		}
		end, _ := strconv.Atoi(m[2])
		return end + 1, nil
	}
	return 0, &UploadError{Status: resp.Status}
}

func (c *Client) DeleteUpload(ctx context.Context, uploadedFilename string) error {
	_, err := c.Do(ctx, &Request{Route: NewRoute(http.MethodDelete, "/attachments/{upload_filename}", Raw("upload_filename", uploadedFilename))})
	return err
}

type RefreshedURL struct {
	Original  string `json:"original"`
	Refreshed string `json:"refreshed"`
}

func (c *Client) RefreshAttachmentURLs(ctx context.Context, urls ...string) ([]RefreshedURL, error) {
	out, err := Call[struct {
		RefreshedURLs []RefreshedURL `json:"refreshed_urls"`
	}](ctx, c, &Request{
		Route: NewRoute(http.MethodPost, "/attachments/refresh-urls"),
		JSON: struct {
			AttachmentURLs []string `json:"attachment_urls"`
		}{urls},
	})
	return out.RefreshedURLs, err
}
