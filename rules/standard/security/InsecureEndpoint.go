package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
)

var InsecureEndpointRuleID rules.RuleID = "InsecureEndpoint"

type InsecureEndpoint struct {
	metadata rules.RuleMetadata
}

func NewInsecureEndpointRule() *InsecureEndpoint {
	return &InsecureEndpoint{
		metadata: rules.RuleMetadata{
			ID:             InsecureEndpointRuleID,
			Name:           "Insecure endpoint",
			Description:    "Use of the unencrypted HTTP protocol puts data communicated at risk of serious privacy concerns or tampering. Prefer encrypted HTTPS, and leverage https://www.ssllabs.com/ssltest/ or similar to test the encryption quality of the endpoint.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$|\\.page$|\\.component$|\\.email$|\\.cmp$",
			ExcludePattern: "(?i)/((force-app-autotest|autotest|systemtest|test)(s)?|yui|sencha|staticresources|deps)/|Tests?\\.cls$|-meta\\.xml$|\\.min\\.js$",
			Pattern:        "http://",
		},
	}
}

func (r *InsecureEndpoint) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *InsecureEndpoint) Run(fileToScan files.File) []rules.Occurrence {
	return findMatchesForInsecureEndpoint(fileToScan, &r.metadata)

}

func findMatchesForInsecureEndpoint(fileToScan files.File, ruleMetadata *rules.RuleMetadata) []rules.Occurrence {
	var output []rules.Occurrence
	httpRegexp := regexp.MustCompile(ruleMetadata.Pattern)
	xmlnsRegexp := regexp.MustCompile(`\s+xmlns\s*(=|:)`)

	for _, line := range fileToScan.Lines {
		isFalsePositive := fileToScan.IsLineMarkedFalsePositive(string(ruleMetadata.ID), line.LineNumber)
		if line.IsCommentedLine || (isFalsePositive && !options.IsBaselineScan()) {
			continue
		}
		isHttpOccurrenceInLine := httpRegexp.MatchString(line.Text)
		if isHttpOccurrenceInLine {
			isExcludeTrue := xmlnsRegexp.MatchString(line.Text)
			if !isExcludeTrue {
				output = append(
					output,
					rules.Occurrence{
						FileName:        fileToScan.FileName,
						LineNumber:      line.LineNumber,
						LineContent:     line.Text,
						ColumnRange:     []int{0, len(line.Text) - 1},
						IsFalsePositive: isFalsePositive,
					},
				)
			}
		}
	}
	return output
}
