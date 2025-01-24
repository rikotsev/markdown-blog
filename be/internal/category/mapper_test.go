package category

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CategoryTestSuite struct {
	suite.Suite
	mapper *Mapper
}

func (c *CategoryTestSuite) SetupSuite() {
	c.mapper = &Mapper{}
}

func (c *CategoryTestSuite) TestMappingToPersistenceLayer() {
	tests := []struct {
		input  gen.Category
		output Entity
	}{
		{
			input: gen.Category{
				EntityType: gen.CategoryTypeCategory,
				Id:         "technology",
				Name:       "Technology",
				UrlId:      "technology",
			},
			output: Entity{
				Id:    "technology",
				Name:  "Technology",
				UrlId: "technology",
			},
		},
		{
			input: gen.Category{
				EntityType: gen.CategoryTypeCategory,
				Id:         "travel",
				Name:       "Travel",
				UrlId:      "travel",
			},
			output: Entity{
				Id:    "travel",
				Name:  "Travel",
				UrlId: "travel",
			},
		},
	}

	for idx, testCase := range tests {
		c.Run(fmt.Sprintf("Mapping category to persistence layer case[%d]", idx), func() {
			c.Require().Equal(testCase.output, c.mapper.ToPersistenceLayer(testCase.input))
		})
	}
}

func (c *CategoryTestSuite) TestMappingToHttpLayer() {
	tests := []struct {
		input  Entity
		output gen.Category
	}{
		{
			input: Entity{
				Id:    "technology",
				UrlId: "technology",
				Name:  "Technology",
			},
			output: gen.Category{
				EntityType: gen.CategoryTypeCategory,
				Id:         "technology",
				UrlId:      "technology",
				Name:       "Technology",
			},
		},
		{
			input: Entity{
				Id:    "travel",
				UrlId: "travel",
				Name:  "Travel",
			},
			output: gen.Category{
				EntityType: gen.CategoryTypeCategory,
				Id:         "travel",
				UrlId:      "travel",
				Name:       "Travel",
			},
		},
	}

	for idx, testCase := range tests {
		c.Run(fmt.Sprintf("Mapping category to http layer case[%d]", idx), func() {
			c.Require().Equal(testCase.output, c.mapper.ToHttpLayer(testCase.input))
		})
	}
}

func TestCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}
