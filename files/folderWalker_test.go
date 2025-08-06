package files

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetAllFilePaths_WhenValidInput_ReturnsFilePaths(t *testing.T) {
	//Given
	mockRootPath := "./testData"
	isGitIgnoreDisabled := true
	isForceIgnoreDisabled := true
	MOCKED_FILE_OPTIONS := FileOptions{
		RootPath:        mockRootPath,
		DontGitIgnore:   isGitIgnoreDisabled,
		DontForceIgnore: isForceIgnoreDisabled,
	}

	filepath1, _ := filepath.Abs("./testData/sampleField-meta.xml")
	filepath2, _ := filepath.Abs("./testData/testFile.cls")
	filepath3, _ := filepath.Abs("./testData/testFile2.cls")
	expectedResult := []string{filepath1, filepath2, filepath3}

	//When
	actualResult, _ := GetAllFilePaths(MOCKED_FILE_OPTIONS)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Filepaths are mismatched!", actualResult, expectedResult)
	}
}

func TestGetAllFilePaths_WhenOsStatFails_ReturnsError(t *testing.T) {
	//Given
	mockRootPath := "./invalid/path"
	isGitIgnoreDisabled := true
	isForceIgnoreDisabled := true
	MOCKED_FILE_OPTIONS := FileOptions{
		RootPath:        mockRootPath,
		DontGitIgnore:   isGitIgnoreDisabled,
		DontForceIgnore: isForceIgnoreDisabled,
	}
	expectedResult := []string{}

	//When
	actualResult, err := GetAllFilePaths(MOCKED_FILE_OPTIONS)

	//Then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	if len(actualResult) != len(expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Filepaths are mismatched!", actualResult, expectedResult)
	}
}

func TestGetAllFilePaths_WhenSingleFileScans_ReturnsSameFilePath(t *testing.T) {
	//Given
	mockFilePath := "./testData/testFile.cls"
	expectedFiles := []string{mockFilePath}

	//When
	actualFiles, _ := GetAllFilePaths(FileOptions{RootPath: mockFilePath})

	//Then
	if !reflect.DeepEqual(actualFiles, expectedFiles) {
		t.Errorf("Actual %+v, Expected %+v FilePaths are mismatched!", actualFiles, expectedFiles)
	}
}
