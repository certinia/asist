package test

import (
	"path/filepath"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/finding"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/ruleset"
)

var filePaths []string
var standardRuleIDs []rules.RuleID
var customRuleIDs []rules.RuleID
var configFile *config.Config
var ruleInstances []*rules.Rule

func GetAbsPath(relPath string) string {
	absPath, _ := filepath.Abs(relPath)
	return absPath
}

// Initialise the data for each function
func createData(path string, standardRuleID rules.RuleID, configPath string) {
	fileOptions := files.FileOptions{
		RootPath:        path,
		DontGitIgnore:   true,
		DontForceIgnore: true,
	}
	filePaths, _ = files.GetAllFilePaths(fileOptions)
	standardRuleIDs = []rules.RuleID{standardRuleID}
	customRuleIDs = []rules.RuleID{}
	configFile, _ = config.ParseConfig(configPath)
	ruleInstances, _ = ruleset.CreateAndOverrideRules(standardRuleIDs, customRuleIDs, configFile)
}

type PartialFinding struct {
	ID         rules.RuleID
	Occurrence PartialOccurrence
}

type PartialOccurrence struct {
	FileName    string `json:"File"`
	LineNumber  int    `json:"LineNumber"`
	ColumnRange []int  `json:"ColumnRange"`
}

type PartialOutput struct {
	Count   int
	Results []PartialFinding
}

func projectOutputToPartial(out finding.Output) PartialOutput {
	partialFindings := make([]PartialFinding, len(out.Results))
	for index, finding := range out.Results {
		partialFindings[index] = PartialFinding{
			ID: finding.ID,
			Occurrence: PartialOccurrence{
				FileName:    finding.Occurrence.FileName,
				LineNumber:  finding.Occurrence.LineNumber,
				ColumnRange: finding.Occurrence.ColumnRange,
			},
		}
	}
	return PartialOutput{
		Count:   out.Count,
		Results: partialFindings,
	}
}
