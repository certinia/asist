package finding

import (
	"crypto/sha256"
	"fmt"

	"github.com/certinia/asist/rules"
)

type RecordType string

const (
	BaselineFinding RecordType = "Finding"
	BaselineConfig  RecordType = "Config"
)

type BaselineOutputContent struct {
	FindingID       string             `json:"FindingID"`
	IsCustom        bool               `json:"IsCustom"`
	IsFalsePositive bool               `json:"IsFalsePositive"`
	Id              rules.RuleID       `json:"RuleID"`
	Severity        rules.Severity     `json:"Severity"`
	RuleCategory    rules.RuleCategory `json:"RuleCategory"`
}

type BaselineOutput struct {
	RepositoryName string      `json:"RepositoryName"`
	RepositoryURL  string      `json:"RepositoryURL"`
	RecordType     RecordType  `json:"RecordType"`
	Content        interface{} `json:"Content"`
}

type Finding struct {
	ID           rules.RuleID       `json:"ID"`
	Name         string             `json:"Name"`
	Description  string             `json:"Description"`
	Severity     rules.Severity     `json:"Severity"`
	RuleCategory rules.RuleCategory `json:"RuleCategory"`
	Occurrence   rules.Occurrence   `json:"Occurrence"`
}

/**
 * CreateFindingID - method used to create a unique finding ID for a particular occurrence
 */
func (finding *Finding) CreateFindingID() string {
	contentToHash := fmt.Sprintf("%s-%s", finding.ID, finding.Occurrence.CreateHashableString())

	hash := sha256.New()
	hash.Write([]byte(contentToHash))
	bs := hash.Sum(nil)[0:11]
	return fmt.Sprintf("%x", bs)
}

type Output struct {
	Count           int       `json:"Count"`
	ScanStartedTime string    `json:"Started"`
	ScanEndingTime  string    `json:"Ended"`
	Results         []Finding `json:"Result"`
}
