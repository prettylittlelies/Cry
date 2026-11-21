package voice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/prettylittlelies/cry/voice/dave"
)

type MediaType string

const (
	Camera MediaType = "video"
	Screen MediaType = "screen"
)

var (
	ErrNoVideo    = errors.New("voice: video was not negotiated on this connection")
	ErrVideoCodec = errors.New("voice: negotiated video codec is not VP8")
)

type VideoOptions struct {
	Width   int
	Height  int
	FPS     int
	Bitrate int
}

func (o VideoOptions) defaults() VideoOptions {
	if o.FPS <= 0 {
		o.FPS = 30
	}
	if o.Bitrate <= 0 {
		o.Bitrate = 2_500_000
	}
	return o
}

func (v *Conn) SetVideo(active bool, opts VideoOptions) error {
	s := v.video
	ws := v.ws.Load()
	if s == nil || ws == nil {
		return ErrNoVideo
	}
	opts = opts.defaults()
	state := videoState{AudioSSRC: v.ssrc, Streams: []streamInfo{}}
	if active {
		state.VideoSSRC, state.RTXSSRC = s.ssrc, s.rtx
		res := &resolution{Type: "source"}
		if opts.Width > 0 && opts.Height > 0 {
			res = &resolution{Type: "fixed", Width: opts.Width, Height: opts.Height}
		}
		state.Streams = []streamInfo{{
			Type: string(v.opts.Video), RID: primaryRID, Quality: 100, Active: true,
			MaxBitrate: opts.Bitrate, MaxFramerate: opts.FPS, MaxResolution: res,
			SSRC: s.ssrc, RTXSSRC: s.rtx,
		}}
	}
	v.videoFPS.Store(int32(opts.FPS))
	return ws.send(OpVideo, state)
}

func (v *Conn) WriteVP8(frame []byte) error {
	s := v.video
	if s == nil {
		return ErrNoVideo
	}
	if codec := v.videoCodec.Load(); codec != nil && !strings.EqualFold(*codec, "VP8") {
		return fmt.Errorf("%w: %s", ErrVideoCodec, *codec)
	}
	t, u := v.crypt.Load(), v.udp.Load()
	if t == nil || u == nil {
		return ErrNotReady
	}
	payload := frame
	if v.e2ee != nil {
		var err error
		if payload, err = v.e2ee.protectRanges(frame, dave.VP8Ranges(frame)); err != nil {
			return err
		}
	}
	fps := max(v.videoFPS.Load(), 1)
	for _, p := range s.packetize(payload, uint32(videoClock/fps)) {
		if err := u.write(t.seal(p.header, p.payload)); err != nil {
			return err
		}
	}
	return nil
}

func (v *Conn) PlayIVF(ctx context.Context, r io.Reader, opts VideoOptions) error {
	ivf, err := NewIVF(r)
	if err != nil {
		return err
	}
	if !strings.EqualFold(ivf.Header.FourCC, "VP80") {
		return fmt.Errorf("%w: %s", ErrVideoCodec, ivf.Header.FourCC)
	}
	if opts.Width == 0 {
		opts.Width, opts.Height = ivf.Header.Width, ivf.Header.Height
	}
	if opts.FPS == 0 && ivf.Header.FrameRate > 0 {
		opts.FPS = int(ivf.Header.FrameRate + 0.5)
	}
	opts = opts.defaults()
	if err := v.SetVideo(true, opts); err != nil {
		return err
	}
	defer v.SetVideo(false, opts)
	t := time.NewTicker(time.Second / time.Duration(opts.FPS))
	defer t.Stop()
	for {
		frame, err := ivf.Next()
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return nil
		}
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-v.ctx.Done():
			return ErrClosed
		case <-t.C:
		}
		if err := v.WriteVP8(frame); err != nil {
			return err
		}
	}
}
