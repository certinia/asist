package config

import (
	"reflect"
	"strings"
	"testing"

	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/utils"
)

func TestParseConfig_WhenYAMLConfigFileExist_ReturnsConfigInstance(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.yaml"

	ENABLED_TRUE := true
	ENABLED_FALSE := false
	MAX_ISSUES_10 := 10
	CUSTOM_MAX_5 := 5

	expectedConfigFile := Config{
		EnableAllStandardRules: &ENABLED_TRUE,
		DontGitIgnore:          true,
		DontForceIgnore:        true,
		ExcludeFilesAndFolders: []string{"/force-app-autotests/"},
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				Severity:      "Medium",
				Enabled:       &ENABLED_TRUE,
				CicdMaxIssues: &MAX_ISSUES_10,
			},
		},
		CustomRegexRules: map[string]CustomRegexRule{
			"CustomRule1": {
				Name:           "customName1",
				Description:    "Please fix this",
				Enabled:        &ENABLED_FALSE,
				Severity:       "High",
				RuleCategory:   "Security",
				Pattern:        "Label",
				IncludePattern: "\\.component$|\\.page$|\\.cls$|\\.email",
				ExcludePattern: "",
				CicdMaxIssues:  &CUSTOM_MAX_5,
			},
		},
		CICDRules: []string{
			"XSSLabel",
		},
	}

	//When
	actualConfigFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if err != nil {
		t.Errorf("Parse config method should not return error!")
	}
	if !reflect.DeepEqual(actualConfigFile, &expectedConfigFile) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file should be equal!", actualConfigFile, expectedConfigFile)
	}
}

func TestParseConfig_WhenYAMLConfigFileNotExist_ReturnsFileNotExistError(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/test.yaml"
	expectedFileNotExistError := "Invalid template file  open ./testData/test.yaml"

	//When
	_, actualFileNotExistError := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !strings.Contains(actualFileNotExistError.Error(), expectedFileNotExistError) {
		t.Errorf("File not exist error mismatched!")
	}
}

func TestParseConfig_WhenYAMLConfigFileWithErrorExist_ReturnsYamlParsingError(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/test.yaml"
	expectedError := "Error during Unmarshaling of file yaml: unmarshal errors"
	utils.CreateFile(MOCK_CONFIG_FILE_PATH)
	utils.WriteFile(MOCK_CONFIG_FILE_PATH, []byte("123"))

	//When
	_, err := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Parse config error mismatched!")
	}
	utils.DeleteFile(MOCK_CONFIG_FILE_PATH)
}

func TestParseConfig_WhenJSONConfigFileExist_ReturnsConfigInstance(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.json"
	ENABLED_TRUE := true
	ENABLED_FALSE := false
	MAX_ISSUES_10 := 10
	CUSTOM_MAX_5 := 5

	expectedConfigFile := Config{
		EnableAllStandardRules: &ENABLED_TRUE,
		DontGitIgnore:          true,
		DontForceIgnore:        true,
		ExcludeFilesAndFolders: []string{"/force-app-autotests/"},
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				Severity:      "Medium",
				Enabled:       &ENABLED_TRUE,
				CicdMaxIssues: &MAX_ISSUES_10,
			},
		},
		CustomRegexRules: map[string]CustomRegexRule{
			"CustomRule1": {
				Name:           "customName1",
				Description:    "Please fix this",
				Enabled:        &ENABLED_FALSE,
				Severity:       "High",
				RuleCategory:   "Security",
				Pattern:        "Label",
				IncludePattern: "\\.component$|\\.page$|\\.cls$|\\.email",
				ExcludePattern: "",
				CicdMaxIssues:  &CUSTOM_MAX_5,
			},
		},
		CICDRules: []string{
			"XSSLabel",
		},
	}

	//When
	actualConfigFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !reflect.DeepEqual(actualConfigFile, &expectedConfigFile) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file should be equal!", actualConfigFile, expectedConfigFile)
	}
	if err != nil {
		t.Errorf("Parse config method should not return error!")
	}
}

func TestParseConfig_WhenJSONConfigFileNotExist_ReturnsFileNotExistError(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/test.json"
	expectedFileNotExistError := "Invalid template file  open ./testData/test.json"

	//When
	_, actualFileNotExistError := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !strings.Contains(actualFileNotExistError.Error(), expectedFileNotExistError) {
		t.Errorf("File not exist error mismatched!")
	}
}

