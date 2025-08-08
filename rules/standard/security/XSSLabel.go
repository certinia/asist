package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSLabelRuleID rules.RuleID = "XSSLabel"

type XSSLabelRule struct {
	metadata rules.RuleMetadata
}

func NewXSSLabelRule() *XSSLabelRule {
	return &XSSLabelRule{
		metadata: rules.RuleMetadata{
			ID:             XSSLabelRuleID,
			Name:           "Potential XSS with $Label",
			Description:    "Use of $Label may introduce an XSS issue. Consider adding the Encode method before $Label, otherwise, ensure that no Label is rendered unescaped. While labels may only be changed by system admins, if unescaped HTML special characters are left in labels unintentionally, they can open an opportunity for attackers to create XSS exploits.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.component$|\\.page$",
			ExcludePattern: "",
			Pattern:        "\\$(?i)Label\\.\\w+",
		},
	}
}

func (r *XSSLabelRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSLabelRule) Run(fileToScan files.File) []rules.Occurrence {
	return findMatchesForXssLabel(fileToScan, &r.metadata)
}

func findMatchesForXssLabel(fileToScan files.File, ruleMetadata *rules.RuleMetadata) []rules.Occurrence {
	var output []rules.Occurrence

	labelRegexp := regexp.MustCompile(ruleMetadata.Pattern)
	mergeFieldRegexp := regexp.MustCompile(`\{![^\}]*`)
	encodeMethodRegexp := regexp.MustCompile(`(JSENCODE|HTMLENCODE|JSINHTMLENCODE|URLFOR)\(`)

	vulnerableLinkTagsRegex := `<\s*((a\s+href)|(apex\s*:\s*outputLink\s*.*value))\s*=\s*"[^\w\/].*?"`
	vulnerableOnEventRegex := `\bon\w+\s*=\s*"(.*?"|.*)`

	extraVulnerableTagsRegexp := []string{vulnerableLinkTagsRegex, vulnerableOnEventRegex}

	for _, line := range findVulnerableLinesBetweenTags(fileToScan, string(ruleMetadata.ID), extraVulnerableTagsRegexp, false) {
		subLineRange1 := 0
		subLineRange2 := 0
		tempLineText := line.LineContent

		if len(line.ColumnRange) != 0 {
			subLineRange1 = line.ColumnRange[0]
			subLineRange2 = line.ColumnRange[1]
			tempLineText = line.LineContent[subLineRange1:subLineRange2]
		}
		for _, occurrance := range mergeFieldRegexp.FindAllStringIndex(tempLineText, -1) {

			tempStr := regexrulehelper.ReplaceEncodePartByHash(tempLineText[occurrance[0]:occurrance[1]], encodeMethodRegexp)
			for _, label := range labelRegexp.FindAllStringIndex(tempStr, -1) {
				labelStartIndex := subLineRange1 + occurrance[0] + label[0]
				labelEndIndex := subLineRange1 + occurrance[0] + label[1]
				output = append(output,
					rules.Occurrence{
						FileName:        fileToScan.FileName,
						LineContent:     line.LineContent,
						LineNumber:      line.LineNumber,
						ColumnRange:     []int{labelStartIndex, labelEndIndex},
						IsFalsePositive: line.IsFalsePositive,
					},
				)
			}
		}

	}
	return output
}
