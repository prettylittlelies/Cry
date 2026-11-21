package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Field struct {
	Name        string
	Optional    bool
	Type        string
	Description string
}

type Structure struct {
	Doc    string
	Title  string
	Anchor string
	Fields []Field
}

type EnumValue struct {
	Value       string
	Name        string
	Description string
}

type Enum struct {
	Doc    string
	Title  string
	Anchor string
	Values []EnumValue
}

type Endpoint struct {
	Doc      string
	Method   string
	URL      string
	Name     string
	Params   *Structure
	Query    *Structure
	Response *Structure
	Prose    []string
}

type GatewayEvent struct {
	Name    string
	Prose   []string
	Structs []*Structure
}

type Spec struct {
	Structures []*Structure
	Enums      []*Enum
	Endpoints  []*Endpoint
	Events     []*GatewayEvent
}

var (
	headingPattern  = regexp.MustCompile(`^#{2,6}\s+(.+?)\s*$`)
	footnotePattern = regexp.MustCompile(`\^[^^]*\^`)
	boldPattern     = regexp.MustCompile(`\*\*[^*]*\*\*`)
	linkPattern     = regexp.MustCompile(`\[([^\[\]]*)\]\(([^)]*)\)`)
	tagPattern      = regexp.MustCompile(`<[^>]+>`)
	routePattern    = regexp.MustCompile(`<RouteHeader\s+method="([A-Z]+)"\s+url="([^"]+)"`)
	sectionPattern  = regexp.MustCompile(`^#{1,3}\s`)
	parenShift      = regexp.MustCompile(`\((\d+\s*<<\s*\d+)\)`)
)

func anchor(title string) string {
	s := strings.ToLower(tagPattern.ReplaceAllString(title, ""))
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-':
			b.WriteRune('-')
		}
	}
	return strings.Trim(b.String(), "-")
}

func cells(line string) []string {
	parts := strings.Split(strings.Trim(strings.TrimSpace(line), "|"), "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func clean(s string) string {
	s = footnotePattern.ReplaceAllString(s, "")
	s = boldPattern.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "`", "")
	s = strings.ReplaceAll(s, "~~", "")
	s = strings.ReplaceAll(s, `\`, "")
	return strings.TrimSpace(s)
}

func plainType(s string) string {
	return strings.TrimSpace(linkPattern.ReplaceAllString(clean(s), "$1"))
}

func ParseDir(root string) (*Spec, error) {
	spec := &Spec{}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".mdx" {
			return err
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		parseDoc(spec, strings.TrimSuffix(filepath.Base(path), ".mdx"), string(raw))
		return nil
	})
	return spec, err
}

func parseDoc(spec *Spec, doc, text string) {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	title := ""
	var (
		route *Endpoint
		event *GatewayEvent
	)
	for i := 0; i < len(lines); i++ {
		if m := routePattern.FindStringSubmatch(lines[i]); m != nil {
			route = &Endpoint{Doc: doc, Method: m[1], URL: m[2]}
			for i++; i < len(lines) && !strings.Contains(lines[i], "</RouteHeader>"); i++ {
				if n := strings.TrimSpace(tagPattern.ReplaceAllString(lines[i], "")); n != "" && route.Name == "" {
					route.Name = n
				}
			}
			if route.Name != "" {
				spec.Endpoints = append(spec.Endpoints, route)
			}
			title = ""
			continue
		}
		if sectionPattern.MatchString(lines[i]) {
			route = nil
		}
		if m := headingPattern.FindStringSubmatch(lines[i]); m != nil {
			title = tagPattern.ReplaceAllString(m[1], "")
			if doc == "gateway-events" && strings.HasPrefix(lines[i], "#### ") {
				event = &GatewayEvent{Name: title}
				spec.Events = append(spec.Events, event)
			} else if strings.HasPrefix(lines[i], "## ") || strings.HasPrefix(lines[i], "### ") {
				event = nil
			}
			continue
		}
		if event != nil && !strings.HasPrefix(lines[i], "|") {
			event.Prose = append(event.Prose, lines[i])
		}
		if route != nil && !strings.HasPrefix(lines[i], "|") && !strings.HasPrefix(lines[i], "#") {
			route.Prose = append(route.Prose, lines[i])
		}
		if !strings.HasPrefix(lines[i], "|") || i+1 >= len(lines) || !strings.HasPrefix(lines[i+1], "| -") {
			continue
		}
		header := cells(lines[i])
		var rows [][]string
		for i += 2; i < len(lines) && strings.HasPrefix(lines[i], "|"); i++ {
			rows = append(rows, cells(lines[i]))
		}
		i--
		structTitle, slot := structureTitle(title, route)
		switch {
		case structTitle != "" && len(header) >= 2 && (header[0] == "Field" || header[0] == "Name") && header[1] == "Type":
			s := &Structure{Doc: doc, Title: structTitle, Anchor: anchor(title)}
			for _, r := range rows {
				if len(r) < 2 {
					continue
				}
				name := clean(r[0])
				optional := strings.HasSuffix(name, "?")
				name = strings.TrimSuffix(name, "?")
				if name == "" || strings.ContainsAny(name, " .") {
					continue
				}
				desc := ""
				if len(r) > 2 {
					desc = r[2]
				}
				s.Fields = append(s.Fields, Field{Name: name, Optional: optional, Type: plainType(r[1]), Description: desc})
			}
			if len(s.Fields) > 0 {
				spec.Structures = append(spec.Structures, s)
				if slot != nil {
					*slot = s
				}
				if event != nil {
					event.Structs = append(event.Structs, s)
				}
			}
		case len(header) >= 2 && header[0] == "Value" && (header[1] == "Name" || header[1] == "Description"):
			named := header[1] == "Name"
			e := &Enum{Doc: doc, Title: title, Anchor: anchor(title)}
			for _, r := range rows {
				if len(r) < 2 {
					continue
				}
				value := clean(r[0])
				if m := parenShift.FindStringSubmatch(value); m != nil {
					value = m[1]
				}
				if named {
					desc := ""
					if len(r) > 2 {
						desc = r[2]
					}
					e.Values = append(e.Values, EnumValue{Value: value, Name: clean(r[1]), Description: desc})
				} else {
					e.Values = append(e.Values, EnumValue{Value: value, Name: value, Description: r[1]})
				}
			}
			if len(e.Values) > 0 && title != "" {
				spec.Enums = append(spec.Enums, e)
			}
		}
	}
}

func structureTitle(title string, route *Endpoint) (string, **Structure) {
	switch {
	case strings.HasSuffix(title, " Extra Fields"):
		return strings.Replace(title, " Structure Extra Fields", " Extra Fields", 1), nil
	case strings.HasSuffix(title, " Structure"):
		return strings.TrimSuffix(title, " Structure"), nil
	case strings.HasSuffix(title, " Object"):
		return strings.TrimSuffix(title, " Object"), nil
	case route == nil:
		return "", nil
	case title == "Response Body":
		return route.Name + " Response", &route.Response
	case title == "JSON Params" || title == "JSON/Form Params" || title == "Form Params":
		return route.Name + " Params", &route.Params
	case title == "Query String Params":
		return route.Name + " Query", &route.Query
	}
	return "", nil
}
