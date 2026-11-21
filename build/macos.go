package build

import (
	"archive/zip"
	"context"
	"io"
	"strings"
	"sync"

	"github.com/prettylittlelies/cry/internal/fetch"
	"howett.net/plist"
)

const electronPlist = "/Contents/Frameworks/Electron Framework.framework/Versions/A/Resources/Info.plist"

type macUpdate struct {
	Name Version `json:"name"`
	URL  string  `json:"url"`
}

type bundleInfo struct {
	Version string `plist:"CFBundleVersion"`
}

type electronRelease struct {
	Version string `json:"version"`
	Chrome  string `json:"chrome"`
}

type releaseIndex func() ([]electronRelease, error)

func (s *Scraper) macOS(ctx context.Context, c Channel, index releaseIndex) (MacOS, error) {
	upd, err := fetch.JSON[macUpdate](ctx, s.client, s.endpoints.macUpdate(c))
	if err != nil {
		return MacOS{}, err
	}
	electron, err := s.electronVersion(ctx, upd.URL)
	if err != nil {
		return MacOS{}, err
	}
	releases, err := index()
	if err != nil {
		return MacOS{}, err
	}
	chrome, err := chromeFor(releases, electron)
	if err != nil {
		return MacOS{}, err
	}
	return MacOS{Host: upd.Name, Electron: electron, Chrome: chrome}, nil
}

func (s *Scraper) electronVersion(ctx context.Context, url string) (Version, error) {
	r, err := fetch.OpenRanged(ctx, s.client, url)
	if err != nil {
		return nil, err
	}
	archive, err := zip.NewReader(r, r.Size())
	if err != nil {
		return nil, err
	}
	for _, f := range archive.File {
		if strings.HasSuffix("/"+f.Name, electronPlist) {
			return readBundleVersion(f)
		}
	}
	return nil, &ExtractError{Source: "macOS package", Field: "Electron Framework Info.plist"}
}

func readBundleVersion(f *zip.File) (Version, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	raw, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	var info bundleInfo
	if _, err := plist.Unmarshal(raw, &info); err != nil {
		return nil, err
	}
	return ParseVersion(info.Version)
}

func (s *Scraper) ChromeFor(ctx context.Context, electron Version) (Version, error) {
	releases, err := s.electronIndex(ctx)()
	if err != nil {
		return nil, err
	}
	return chromeFor(releases, electron)
}

func chromeFor(releases []electronRelease, electron Version) (Version, error) {
	want := electron.String()
	for _, r := range releases {
		if r.Version == want {
			return ParseVersion(r.Chrome)
		}
	}
	return nil, &ExtractError{Source: "electron release index", Field: "electron " + electron.String()}
}

func (s *Scraper) electronIndex(ctx context.Context) releaseIndex {
	return sync.OnceValues(func() ([]electronRelease, error) {
		return fetch.JSON[[]electronRelease](ctx, s.client, s.endpoints.ElectronIndex)
	})
}
