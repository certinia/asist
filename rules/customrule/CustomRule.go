package customrule

import (
	"github.com/certinia/asist/config"
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

type CustomRule struct {
	metadata rules.RuleMetadata
}

/**
 * NewCustomRule - method used create a custom rule
 */
func NewCustomRule(customRule config.CustomRegexRule, customRuleID rules.RuleID) *CustomRule {
	return &CustomRule{
		metadata: rules.RuleMetadata{
			ID:             customRuleID,
			Name:           customRule.Name,
			Description:    customRule.Description,
			Severity:       rules.Severity(customRule.Severity),
			RuleCategory:   rules.RuleCategory(customRule.RuleCategory),
			IncludePattern: customRule.IncludePattern,
			ExcludePattern: customRule.ExcludePattern,
			Pattern:        customRule.Pattern,
		},
	}
}

func (r *CustomRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

/**
 * Run - method used to run a custom rule
 */
func (r *CustomRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
