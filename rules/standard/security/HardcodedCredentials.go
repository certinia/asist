package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var HardcodedCredentialsRuleID rules.RuleID = "HardcodedCredentials"

type HardcodedCredentialsRule struct {
	metadata rules.RuleMetadata
}

func NewHardcodedCredentialsRule() *HardcodedCredentialsRule {
	return &HardcodedCredentialsRule{
		metadata: rules.RuleMetadata{
			ID:             HardcodedCredentialsRuleID,
			Name:           "Potential Issues with Hardcoded Credential",
			Description:    "Hardcoding credentials in source code is a common and severe security weakness. Consider using appropriate storage mechanisms such as protected custom settings or protected custom metadata, or use dummy credentials in the case of test classes. Also, ensure valid credentials do not remain in source control history, and are rotated regularly according to industry best practices.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$|\\.trigger$",
			ExcludePattern: "",
			Pattern:        "(?i)\\w*(Secret|Credential|Password|Session|Api\\s*Key|(auth|api|access|refresh|session)[-_]*Token|Username|Authori(s|z)+ation\\s*Code)\\w*\\s*=\\s*['\"].*['\"]",
		},
	}
}

func (r *HardcodedCredentialsRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *HardcodedCredentialsRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, true)
}
