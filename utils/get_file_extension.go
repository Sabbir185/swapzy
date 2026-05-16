package utils

import (
	"strings"
)

func GetFileExtension(filePath string) string {
	e := strings.TrimPrefix(filePath, "./data/config.")
	return e
}
