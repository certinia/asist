package ignore

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/certinia/asist/utils"
)

const TEST_ENVIRONMENT_DIRECTORY = "testData"

func TestMain(m *testing.M) {
	//Preparing setup
	createTestEnvironment()
	//Exit test case with exited code
	exitCode := m.Run()
	//Cleaning all the setup
	utils.DeleteFolder(TEST_ENVIRONMENT_DIRECTORY)
	//Exit process
	os.Exit(exitCode)
}

func TestGetIgnoreFilesPatterns_WhenGitIgnoreAndForceIgnoreEnabled_ReturnsFilesInIgnoreList(t *testing.T) {
	//Given
	currentAbsolutePath, _ := filepath.Abs(TEST_ENVIRONMENT_DIRECTORY)
	expectedIgnoredFilePath := currentAbsolutePath + "/testDataForGitIgnore.txt"
	fileOptions := IgnoreOptions{
		RootPath:        currentAbsolutePath,
		DontGitIgnore:   false,
		DontForceIgnore: false,
	}

	//When
	ignoreList, _ := GetIgnoreFilesPatterns(fileOptions)

	//Then
	actualResult := ignoreList.Match(strings.Split(filepath.ToSlash(expectedIgnoredFilePath), "/"), false)
	if !actualResult {
		t.Errorf("Ignore list doesn't have entry of testDataForGitIgnore!!!")
	}
}

func TestCreateIgnoreList_WhenGitIgnoreAndForceIgnoreDisabled_ReturnsFilesExcludedFromIgnoreList(t *testing.T) {
	//Given
	currentAbsolutePath, _ := filepath.Abs(TEST_ENVIRONMENT_DIRECTORY)
	fileOptions := IgnoreOptions{
		RootPath:        currentAbsolutePath,
		DontGitIgnore:   true,
		DontForceIgnore: true,
	}

	//When
	actualIgnoreList, _ := GetIgnoreFilesPatterns(fileOptions)

	//Then
	if actualIgnoreList != nil {
		t.Errorf("Should not found the match for disabled git,force ignore files!")
	}
}

func createTestEnvironment() {
	utils.CreateFolder(TEST_ENVIRONMENT_DIRECTORY)
	utils.WriteFile(TEST_ENVIRONMENT_DIRECTORY+"/.gitignore", []byte("testDataForGitIgnore.txt"))
	utils.WriteFile(TEST_ENVIRONMENT_DIRECTORY+"/testDataForGitIgnore.txt", []byte("Sample data for the gitignore functionality!!"))
}
