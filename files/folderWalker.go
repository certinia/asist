package files

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/certinia/asist/debugger"
	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/files/ignore"
	"github.com/certinia/asist/message"
)

type FileOptions struct {
	RootPath        string
	DontGitIgnore   bool
	DontForceIgnore bool
}

/**
 * GetAllFilePaths - method used to get all file paths present in the user provided root path.
 */
func GetAllFilePaths(fileOptions FileOptions) ([]string, error) {
	var filePaths []string

	info, err := os.Stat(fileOptions.RootPath)
	if err != nil {
		return nil, errorhandler.NewUserError(message.GetFileReadError(fileOptions.RootPath, err))
	}
	ignoreFileOptions := ignore.IgnoreOptions{
		RootPath:        fileOptions.RootPath,
		DontGitIgnore:   fileOptions.DontGitIgnore,
		DontForceIgnore: fileOptions.DontForceIgnore,
	}
	if !info.IsDir() {
		return []string{fileOptions.RootPath}, nil
	}
	matcher, err := ignore.GetIgnoreFilesPatterns(ignoreFileOptions)
	if err != nil {
		return nil, err
	}
	debugger.Debug("created .gitignore list")

	//ignoreFilesList is an instance of gitignore package. It holds the list of all ignore files which should be ignored from the provided rootpath.
	err = filepath.Walk(fileOptions.RootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(fileOptions.RootPath, path)
		if err != nil {
			return err
		}

		if matcher != nil && matcher.Match(strings.Split(filepath.ToSlash(relPath), "/"), info.IsDir()) {
			return nil
		}

		if !info.IsDir() {
			currentFilePath, absError := filepath.Abs(path)
			if absError != nil {
				return absError
			}
			filePaths = append(filePaths, currentFilePath)
		}
		return nil
	})
	debugger.Debug("walked directory")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, errorhandler.NewInternalError(message.GetFilesFetchingError(err))
	}
	return filePaths, nil
}
