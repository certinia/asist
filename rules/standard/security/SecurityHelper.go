package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/utils"
)

var openScriptStyleTagRegexp = regexp.MustCompile(`<\s*(script|style)`)
var closeScriptStyleTagRegexp = regexp.MustCompile(`<\s*/\s*(script|style)\s*>`)
var selfCloseScriptStyleTagRegexp = regexp.MustCompile(`<\s*(script|style).*/\s*>`)

/**
 * findVulnerableLinesBetweenTags - method used to find all the vulnerable lines between (script/style) tag or any vulnerable tag in a file based on rule Id
 */
func findVulnerableLinesBetweenTags(fileToScan files.File, ruleId string, extraVulnerableTags []string, isCommentedLinesIncluded bool) []rules.Occurrence {
	var isInsideScriptOrStyleTag = false
	var hasOpeningScriptOrStyleTagFound = false
	var vulnerableLines = []rules.Occurrence{}
	var columnRange = [2]int{}
	var extraVulnerableTagsPresent = len(extraVulnerableTags)
	var extraVulnerableTagsRegexp *regexp.Regexp

	if extraVulnerableTagsPresent != 0 {
		extraVulnerableTagsRegexp = regexp.MustCompile(utils.CreateRegex(extraVulnerableTags))
	}

	for _, line := range fileToScan.Lines {
		columnRange = [2]int{}
		isInsideScriptOrStyleTag = false
		isFalsePositive := fileToScan.IsLineMarkedFalsePositive(ruleId, line.LineNumber)
		if (!isCommentedLinesIncluded && line.IsCommentedLine) || (isFalsePositive && !options.IsBaselineScan()) {
			continue
		}
		if hasOpeningScriptOrStyleTagFound {
			// Search for closing script or style tag
			if closeScriptStyleTagRegexp.MatchString(line.Text) {
				hasOpeningScriptOrStyleTagFound = false
			}
			isInsideScriptOrStyleTag = true

		} else {
			// Search for opening script or style tag
			if openScriptStyleTagRegexp.MatchString(line.Text) {

				//Search for closing script or style tag at same line
				isScriptOrStyleClosingTag := closeScriptStyleTagRegexp.MatchString(line.Text)
				if isScriptOrStyleClosingTag {
					columnRange[0] = openScriptStyleTagRegexp.FindStringIndex(line.Text)[1]
					columnRange[1] = closeScriptStyleTagRegexp.FindStringIndex(line.Text)[0]
				} else if !isScriptOrStyleClosingTag && !selfCloseScriptStyleTagRegexp.MatchString(line.Text) {
					hasOpeningScriptOrStyleTagFound = true
				}
				isInsideScriptOrStyleTag = true
			}
		}
		//IF line is between script or style tag
		if isInsideScriptOrStyleTag {
			//If column range is present
			if columnRange[0] != columnRange[1] {
				vulnerableLines = append(vulnerableLines, rules.Occurrence{
					LineNumber:      line.LineNumber,
					LineContent:     line.Text,
					ColumnRange:     []int{columnRange[0], columnRange[1]},
					IsFalsePositive: isFalsePositive,
				})
			} else {
				vulnerableLines = append(vulnerableLines, rules.Occurrence{
					LineNumber:      line.LineNumber,
					LineContent:     line.Text,
					IsFalsePositive: isFalsePositive,
				})
			}
		} else if extraVulnerableTagsPresent != 0 {
			for _, extraVulnerableTag := range extraVulnerableTagsRegexp.FindAllStringIndex(line.Text, -1) {
				vulnerableLines = append(vulnerableLines, rules.Occurrence{
					LineNumber:      line.LineNumber,
					LineContent:     line.Text,
					ColumnRange:     extraVulnerableTag,
					IsFalsePositive: isFalsePositive,
				})
			}
		}
	}
	return vulnerableLines
}
