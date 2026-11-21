package rest

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"strconv"
	"strings"
)

const (
	boundaryPrefix   = "----WebKitFormBoundary"
	boundaryAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789AB"
	octetStream      = "application/octet-stream"
)

type File struct {
	Name        string
	ContentType string
	Data        []byte
}

type Part struct {
	Name  string
	File  *File
	Value string
}

type Form struct {
	Parts []Part
}

func MessageForm(payload any, files ...File) (*Form, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	f := &Form{Parts: make([]Part, 0, len(files)+1)}
	for i := range files {
		f.Parts = append(f.Parts, Part{Name: "files[" + strconv.Itoa(i) + "]", File: &files[i]})
	}
	f.Parts = append(f.Parts, Part{Name: "payload_json", Value: string(raw)})
	return f, nil
}

func boundary() string {
	var raw [16]byte
	rand.Read(raw[:])
	b := []byte(boundaryPrefix)
	for _, c := range raw {
		b = append(b, boundaryAlphabet[c&0x3f])
	}
	return string(b)
}

func quote(buf *bytes.Buffer, s string, normalize bool) {
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == '\n' && normalize:
			buf.WriteString("%0D%0A")
		case c == '\n':
			buf.WriteString("%0A")
		case c == '\r' && normalize:
			buf.WriteString("%0D%0A")
			if i+1 < len(s) && s[i+1] == '\n' {
				i++
			}
		case c == '\r':
			buf.WriteString("%0D")
		case c == '"':
			buf.WriteString("%22")
		default:
			buf.WriteByte(c)
		}
	}
}

func normalizeCRLF(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return strings.ReplaceAll(s, "\n", "\r\n")
}

func (f *Form) Encode() ([]byte, string) {
	b := boundary()
	var buf bytes.Buffer
	for _, p := range f.Parts {
		buf.WriteString("--" + b + "\r\n")
		buf.WriteString(`Content-Disposition: form-data; name="`)
		quote(&buf, p.Name, true)
		buf.WriteByte('"')
		if p.File == nil {
			buf.WriteString("\r\n\r\n")
			buf.WriteString(normalizeCRLF(p.Value))
			buf.WriteString("\r\n")
			continue
		}
		buf.WriteString(`; filename="`)
		quote(&buf, p.File.Name, false)
		buf.WriteString("\"\r\nContent-Type: ")
		ct := p.File.ContentType
		if ct == "" {
			ct = octetStream
		}
		buf.WriteString(ct)
		buf.WriteString("\r\n\r\n")
		buf.Write(p.File.Data)
		buf.WriteString("\r\n")
	}
	buf.WriteString("--" + b + "--\r\n")
	return buf.Bytes(), "multipart/form-data; boundary=" + b
}
