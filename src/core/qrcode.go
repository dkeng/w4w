package core

import (
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"
)

// CreateQrcodeBase64 创建二维码
func CreateQrcodeBase64(s string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(s, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(png), nil
}
