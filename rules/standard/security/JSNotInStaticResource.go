package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var JSNotInStaticResourceRuleID rules.RuleID = "JSNotInStaticResource"

type JSNotInStaticResourceRule struct {
	metadata rules.RuleMetadata
}

func NewJSNotInStaticResourceRule() *JSNotInStaticResourceRule {
	return &JSNotInStaticResourceRule{
		metadata: rules.RuleMetadata{
			ID:             JSNotInStaticResourceRuleID,
			Name:           "Potential Issues with JS not in Static Resource",
			Description:    "Externally-hosted JavaScript files with <script> tags can cause security issues if the external source is compromised, considering storing them as static resources. This is known to cause Salesforce Security Review failures.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.js$|\\.component$|\\.cmp$",
			ExcludePattern: "",
			Pattern:        "(?i)((<\\s*script[^><]*\\s+src|<\\s*apex\\s*:\\s*includeScript[^><]*\\s+value)\\s*=\\s*[\"']+)(https?|\\s*{\\s*!\\$CurrentPage\\.parameters\\.\\w+)",
		},
	}
}

func (r *JSNotInStaticResourceRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *JSNotInStaticResourceRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)

}
