package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSIsRichTextRuleID rules.RuleID = "XSSIsRichText"

type XSSIsRichTextRule struct {
	metadata rules.RuleMetadata
}

func NewXSSIsRichTextRule() *XSSIsRichTextRule {
	return &XSSIsRichTextRule{
		metadata: rules.RuleMetadata{
			ID:             XSSIsRichTextRuleID,
			Name:           "Potential XSS with isRichText",
			Description:    "Use of isRichText may introduce a chatter injection issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.field-meta\\.xml$|\\.object$",
			ExcludePattern: "",
			Pattern:        "(?i)isRichText",
		},
	}
}

func (r *XSSIsRichTextRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSIsRichTextRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
