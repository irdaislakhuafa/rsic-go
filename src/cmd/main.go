package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"github.com/irdaislakhuafa/rsic-go/files"
	imageUsecase "github.com/irdaislakhuafa/rsic-go/src/usecase/image"
)

func main() {
	// TODO: set this as parameter/flags
	// path of needed image files
	sourceImg := "docs/img/4.2.png"
	outputImg := "output/img/4.2.png"

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
	sourceImgExt, err := files.GetFileExtension(sourceImg)
	if err != nil {
		return err
	}

	// get decoder and encoder for supported file extension
	decode, encode, _, err := imageUsecase.GetDecoderEncoderFromExt(*sourceImgExt)
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
	newOutputImg := image.NewRGBA(image.Rect(0, 0, srcImgWidth, srcImgHeight))

	// TODO: set this in args/parameter
	requiredColor := color.RGBA{R: 14, G: 14, B: 18, A: 255}
	replacementColor := color.RGBA{R: 244, G: 245, B: 245, A: 255}

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

	// read output image dir and create it if not exists
	outputDirPath := filepath.Dir(outputImg)
	if _, err := os.Stat(outputDirPath); err != nil {
		if os.IsNotExist(err) {
			// create dirs with default os permission
			if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create output dirs '%v', %v", outputDirPath, err)
			} else {
				fmt.Printf("directory '%v' doesn't exist and automatic created\n", outputDirPath)
			}
		} else {
			return fmt.Errorf("error while reading directory '%v', %v", outputDirPath, err)
		}
	}

	// create empty file for output image
	outputImgFile, err := os.Create(outputImg)
	if err != nil {
		return fmt.Errorf("cannot create file '%v', %v", outputImg, err)
	}
	defer outputImgFile.Close()

	// encode it and save image same as source format
	if err := encode(outputImgFile, newOutputImg); err != nil {
		return fmt.Errorf("cannot encode new output image, %v", err)
	}

	fmt.Printf("image processing complete and saved to %v\n", outputImg)
	return nil
}
