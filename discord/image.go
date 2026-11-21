package discord

import (
	"encoding/base64"
	"net/http"
)

type ImageData string

func NewImageData(data []byte) ImageData {
	return ImageData("data:" + http.DetectContentType(data) + ";base64," + base64.StdEncoding.EncodeToString(data))
}
