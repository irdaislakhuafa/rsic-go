package utility

import (
	"image"
	"io"
)

type ImgDecoder func(r io.Reader) (image.Image, string, error)
type ImgEncoder func(w io.Writer, m image.Image) error
