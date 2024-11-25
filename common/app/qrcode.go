package app

import (
	"github.com/skip2/go-qrcode"
)

func QrcodeEncode(content string, size int) []byte {
	qrByte, err := qrcode.Encode(content, qrcode.Medium, size)
	if err == nil {
		return qrByte
	} else {
		return nil
	}
}
