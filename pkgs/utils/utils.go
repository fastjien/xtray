package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
)

func NormalizeSSR(s string) (r string) {
	r = strings.ReplaceAll(s, "-", "+")
	r = strings.ReplaceAll(r, "_", "/")
	return r
}

func DecodeBase64(str string) (res string) {
	if s, err := base64.StdEncoding.DecodeString(str); err == nil {
		res = string(s)
	}
	return
}

func StringToReader(str string) io.Reader {
	return bytes.NewReader([]byte(str))
}
