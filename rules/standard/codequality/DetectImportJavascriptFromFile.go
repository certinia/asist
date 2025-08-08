package codequality

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var DetectImportJavascriptFromFileRuleID rules.RuleID = "DetectImportJavascriptFromFile"

type DetectImportJavascriptFromFileRule struct {
	metadata rules.RuleMetadata
}

func NewDetectImportJavascriptFromFileRule() *DetectImportJavascriptFromFileRule {
	return &DetectImportJavascriptFromFileRule{
		metadata: rules.RuleMetadata{
			ID:             DetectImportJavascriptFromFileRuleID,
			Name:           "Detect Import statement in LWC with file extension specified",
			Description:    "JavaScript should be imported from folders, not files directly. In the import statement, specify the folder to import from, not the file (don't specify a file extension).\nReference: https://developer.salesforce.com/docs/platform/lwc/guide/js-share-code.html\n\nThis is known to cause Salesforce Security Review failures.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategoryCodeQuality,
			IncludePattern: "/lwc/.*\\.js$",
			ExcludePattern: "/__tests__/",
			Pattern:        "\\bimport(\\s+([\\w\\-\\$]+(,\\s*\\*)?|\\*|([\\w\\-\\$]+,\\s*)?\\{\\s*[\\w\\-\\$,\"'\\s\\/\\*]+\\s*\\},?)(\\s+as\\s+[\\w\\-\\$]+)?\\s+from)?\\s+(\"[^\"]+\\.js\"|'[^']+\\.js')\\s*;",
		},
	}
}

func (r *DetectImportJavascriptFromFileRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *DetectImportJavascriptFromFileRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
