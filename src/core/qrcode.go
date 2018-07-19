package core

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	qrcode "github.com/skip2/go-qrcode"
)

// CreateQrcodeBase64 创建二维码
func CreateQrcodeBase64(s string) (string, error) {
	var pngQrcode []byte
	pngQrcode, err := qrcode.Encode(s, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	pngBuffer := bytes.NewBuffer(pngQrcode)
	// 图片文件解码
	img, _, err := image.Decode(pngBuffer)
	if err != nil {
		return "", err
	}
	rgbImg := img.(*image.Paletted)
	//图片裁剪x0 y0 x1 y1
	subImg := rgbImg.SubImage(image.Rect(35, 35, 220, 220)).(*image.Paletted)
	emptyBuffer := bytes.NewBuffer(nil)
	png.Encode(emptyBuffer, subImg)
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuffer.Bytes()), nil
}
