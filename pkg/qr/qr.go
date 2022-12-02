package qr

import (
	"bytes"

	qrcode "github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type QR struct {
	Buffer *bytes.Buffer
}

func (q QR) Close() error {
	return nil
}

func (q QR) Write(p []byte) (n int, err error) {
	return q.Buffer.Write(p)
}

func (q QR) Bytes() []byte {
	return q.Buffer.Bytes()
}

func QRWriter() (QR, *standard.Writer) {
	qr := QR{Buffer: &bytes.Buffer{}}

	// , standard.WithBgColor(color.RGBA{R: 3, G: 3, B: 4, A: 3})
	// standard.WithCircleShape()
	return qr, standard.NewWithWriter(
		qr,
		standard.WithQRWidth(14),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithBorderWidth(12),
	)
}

func New(content string, writer qrcode.Writer) error {
	q, err := qrcode.New(content)
	if err != nil {
		return err
	}

	err = q.Save(writer)
	if err != nil {
		return err
	}

	return nil
}
