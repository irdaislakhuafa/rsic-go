package images

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"path/filepath"

	imgErr "github.com/irdaislakhuafa/rsic-go/errors"
	"github.com/irdaislakhuafa/rsic-go/files"
)

type ImgDecoder func(r io.Reader) (image.Image, string, error)
type ImgEncoder func(w io.Writer, m image.Image) error

var (
	// Supported format to decode and encode
	supportedFormat = map[string](func() (ImgDecoder, ImgEncoder, string, error)){
		".png": func() (ImgDecoder, ImgEncoder, string, error) {
			decoder := func(r io.Reader) (image.Image, string, error) {
				img, err := png.Decode(r)
				if err != nil {
					return nil, "", errors.Join(imgErr.ErrImgDecoder, err)
				}
				return img, "", nil
			}
			return decoder, png.Encode, "", nil
		},
	}
)

// Get specific decoder and encoder image from specific extension and return error if not extension not supported
func GetDecoderEncoderFromExt(fileExt string) (ImgDecoder, ImgEncoder, string, error) {
	decEnc, isExists := supportedFormat[fileExt]
	if !isExists {
		return nil, nil, "", errors.Join(imgErr.New(fmt.Sprintf("currently supported format is %+v", GetSupportedFormat())), imgErr.ErrExtFormatNotSupported)
	}
	return decEnc()
}

// Method to get list of supported format
func GetSupportedFormat() []string {
	supported := []string{}
	for k := range supportedFormat {
		supported = append(supported, k)
	}
	return supported
}

// replace specific color from image and return image.RGBA if success
func ReplaceSpecificImgColorWithWrite(sourceImg, outputImg string, requiredColor, replacementColor color.RGBA) (string, error) {
	sourceImgExt, err := files.GetFileExtension(sourceImg)
	if err != nil {
		return "", err
	}

	_, encode, _, err := GetDecoderEncoderFromExt(*sourceImgExt)
	if err != nil {
		return "", err
	}

	newOutputImg, err := ReplaceSpecificImgColorWithoutWrite(sourceImg, requiredColor, replacementColor)
	if err != nil {
		return "", err
	}

	// read output image dir and create it if not exists
	outputDirPath := filepath.Dir(outputImg)
	if _, err := os.Stat(outputDirPath); err != nil {
		if os.IsNotExist(err) {
			// create dirs with default os permission
			if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
				return "", errors.Join(imgErr.New(fmt.Sprintf("failed to create output dirs '%v'", outputDirPath)), imgErr.ErrFailedCreateOutputDir, err)
			}
		} else {
			return "", errors.Join(imgErr.New(fmt.Sprintf("error while reading directory '%v'", outputDirPath)), imgErr.ErrFailedReadOutputDir, err)
		}
	}

	// create empty file for output image
	outputImgFile, err := os.Create(outputImg)
	if err != nil {
		return "", errors.Join(imgErr.New(fmt.Sprintf("cannot create file '%v'", outputImg)), imgErr.ErrFailedCreateOutputFile, err)
	}
	defer outputImgFile.Close()

	// encode it and save image same as source format
	if err := encode(outputImgFile, newOutputImg); err != nil {
		return "", errors.Join(imgErr.New(fmt.Sprintf("cannot encode new output image")), imgErr.ErrImgEncoder, err)
	}

	return fmt.Sprintf("image processing complete and saved to %v", outputImg), nil
}

func ReplaceSpecificImgColorWithoutWrite(sourceImg string, requiredColor, replacementColor color.RGBA) (*image.RGBA, error) {
	// open source image file
	sourceFile, err := os.Open(sourceImg)
	if err != nil {
		return nil, errors.Join(imgErr.ErrFailedOpenSrcFile, err)
	}
	defer sourceFile.Close()

	// get file extension
	sourceImgExt, err := files.GetFileExtension(sourceImg)
	if err != nil {
		return nil, err
	}

	// get decoder and encoder for supported file extension
	decode, _, _, err := GetDecoderEncoderFromExt(*sourceImgExt)
	if err != nil {
		return nil, err
	}

	// decode source image file
	decodedSrcImg, _, err := decode(sourceFile)
	if err != nil {
		return nil, errors.Join(imgErr.New(fmt.Sprintf("failed to decode source file '%v'", sourceImg)), err)
	}

	srcImgWidth := decodedSrcImg.Bounds().Dx()
	srcImgHeight := decodedSrcImg.Bounds().Dy()
	newOutputImg := image.NewRGBA(image.Rect(0, 0, srcImgWidth, srcImgHeight))

	// loop each pixel of image, compare current pixel color with required color and if match change it with replacement color
	for width := 0; width < srcImgWidth; width++ {
		for height := 0; height < srcImgHeight; height++ {
			// get pixel color at pixel (x, y)/(width, height)
			srcImgPixelColor := decodedSrcImg.At(width, height)

			if srcImgPixelColor == requiredColor {
				newOutputImg.Set(width, height, replacementColor)
			} else {
				newOutputImg.Set(width, height, srcImgPixelColor)
			}
		}
	}

	return newOutputImg, nil
}
