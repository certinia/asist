package scanner

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/debugger"
	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/finding"
	"github.com/certinia/asist/message"
	"github.com/certinia/asist/output"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/ruleset"
)

var Version = "0.0.0"

/**
*	LoadResources - Method will load and setup the required resources for scan
 */
func LoadResources() ([]string, []*rules.Rule, error) {
	opts := options.Initilize()
	debugger.Debug("start")
	output.DisplayVersion(Version)

	configFile, configErr := loadConfigFile(opts)
	if configErr != nil {
		return nil, nil, configErr
	}
	rules, rulesErr := loadRules(opts, configFile)
	if rulesErr != nil {
		return nil, nil, rulesErr
	}
	// List rules and exit if requested
	output.ListRules(rules)

	paths, pathsErr := loadAndFilterFilePath(configFile)
	if pathsErr != nil {
		return nil, nil, pathsErr
	}
	return paths, rules, nil
}

/**
 * RunRulesOnFiles - method used to run active rules (standard or custom) on the file paths provided by user.
 */
func RunRulesOnFiles(filePaths []string, rules []*rules.Rule) (*finding.Output, error) {
	var finalResult finding.Output
	allfindings := []finding.Finding{}
	for _, path := range filePaths {
		debugger.Debug(fmt.Sprintf("checking if eligible to scan file %s", path))
		rulesToRun := getValidRulesForFile(path, rules)
		if len(rulesToRun) == 0 {
			debugger.Debug(fmt.Sprintf("file is not eligible to scan for enabled rules %s", path))
			continue
		}
		debugger.Debug(fmt.Sprintf("file is eligible to scan for enabled rules %s", path))
		findings, err := runRulesOnFile(rulesToRun, path)
		if err != nil {
			return nil, err
		}
		allfindings = append(allfindings, findings...)
	}
	finalResult.Count = len(allfindings)
	finalResult.Results = allfindings
	return &finalResult, nil
}

/**
 * runRulesOnFile - method used to scan a particular file using the enabled rules and
 *	returns finding for particular file
 */
func runRulesOnFile(rulesToRun []*rules.Rule, fileName string) ([]finding.Finding, error) {
	debugger.Debug(fmt.Sprintf("running rules on %s", fileName))
	allFindings := []finding.Finding{}
	fileMaster, err := files.Read(fileName)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, errorhandler.NewInternalError(message.GetFileReadError(fileName, err))
	}
	debugger.Debug(fmt.Sprintf("read file %s into memory", fileName))

	for _, rule := range rulesToRun {
		ruleMetadata := (*rule).GetMetadata()
		//Search result in master file using pattern(Regex)
		currentRuleOccurrence := (*rule).Run(*fileMaster)
		for _, occurrence := range currentRuleOccurrence {
			allFindings = append(allFindings, finding.Finding{
				Occurrence:   occurrence,
				ID:           ruleMetadata.ID,
				Name:         ruleMetadata.Name,
				Description:  ruleMetadata.Description,
				Severity:     ruleMetadata.Severity,
				RuleCategory: ruleMetadata.RuleCategory,
			})
		}
		debugger.Debug(fmt.Sprintf("ran rule %s on %s", ruleMetadata.ID, fileName))
	}

	return allFindings, nil
}

/**
* getValidRulesForFile - Method will return rules which are valid for file.
 */
func getValidRulesForFile(path string, ruleInstances []*rules.Rule) []*rules.Rule {
	rulesToRun := []*rules.Rule{}
	for _, rule := range ruleInstances {
		metadata := (*rule).GetMetadata()
		if regexrulehelper.RunIncludeExcludePatternsOnFile(path, *metadata) {
			rulesToRun = append(rulesToRun, rule)
		}
	}
	return rulesToRun
}

func loadConfigFile(opts *options.Options) (*config.Config, error) {
	var err error
	//Get config file path
	opts.ConfigFile, err = config.GetConfigFilePath(options.GetPathToScan(), opts.ConfigFile, opts.BaselineScan)
	if err != nil {
		return nil, err
	}
	// Parse config file
	configFile, configErr := config.ParseConfig(opts.ConfigFile)
	if configErr != nil {
		return nil, configErr
	}
	debugger.Debug("parsed config")
	return configFile, nil
}

func loadRules(opts *options.Options, configFile *config.Config) ([]*rules.Rule, error) {
	standardRuleIds, customRuleIds, ruleIdsErr := ruleset.GetRuleIdsToRun(configFile, opts)
	if ruleIdsErr != nil {
		return nil, ruleIdsErr
	}
	debugger.Debug("created list of rules to run")

	//Create rules for Ids
	rules, ruleErr := ruleset.CreateAndOverrideRules(standardRuleIds, customRuleIds, configFile)
	if ruleErr != nil {
		return nil, ruleErr
	}
	debugger.Debug("created rules")
	return rules, nil
}

func loadAndFilterFilePath(configFile *config.Config) ([]string, error) {
	fileOptions := files.FileOptions{
		RootPath:        options.GetPathToScan(),
		DontForceIgnore: configFile != nil && configFile.DontForceIgnore,
		DontGitIgnore:   configFile != nil && configFile.DontGitIgnore,
	}
	// Get all file paths to scan
	paths, pathErr := files.GetAllFilePaths(fileOptions)
	if pathErr != nil {
		return nil, pathErr
	}
	// Excludes files and folders from user provided directory using yaml feature 'excludefilesandfolders'
	paths = configFile.FilterExcludedFilesAndFolders(paths)
	debugger.Debug("enumerated files to scan")
	return paths, nil
}
