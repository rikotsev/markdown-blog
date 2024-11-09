package urlid

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UrlIdSuite struct {
	suite.Suite
	transformerSlug  *Transformer
	transformerNaive *Transformer
}

func (s *UrlIdSuite) SetupSuite() {
	s.transformerSlug = NewTransformerWith(Slug)
	s.transformerNaive = NewTransformerWith(NaiveImpl)
}

func (s *UrlIdSuite) TestProcessWithSlug() {

	testCases := []struct {
		input         string
		expectedSlug  string
		expectedNaive string
	}{
		{
			input:         "A cool new javascript framework!",
			expectedSlug:  "a-cool-new-javascript-framework",
			expectedNaive: "A-cool-new-javascript-framework!",
		},
		{
			input:         "Are you for real?",
			expectedSlug:  "are-you-for-real",
			expectedNaive: "Are-you-for-real?",
		},
		{
			input:         "$$ SignS",
			expectedSlug:  "signs",
			expectedNaive: "$$-SignS",
		},
		{
			input:         "Дори на български?",
			expectedSlug:  "dori-na-b-lgarski",
			expectedNaive: "Дори-на-български?",
		},
	}

	for idx, testCase := range testCases {
		s.Run(fmt.Sprintf("test case [%d]: %s", idx, testCase.input), func() {
			actual := s.transformerSlug.Process(testCase.input)
			s.Equal(testCase.expectedSlug, actual, fmt.Sprintf("slug failed for input: %s", testCase.input))
			actual = s.transformerNaive.Process(testCase.input)
			s.Equal(testCase.expectedNaive, actual, fmt.Sprintf("naive impl failed for input: %s", testCase.input))
		})
	}

}

func TestUrlIdSuite(t *testing.T) {
	suite.Run(t, new(UrlIdSuite))
}
