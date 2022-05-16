package utils

import (
	"path"
	"strings"
)

func GetFileNameWithoutExt(filePath string) string {
	fileName := path.Base(filePath)
	return fileName[:len(fileName)-len(path.Ext(fileName))]
}

// Return extension(without dot) if path is a file.
func GetFileExtension(filePath string) string {
	if filePath == "" {
		return ""
	}
	ext := path.Ext(filePath)
	if ext == "" {
		return ""
	}
	return strings.Split(ext, ".")[1]
}
