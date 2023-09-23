package files

import (
	"strings"

	imgErr "github.com/irdaislakhuafa/rsic-go/errors"
)

func GetFileExtension(filePathName string) (*string, error) {
	splitedString := strings.Split(filePathName, ".")
	if len(splitedString) <= 0 {
		return nil, imgErr.ErrNotHaveExt
	}

	fileExtension := "." + splitedString[len(splitedString)-1]
	return &fileExtension, nil
}
