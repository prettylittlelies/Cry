package build

import (
	"context"

	"github.com/prettylittlelies/cry/internal/fetch"
)

type manifest struct {
	Full struct {
		HostVersion Version `json:"host_version"`
	} `json:"full"`
	MetadataVersion *NativeNumber `json:"metadata_version"`
}

func (s *Scraper) windows(ctx context.Context, c Channel) (Windows, error) {
	m, err := fetch.JSON[manifest](ctx, s.client, s.endpoints.windowsManifest(c))
	if err != nil {
		return Windows{}, err
	}
	if m.MetadataVersion == nil {
		return Windows{}, &ExtractError{Source: "windows manifest", Field: "metadata_version"}
	}
	return Windows{Host: m.Full.HostVersion, Native: *m.MetadataVersion}, nil
}