func TestParseConfig_WhenJSONConfigFileWithErrorExist_ReturnsJSONParsingError(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/test.json"
	expectedError := "Error during Unmarshaling of file json"
	utils.CreateFile(MOCK_CONFIG_FILE_PATH)
	utils.WriteFile(MOCK_CONFIG_FILE_PATH, []byte("123"))

	//When
	_, err := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Parse config error mismatched!")
	}
	utils.DeleteFile(MOCK_CONFIG_FILE_PATH)
}

func TestParseConfig_WhenEmptyConfigFilePath_ReturnsNil(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = ""

	//When
	configFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if !(configFile == nil) {
		t.Errorf("Parse Config should return nil")
	}
	if err != nil {
		t.Errorf("Parse config method should not return error!")
	}
}

func TestParseConfig_WhenInvalidConfigFilePath_ReturnsInvalidFileError(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "test.txt"
	expectedInvalidFileError := "Invalid config file. Expecting a \"yaml\" or \"json\" extension test.txt"

	//When
	_, actualInvalidFileError := ParseConfig(MOCK_CONFIG_FILE_PATH)

	//Then
	if actualInvalidFileError.Error() != expectedInvalidFileError {
		t.Errorf("Invalid file path should return error!")
	}
}

func TestGetConfigFilePath_WhenConfigFilePathNotExist_ReturnsEmptyString(t *testing.T) {
	//Given
	utils.CreateFolder("test")
	expectedConfigPath := ""

	//when
	actualConfigPath, err := GetConfigFilePath("test", "", false)

	//Then
	if !reflect.DeepEqual(actualConfigPath, expectedConfigPath) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualConfigPath, expectedConfigPath)
	}
	if err != nil {
		t.Errorf("Parse config method should not return error")
	}
	utils.DeleteFolder("test")
}

func TestGetConfigFilePath_WhenInvalidRootDirectory_ReturnsInvalidDirectoryError(t *testing.T) {
	//Given
	expectedInvalidDirectoryError := "Error fetching path:"

	//When
	_, actualInvalidDirectoryError := GetConfigFilePath("", "", true)
	actualErrorMessage := actualInvalidDirectoryError.Error()

	//Then
	if !strings.Contains(actualErrorMessage, expectedInvalidDirectoryError) {
		t.Errorf("Get config file error mismatched!")
	}
}

func TestGetConfigFilePath_WhenConfigFilePathExist_ReturnsConfigFilePath(t *testing.T) {
	//Given
	expectedConfigPath := "test/.asist.yaml"
	utils.CreateFolder("test")
	utils.CreateFile(expectedConfigPath)

	//when
	actualConfigPath, err := GetConfigFilePath("test", expectedConfigPath, false)

	//Then
	if !reflect.DeepEqual(actualConfigPath, expectedConfigPath) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualConfigPath, expectedConfigPath)
	}
	if err != nil {
		t.Errorf("GetConfigFilePath method should not return error")
	}
	utils.DeleteFolder("test")
}

func TestGetConfigFilePath_WhenConfigPathNotProvidedAndYAMLConfigExistInRootDir_ReturnsYamlFilePath(t *testing.T) {
	//Given
	expectedConfigPath := "test/.asist.yaml"
	utils.CreateFolder("test")
	utils.CreateFile(expectedConfigPath)

	//when
	actualConfigPath, err := GetConfigFilePath("test", "", true)

	//Then
	if !reflect.DeepEqual(actualConfigPath, expectedConfigPath) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualConfigPath, expectedConfigPath)
	}
	if err != nil {
		t.Errorf("GetConfigFilePath method should not return error")
	}
	utils.DeleteFolder("test")
}

func TestGetConfigFilePath_TestGetConfigFilePath_WhenConfigPathNotProvidedAndJSONConfigExistInRootDir_ReturnsJSONFilePath(t *testing.T) {
	//Given
	expectedConfigPath := "test/.asist.json"
	utils.CreateFolder("test")
	utils.CreateFile(expectedConfigPath)

	//when
	actualConfigPath, err := GetConfigFilePath("test", "", true)

	//Then
	if !reflect.DeepEqual(actualConfigPath, expectedConfigPath) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualConfigPath, expectedConfigPath)
	}
	if err != nil {
		t.Errorf("GetConfigFilePath method should not return error")
	}
	utils.DeleteFolder("test")
}

