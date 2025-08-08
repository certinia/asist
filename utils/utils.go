package utils

import (
	"errors"
	"io/fs"
	"os"

	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/message"
)

/**
* CreateRegex - method used to create regex at runtime by combining multiple patterns
 */
func CreateRegex(patterns []string) string {
	combinedPattern := ""
	regexIndex := 0
	if len(patterns) > 0 {
		for regexIndex < len(patterns)-1 {
			combinedPattern += "(" + patterns[regexIndex] + ")" + "|"
			regexIndex++
		}
		combinedPattern += "(" + patterns[regexIndex] + ")"
	}
	return combinedPattern
}

/**
* IsDirectory - method used to check a given path is a directory or not
 */
func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, errorhandler.NewInternalError(message.GetPathFetchingError(err))
	}
	return info.IsDir(), nil
}

/**
* IsFileExists - method used to check a given file path exists or not
 */
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}
