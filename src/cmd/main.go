package main

import (
	"fmt"
	"os"

	imageUsecase "github.com/irdaislakhuafa/rsic-go/src/business/usecase/image"
	"github.com/irdaislakhuafa/rsic-go/src/utility"
)

func main() {
	// path of needed image files
	sourceImg := "docs/img/4.2.png"
	outputImg := "docs/img/4.2-bg-white.png"

	if err := ReplaceSpecificImgColor(sourceImg, outputImg); err != nil {
		panic(err)
	}
}

func ReplaceSpecificImgColor(sourceImg, outputImg string) error {
	// open source image file
	sourceFile, err := os.Open(sourceImg)
	if err != nil {
		return fmt.Errorf("failed to open source file '%v', %v", sourceFile, err)
	}
	defer sourceFile.Close()

	// get file extension
	sourceImgExt, err := utility.GetFileExtension(sourceImg)
	if err != nil {
		return err
	}

	// get decoder and encoder for supported file extension
	decode, _, _, err := imageUsecase.GetDecoderEncoderFromExt(*sourceImgExt)
	if err != nil {
		return err
	}

	// decode source image file
	decodedSrcImg, _, err := decode(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to decode source file '%v', %v", sourceImg, err)
	}

	srcImgWidth := decodedSrcImg.Bounds().Dx()
	srcImgHeight := decodedSrcImg.Bounds().Dy()
	// newOutputImg := image.NewRGBA(image.Rect(0, 0, srcImgWidth, srcImgHeight))

	for width := 0; width < srcImgWidth; width++ {
		for height := 0; height < srcImgHeight; height++ {
			// get pixel color at pixel (x, y)/(width, height)
			srcImgPixelColor := decodedSrcImg.At(width, height)

			// if srcImgPixelColor {

			// }
		}

	}

	return nil
}