func TestFilterExcludedFilesAndFolders_WhenExcludeFileAndFoldersLengthNil_ReturnsAllPaths(t *testing.T) {
	//Given
	filePaths := []string{"testfile1", "testfile2", "testfile3"}
	config := Config{}
	expectedFilePaths := filePaths

	//When
	actualFilePaths := config.FilterExcludedFilesAndFolders(filePaths)

	//Then
	if !reflect.DeepEqual(actualFilePaths, expectedFilePaths) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualFilePaths, expectedFilePaths)
	}
}

func TestFilterExcludedFilesAndFolders_WhenExcludeFileAndFoldersHasExcludedPaths_ReturnsFilteredPaths(t *testing.T) {
	//Given
	filePaths := []string{"testfile1", "testfile2", "testfile3"}
	expectedPaths := []string{"testfile1", "testfile3"}
	config := Config{
		ExcludeFilesAndFolders: []string{"testfile2"},
	}

	//When
	actualPaths := config.FilterExcludedFilesAndFolders(filePaths)

	//Then
	if !reflect.DeepEqual(actualPaths, expectedPaths) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Config file paths are mismatched!", actualPaths, expectedPaths)
	}
}

func TestCICDRules_WhenCICDRulesInConfigFile_ReturnsCICDRulesList(t *testing.T) {
	//Given
	configFile := Config{
		CICDRules: []string{
			"testRuleId",
		},
	}
	expectedResult := []rules.RuleID{"testRuleId"}

	//When
	actualResult := configFile.GetCICDRuleIds()

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "CI/CD rules are mismatched!", actualResult, expectedResult)
	}
}

func TestOverrideRulesIds_WhenOnlyEnabledOverridedRuleIdsExist_ReturnsEnabledRuleIds(t *testing.T) {
	//Given
	ENABLE := true
	configFile := Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"testRuleId": {
				Enabled: &ENABLE,
			},
		},
		CustomRegexRules: map[string]CustomRegexRule{
			"customTestRule": {
				Enabled: &ENABLE,
			},
		},
	}
	ruleIds := []rules.RuleID{"testRuleId"}
	expectedResult := []rules.RuleID{"testRuleId"}

	//When
	actualResult := configFile.GetEnabledOverridedStandardRuleIds(ruleIds)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Overridden rule Ids are mismatched!", actualResult, expectedResult)
	}
}

func TestOverrideRulesIds_WhenEnabledOverrideIdsExistAndEnablesAllStandardRulesIsTrue_ReturnsEnableRuleIds(t *testing.T) {
	//Given
	ENABLE := true
	configFile := Config{
		EnableAllStandardRules: &ENABLE,
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"testRuleId": {
				Enabled: &ENABLE,
			},
		},
		CustomRegexRules: map[string]CustomRegexRule{
			"customTestRule": {
				Enabled: &ENABLE,
			},
		},
	}
	ruleIds := []rules.RuleID{"testRuleId", "testRuleId1"}
	expectedResult := []rules.RuleID{"testRuleId", "testRuleId1"}

	//When
	actualResult := configFile.GetEnabledOverridedStandardRuleIds(ruleIds)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Overridden rule Ids are mismatched!", actualResult, expectedResult)
	}
}

func TestGetCustomRuleIds_WhenCustomRuleIdsExist_ReturnsRuleIds(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.yaml"
	configFile, _ := ParseConfig(MOCK_CONFIG_FILE_PATH)
	expectedCustomRuleIds := []rules.RuleID{
		"CustomRule1",
	}

	//When
	actualCustomRuleIds := configFile.GetCustomRuleIds()

	//Then
	if !reflect.DeepEqual(actualCustomRuleIds, expectedCustomRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom RuleIds are mismatched!", actualCustomRuleIds, expectedCustomRuleIds)
	}
}

func TestGetOverridedRulesId_WhenRuleOverridesEmpty_ReturnsNil(t *testing.T) {
	//Given
	configFile := Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{},
	}

	//When
	actualResult := configFile.GetOverridedRulesId()

	//Then
	if actualResult != nil {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Overrided rule Ids are not nil!!", actualResult, nil)
	}
}

func TestGetOverridedRulesId_WhenRuleOverridesNotEmpty_ReturnsRuleIds(t *testing.T) {
	//Given
	ENABLE := true
	configFile := Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				Severity: "Medium",
				Enabled:  &ENABLE,
			},
		},
	}
	expectedResult := []rules.RuleID{"XSSTooltip"}

	//When
	actualResult := configFile.GetOverridedRulesId()

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Overrided rule Ids are mismatched!!", actualResult, expectedResult)
	}
}

