package image

import (
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/irdaislakhuafa/rsic-go/src/utility"
)

var (
	supportedFormat = map[string](func() (utility.ImgDecoder, utility.ImgEncoder, string, error)){
		".png": func() (utility.ImgDecoder, utility.ImgEncoder, string, error) {
			decoder := func(r io.Reader) (image.Image, string, error) {
				img, err := png.Decode(r)
				if err != nil {
					return nil, "", err
				}
				return img, "", nil
			}
			return decoder, png.Encode, "", nil
		},
	}
)

func GetDecoderEncoderFromExt(fileExt string) (utility.ImgDecoder, utility.ImgEncoder, string, error) {
	decEnc, isExists := supportedFormat[fileExt]
	if !isExists {
		return nil, nil, "", fmt.Errorf("extension format '%v' not supported, currently supported format is %+v", fileExt, GetSupportedFormat())
	}
	return decEnc()
}

func GetSupportedFormat() []string {
	supported := []string{}
	for k := range supportedFormat {
		supported = append(supported, k)
	}
	return supported
}
