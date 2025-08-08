package testrule

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

var mockData []rules.Occurrence

type test struct {
	metadata rules.RuleMetadata
}

func NewTestRule(metadata rules.RuleMetadata) rules.Rule {
	return &test{
		metadata: metadata,
	}
}

func (r *test) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *test) Run(fileToScan files.File) []rules.Occurrence {
	return mockData
}

func SetMockData(data []rules.Occurrence) {
	mockData = data
}
