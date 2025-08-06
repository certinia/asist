package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var InsecureCryptoAlgorithmRuleID rules.RuleID = "InsecureCryptoAlgorithm"

type InsecureCryptoAlgorithmRule struct {
	metadata rules.RuleMetadata
}

func NewInsecureCryptoAlgorithmRule() *InsecureCryptoAlgorithmRule {
	return &InsecureCryptoAlgorithmRule{
		metadata: rules.RuleMetadata{
			ID:             InsecureCryptoAlgorithmRuleID,
			Name:           "Potential issues with insecure crypto algorithm.",
			Description:    "Use of weak/broken or custom crypto algorithms such as MD5 or SHA1 may expose data to unauthorized third parties. Consider using a strong and secure crypto algorithm instead.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "(?i)('|\")(MD5|hmacSHA1|RSA-SHA1|SHA1)('|\")",
		},
	}
}

func (r *InsecureCryptoAlgorithmRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *InsecureCryptoAlgorithmRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
