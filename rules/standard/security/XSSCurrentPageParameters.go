package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSCurrentPageParametersRuleID rules.RuleID = "XSSCurrentPageParameters"

type XSSCurrentPageParametersRule struct {
	metadata rules.RuleMetadata
}

func NewXSSCurrentPageParametersRule() *XSSCurrentPageParametersRule {
	return &XSSCurrentPageParametersRule{
		metadata: rules.RuleMetadata{
			ID:             XSSCurrentPageParametersRuleID,
			Name:           "Potential XSS with CurrentPage.parameters",
			Description:    "Use of CurrentPage.Parameters in script or style context may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.component$",
			ExcludePattern: "",
			Pattern:        "(?i)\\$CurrentPage\\.parameters\\.\\w+",
		},
	}
}

func (r *XSSCurrentPageParametersRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSCurrentPageParametersRule) Run(fileToScan files.File) []rules.Occurrence {
	return findMatchesForCurrentPageParameters(fileToScan, &r.metadata)
}

func findMatchesForCurrentPageParameters(fileToScan files.File, ruleMetadata *rules.RuleMetadata) []rules.Occurrence {
	var output []rules.Occurrence

	currentPageParametersRegexp := regexp.MustCompile(ruleMetadata.Pattern)
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
			for _, currentPageParametersOccurrencesInLine := range currentPageParametersRegexp.FindAllStringIndex(tempStr, -1) {
				currentPageStartIndex := subLineRange1 + occurrance[0] + currentPageParametersOccurrencesInLine[0]
				currentPageEndIndex := subLineRange1 + occurrance[0] + currentPageParametersOccurrencesInLine[1]
				output = append(output,
					rules.Occurrence{
						FileName:        fileToScan.FileName,
						LineContent:     line.LineContent,
						LineNumber:      line.LineNumber,
						ColumnRange:     []int{currentPageStartIndex, currentPageEndIndex},
						IsFalsePositive: line.IsFalsePositive,
					},
				)
			}
		}
	}
	return output
}
