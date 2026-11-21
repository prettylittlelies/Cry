package build

import (
	"context"
	"regexp"
	"strconv"
)

var buildNumberPattern = regexp.MustCompile(`"BUILD_NUMBER"\s*:\s*"(\d+)"`)

func (s *Scraper) number(ctx context.Context, c Channel) (Number, error) {
	page, err := s.client.Bytes(ctx, s.endpoints.App(c))
	if err != nil {
		return 0, err
	}
	m := buildNumberPattern.FindSubmatch(page)
	if m == nil {
		return 0, &ExtractError{Source: "app page", Field: "BUILD_NUMBER"}
	}
	n, err := strconv.ParseUint(string(m[1]), 10, 32)
	if err != nil {
		return 0, &ParseError{Kind: "build number", Input: string(m[1]), Err: err}
	}
	return Number(n), nil
}
