package build

import (
	"context"
	"regexp"
	"strconv"

	"github.com/prettylittlelies/cry/internal/fetch"
)

var playVersionPattern = regexp.MustCompile(`\[\[\["(\d+\.\d+) - [A-Za-z]+"`)

type itunesLookup struct {
	Results []struct {
		Version          string  `json:"version"`
		MinimumOSVersion Version `json:"minimumOsVersion"`
	} `json:"results"`
}

type androidVersions struct {
	Minimum string `json:"discord_android_min_version"`
}

func (s *Scraper) mobile(ctx context.Context) (Mobile, error) {
	var m Mobile
	page, err := s.client.Bytes(ctx, s.endpoints.PlayStore)
	if err != nil {
		return m, err
	}
	pm := playVersionPattern.FindSubmatch(page)
	if pm == nil {
		return m, &ExtractError{Source: "play store page", Field: "version"}
	}
	if m.Android, err = ParseMobileVersion(string(pm[1])); err != nil {
		return m, err
	}
	av, err := fetch.JSON[androidVersions](ctx, s.client, s.endpoints.AndroidVersions)
	if err != nil {
		return m, err
	}
	minimum, err := strconv.ParseUint(av.Minimum, 10, 32)
	if err != nil {
		return m, &ParseError{Kind: "android minimum version", Input: av.Minimum, Err: err}
	}
	m.AndroidMinimum = uint32(minimum)
	lookup, err := fetch.JSON[itunesLookup](ctx, s.client, s.endpoints.ITunesLookup)
	if err != nil {
		return m, err
	}
	if len(lookup.Results) == 0 {
		return m, &ExtractError{Source: "itunes lookup", Field: "results"}
	}
	m.IOSMinimumOS = lookup.Results[0].MinimumOSVersion
	m.IOS, err = ParseMobileVersion(lookup.Results[0].Version)
	return m, err
}
