package qrutils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode generates a QR code from the provided text.
// Returns the image bytes in PNG format.
func GenerateQRCode(text string, size int) ([]byte, error) {
	return qrcode.Encode(text, qrcode.Medium, size)
}

// SaveQRCode saves the image bytes in the specified format.
// Supported formats: "png", "jpg", "jpeg".
func SaveQRCode(data []byte, format, filename string) error {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	switch format {
	case "png":
		err = png.Encode(buf, img)
	case "jpg", "jpeg":
		err = jpeg.Encode(buf, img, nil)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
