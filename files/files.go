package files

import (
	"fmt"
	"strings"
)

func GetFileExtension(filePathName string) (*string, error) {
	splitedString := strings.Split(filePathName, ".")
	if len(splitedString) <= 0 {
		return nil, fmt.Errorf("file '%v' doesn't have any extension", filePathName)
	}

	fileExtension := "." + splitedString[len(splitedString)-1]
	return &fileExtension, nil
}
