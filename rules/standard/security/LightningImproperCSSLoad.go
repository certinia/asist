package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var LightningImproperCSSLoadRuleID rules.RuleID = "LightningImproperCSSLoad"

type LightningImproperCSSLoadRule struct {
	metadata rules.RuleMetadata
}

func NewLightningImproperCSSLoadRule() *LightningImproperCSSLoadRule {
	return &LightningImproperCSSLoadRule{
		metadata: rules.RuleMetadata{
			ID:             LightningImproperCSSLoadRuleID,
			Name:           "Potential issues with improper CSS load",
			Description:    "Using <link> or <style> tags to load CSS is considered an insecure practice. These tags can reference external or inline resources that contain CSS or JavaScript, and Salesforce’s Lightning Web Security (LWS) security architecture does not control or sanitize these.",
			Severity:       rules.SeverityLow,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.(component|cmp|css|page)$|/(aura|lwc)/.*\\.js$",
			ExcludePattern: "(\\.test\\.js|\\.min\\.js|\\.t\\.js|-min\\.js|-debug\\.js)$",
			Pattern:        "(<\\s*(apex\\s*:\\s*stylesheet[^><]*\\s+value\\s*=\\s*['\"]http|link\\s+([^><]*rel\\s*=\\s*\"stylesheet\"[^><]*href\\s*=\\s*\"http|[^><]*href\\s*=\\s*\"http[^><]*rel\\s*=\\s*\"stylesheet\")))|<\\s*style(\\s+type\\s*=\\s*['\"]text\\/css['\"])?\\s*>|:\\s*url\\(\\s*['\"]*http",
		},
	}
}

func (r *LightningImproperCSSLoadRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *LightningImproperCSSLoadRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