func TestGetEnabledCustomRuleIds_WhenCustomRulesEmpty_ReturnsEmptySlice(t *testing.T) {
	//Given
	configFile := Config{
		CustomRegexRules: map[string]CustomRegexRule{},
	}
	expectedResult := []rules.RuleID{}

	//When
	actualResult := configFile.GetEnabledCustomRuleIds()

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom Regex Rule Ids are not empty!!", actualResult, expectedResult)
	}
}

func TestGetEnabledCustomRuleIds_WhenCustomRulesNotEmpty_ReturnsRuleIds(t *testing.T) {
	//Given
	ENABLE := true
	configFile := Config{
		CustomRegexRules: map[string]CustomRegexRule{
			"customTestRule": {
				Enabled: &ENABLE,
			},
		},
	}
	expectedResult := []rules.RuleID{"customTestRule"}

	//When
	actualResult := configFile.GetEnabledCustomRuleIds()

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom Regex Rule Ids are mismatched!!", actualResult, expectedResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenConfigNil_ReturnsZero(t *testing.T) {
	//Given
	var configFile *Config

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("SomeRule")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenRuleOverridesNil_ReturnsZero(t *testing.T) {
	//Given
	configFile := &Config{
		RuleOverrides: nil,
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("SomeRule")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenRuleNotInOverrides_ReturnsZero(t *testing.T) {
	//Given
	ENABLE := true
	configFile := &Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"OtherRule": {
				Enabled: &ENABLE,
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("SomeRule")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenMaxIssuesNil_ReturnsZero(t *testing.T) {
	//Given
	ENABLE := true
	configFile := &Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				Severity: "Medium",
				Enabled:  &ENABLE,
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("XSSTooltip")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0 when CicdMaxIssues is nil, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenMaxIssuesExplicitlyZero_ReturnsZero(t *testing.T) {
	//Given
	MAX_ISSUES_ZERO := 0
	configFile := &Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				CicdMaxIssues: &MAX_ISSUES_ZERO,
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("XSSTooltip")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0 when CicdMaxIssues is explicitly 0, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenMaxIssuesSet_ReturnsValue(t *testing.T) {
	//Given
	MAX_ISSUES := 50
	configFile := &Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				CicdMaxIssues: &MAX_ISSUES,
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("XSSTooltip")

	//Then
	if actualResult != 50 {
		t.Errorf("Expected 50, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenParsedFromYAML_ReturnsValue(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.yaml"
	configFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("XSSTooltip")

	//Then
	if actualResult != 10 {
		t.Errorf("Expected 10 from YAML config, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenParsedFromJSON_ReturnsValue(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.json"
	configFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("XSSTooltip")

	//Then
	if actualResult != 10 {
		t.Errorf("Expected 10 from JSON config, got %d", actualResult)
	}
}
func TestGetRuleCicdMaxIssues_WhenCustomRuleHasCicdMaxIssues_ReturnsValue(t *testing.T) {
	//Given
	MAX_ISSUES := 25
	configFile := &Config{
		CustomRegexRules: map[string]CustomRegexRule{
			"MyCustomRule": {
				CicdMaxIssues: &MAX_ISSUES,
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("MyCustomRule")

	//Then
	if actualResult != 25 {
		t.Errorf("Expected 25 for custom rule cicdmaxissues, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenCustomRuleHasNoCicdMaxIssues_ReturnsZero(t *testing.T) {
	//Given
	configFile := &Config{
		CustomRegexRules: map[string]CustomRegexRule{
			"MyCustomRule": {
				Name: "MyCustomRule",
			},
		},
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("MyCustomRule")

	//Then
	if actualResult != 0 {
		t.Errorf("Expected 0 when custom rule has no cicdmaxissues, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenCustomRuleParsedFromYAML_ReturnsValue(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.yaml"
	configFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("CustomRule1")

	//Then
	if actualResult != 5 {
		t.Errorf("Expected 5 from YAML custom rule cicdmaxissues, got %d", actualResult)
	}
}

func TestGetRuleCicdMaxIssues_WhenCustomRuleParsedFromJSON_ReturnsValue(t *testing.T) {
	//Given
	const MOCK_CONFIG_FILE_PATH = "./testData/config.json"
	configFile, err := ParseConfig(MOCK_CONFIG_FILE_PATH)
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	//When
	actualResult := configFile.GetRuleCicdMaxIssues("CustomRule1")

	//Then
	if actualResult != 5 {
		t.Errorf("Expected 5 from JSON custom rule cicdmaxissues, got %d", actualResult)
	}
}