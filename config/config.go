package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"

	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/message"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/utils"
)

const (
	YAML_CONFIG_FILE_PATH string = "/.asist.yaml"
	JSON_CONFIG_FILE_PATH string = "/.asist.json"
)

type Config struct {
	EnableAllStandardRules *bool
	DontGitIgnore          bool
	DontForceIgnore        bool
	ExcludeFilesAndFolders []string
	RuleOverrides          map[string]rules.RuleMetadataOverride
	CustomRegexRules       map[string]CustomRegexRule
	CICDRules              []string
}

type CustomRegexRule struct {
	Name           string
	Description    string
	Severity       string
	RuleCategory   string
	Enabled        *bool
	Pattern        string
	IncludePattern string
	ExcludePattern string
}

var config *Config

func GetConfigInstance() *Config {
	return config
}

/**
 * ParseConfig - This method reads and parses a configuration file available in either YAML or JSON format.
 */
func ParseConfig(path string) (*Config, error) {
	if path == "" {
		return nil, nil
	}
	fileExt := filepath.Ext(path)

	switch fileExt {
	case ".json":
		return config, parseJSON(path)
	case ".yaml":
		return config, parseYAML(path)
	default:
		return nil, errorhandler.NewUserError(message.GetInvalidConfigFileError(path))
	}
}

/**
*	GetConfigFilePath - Returns the file path of the config file (JSON or YAML) if it exists, otherwise an empty string.
 */
func GetConfigFilePath(rootPath string, configFilePath string, isBaselineEnabled bool) (string, error) {
	if configFilePath != "" || !isBaselineEnabled {
		return configFilePath, nil
	}
	isDir, err := utils.IsDirectory(rootPath)
	if err != nil {
		return "", err
	}
	if isDir {
		if utils.IsFileExists(rootPath + YAML_CONFIG_FILE_PATH) {
			configFilePath = rootPath + YAML_CONFIG_FILE_PATH
		} else if utils.IsFileExists(rootPath + JSON_CONFIG_FILE_PATH) {
			configFilePath = rootPath + JSON_CONFIG_FILE_PATH
		}
	}
	return configFilePath, nil
}

/**
*	FilterExcludedFilesAndFolders - Filter out the paths from the provided list that are present in the ExcludeFilesAndFolders property of config file.
 */
func (c *Config) FilterExcludedFilesAndFolders(paths []string) []string {
	var filteredPaths []string
	if c == nil || len(c.ExcludeFilesAndFolders) == 0 {
		return paths
	}

	patternToExcludeRegex := regexp.MustCompile(utils.CreateRegex(c.ExcludeFilesAndFolders))
	for _, path := range paths {
		if !patternToExcludeRegex.MatchString(filepath.ToSlash(path)) {
			filteredPaths = append(filteredPaths, path)
		}
	}
	return filteredPaths
}

/**
 * GetCustomRuleIds - Returns a list of all custom rule IDs defined in the configuration file.
 */
func (c *Config) GetCustomRuleIds() []rules.RuleID {
	customRuleIds := []rules.RuleID{}
	if c != nil {
		for customRuleId := range c.CustomRegexRules {
			customRuleIds = append(customRuleIds, rules.RuleID(customRuleId))
		}
	}
	return customRuleIds
}

/**
 * GetCICDRuleIds - Returns a list of all ci/cd rule IDs defined in the configuration file.
 */
func (c *Config) GetCICDRuleIds() []rules.RuleID {
	ruleIds := []rules.RuleID{}
	for _, ruleId := range c.CICDRules {
		ruleIds = append(ruleIds, rules.RuleID(ruleId))
	}
	return ruleIds
}

/**
* GetEnabledOverridedStandardRuleIds - Return the standard rule Ids
*	Where RuleOveride property is defined and override rules enabled property is set to true or nil
 */
func (c *Config) GetEnabledOverridedStandardRuleIds(standarRuleIds []rules.RuleID) []rules.RuleID {
	enabledRuleIds := []rules.RuleID{}
	for _, ruleId := range standarRuleIds {
		override, ok := c.RuleOverrides[string(ruleId)]
		if ok {
			if override.Enabled == nil || *override.Enabled {
				enabledRuleIds = append(enabledRuleIds, ruleId)
			}
		} else if c.EnableAllStandardRules == nil || *c.EnableAllStandardRules {
			enabledRuleIds = append(enabledRuleIds, ruleId)
		}
	}

	return enabledRuleIds
}

/**
* GetOverridedRulesId - Method returns the overrided rules Id
 */
func (c *Config) GetOverridedRulesId() []rules.RuleID {
	var ruleIds []rules.RuleID
	if len(c.RuleOverrides) == 0 {
		return nil
	}
	for key := range c.RuleOverrides {
		ruleIds = append(ruleIds, rules.RuleID(key))
	}
	return ruleIds
}

/**
* GetEnabledCustomRuleIds - Return custom ruleIds where enabled property is set to true or nil
 */
func (c *Config) GetEnabledCustomRuleIds() []rules.RuleID {
	enabledRuleIds := []rules.RuleID{}
	if len(c.CustomRegexRules) == 0 {
		return enabledRuleIds
	}
	for customRuleId, customRule := range c.CustomRegexRules {
		if customRule.Enabled == nil || *customRule.Enabled {
			enabledRuleIds = append(enabledRuleIds, rules.RuleID(customRuleId))
		}
	}
	return enabledRuleIds
}

/**
 * parseJSON - Reads and parses the provided JSON configuration file into the corresponding Config structure.
 */
func parseJSON(path string) error {
	fileContent, fileError := readFile(path)
	if fileError != nil {
		return errorhandler.NewInternalError(message.GetInvalidTemplateFileError(fileError))
	}

	fileUnmarshalError := json.Unmarshal(fileContent, &config)
	if fileUnmarshalError != nil {
		return errorhandler.NewInternalError(message.GetFileUnmarshalingError(fileUnmarshalError))
	}
	return nil
}

/**
 * parseYAML - Reads and parses the provided YAML configuration file into the corresponding Config structure.
 */
func parseYAML(path string) error {
	fileContent, fileError := readFile(path)
	if fileError != nil {
		return errorhandler.NewInternalError(message.GetInvalidTemplateFileError(fileError))
	}

	fileUnmarshalError := yaml.Unmarshal(fileContent, &config)
	if fileUnmarshalError != nil {
		return errorhandler.NewInternalError(message.GetFileUnmarshalingError(fileUnmarshalError))
	}
	return nil
}

func readFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	return byteValue, nil
}
