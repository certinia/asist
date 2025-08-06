package files

type Line struct {
	LineNumber      int
	Text            string
	IsCommentedLine bool
}

type File struct {
	Lines           []Line
	FileName        string
	IgnoresSelected []IgnoreSelected
}

type IgnoreSelected struct {
	BeginLine int
	EndLine   int
	RuleIDs   map[string]bool
}

/**
 * IsLineMarkedFalsePositive - method used to check an occurrence is false positive or not for
 * a particular RuleID in a line
 */
func (f *File) IsLineMarkedFalsePositive(ruleID string, lineNumber int) bool {
	for _, ignore := range *&f.IgnoresSelected {
		_, hasRuleFound := ignore.RuleIDs[ruleID]
		if hasRuleFound && ignore.BeginLine <= lineNumber && ignore.EndLine >= lineNumber {
			return true
		}
	}
	return false
}
