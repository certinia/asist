package message

import (
	"fmt"
)

var TextColor = map[string]string{
	"Error":   "\033[0;31m", // Red
	"Info":    "\033[0;32m", // Green
	"Warning": "\033[0;33m", // Yellow
	"Debug":   "\033[0;34m", // Blue
	"Reset":   "\033[0m",
}

func SetLogType(logType string, msg string) string {
	return TextColor[logType] + logType + TextColor["Reset"] + ": " + msg
}
func GetInvalidRuleIdWarning(ruleId string) string {
	return SetLogType("Warning", fmt.Sprintf("Invalid rule ID %s\n", ruleId))
}

func GetMissingFileOrFolderError() string {
	return "Specify a file or folder path to scan\n"
}

func GetPathFetchingError(err error) string {
	return fmt.Sprintf("Error fetching path: %+v\n", err)
}

func GetFilesFetchingError(err error) string {
	return fmt.Sprintf("Error in fetching files from directory: %+v\n", err)
}

func GetMarshallingOutputError(err error) string {
	return fmt.Sprintf("Error marshalling output: %+v\n", err)
}

func GetFileReadError(fileName string, err error) string {
	return fmt.Sprintf("Error reading file %s: %v", fileName, err.Error())
}

func GetInvalidTemplateFileError(fileError error) string {
	return fmt.Sprintf("Invalid template file  %s", fileError)
}

func GetInvalidConfigFileError(path string) string {
	return fmt.Sprintf("Invalid config file. Expecting a \"yaml\" or \"json\" extension %s", path)
}

func GetFileUnmarshalingError(err error) string {
	return fmt.Sprintf("Error during Unmarshaling of file %s", err)
}

func GetMarshalIndentError(err error) string {
	return fmt.Sprintf("%v", err)
}
