package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSJavascriptButtonRuleID rules.RuleID = "XSSJavascriptButton"

type XSSJavascriptButtonRule struct {
	metadata rules.RuleMetadata
}

func NewXSSJavascriptButtonRule() *XSSJavascriptButtonRule {
	return &XSSJavascriptButtonRule{
		metadata: rules.RuleMetadata{
			ID:             XSSJavascriptButtonRuleID,
			Name:           "Potential XSS with javascript button",
			Description:    "Use of a JavaScript button may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.webLink-meta\\.xml$|\\.object$",
			ExcludePattern: "",
			Pattern:        "<linkType>javascript</linkType>",
		},
	}
}

func (r *XSSJavascriptButtonRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSJavascriptButtonRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
