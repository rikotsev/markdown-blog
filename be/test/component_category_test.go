package test

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
)

const (
	CategoryPath = "category"
)

func (s *ApplicationSuite) TestCategoryGetWithNoResults() {
	responseContent := gen.CategoryResponseList{}
	s.httpGet(CategoryPath, &responseContent)

	s.Require().Equal(0, len(responseContent.Data))
}

func (s *ApplicationSuite) TestCategoryGetWithNewlyCreatedCategory() {
	req := gen.CategoryCreateJSONBody{Name: "tech"}
	res := gen.Category{}

	responseCode := s.httpPost(CategoryPath, req, &res)

	s.Require().Equal(201, responseCode)
	s.Require().Equal("tech", res.Name)
	s.Require().Equal("tech", res.UrlId)

	allCategories := gen.CategoryResponseList{}
	s.httpGet(CategoryPath, &allCategories)

	s.Require().Equal(1, len(allCategories.Data))
}

func (s *ApplicationSuite) TestCreateMultipleCategories() {
	req := gen.CategoryCore{Name: ptr("tech")}
	_ = s.httpPost(CategoryPath, req, &gen.Category{})
	req.Name = ptr("travel")
	_ = s.httpPost(CategoryPath, req, &gen.Category{})
	req.Name = ptr("philosophy")
	_ = s.httpPost(CategoryPath, req, &gen.Category{})

	res := gen.CategoryResponseList{}
	s.httpGet(CategoryPath, &res)

	s.Require().Equal(3, len(res.Data))
	s.Require().Equal("tech", res.Data[0].Name)
	s.Require().Equal("tech", res.Data[0].UrlId)
	s.Require().Equal("travel", res.Data[1].Name)
	s.Require().Equal("travel", res.Data[1].UrlId)
	s.Require().Equal("philosophy", res.Data[2].Name)
	s.Require().Equal("philosophy", res.Data[2].UrlId)
}

func (s *ApplicationSuite) TestCreateExistingCategory() {
	req := gen.CategoryCore{Name: ptr("tech")}
	_ = s.httpPost(CategoryPath, req, &gen.Category{})
	res := gen.Problem{}
	responseCode := s.httpPost(CategoryPath, req, &res)

	s.Require().Equal(409, responseCode)
	s.Require().Equal(409, *res.Status)
	s.Require().Equal("category.exists.title", *res.Title)
	s.Require().Equal("/category", *res.Instance)
	s.Require().NotNil(res.Detail)
	s.Require().NotNil(res.ErrorInstanceId)
}

func (s *ApplicationSuite) TestDeleteExistingCategory() {
	_ = s.httpPost(CategoryPath, gen.CategoryCreateJSONBody{Name: "tech"}, &gen.Category{})
	toBeDeleted := gen.Category{}
	_ = s.httpPost(CategoryPath, gen.CategoryCreateJSONBody{Name: "philosophy"}, &toBeDeleted)

	categoriesBeforeDelete := gen.CategoryResponseList{}
	s.httpGet(CategoryPath, &categoriesBeforeDelete)
	s.Require().Equal(2, len(categoriesBeforeDelete.Data))
	s.Require().Equal("tech", categoriesBeforeDelete.Data[0].Name)
	s.Require().Equal("philosophy", categoriesBeforeDelete.Data[1].Name)

	res := s.httpDelete(fmt.Sprintf("%s/%s", CategoryPath, toBeDeleted.UrlId))
	s.Require().Equal(200, res)

	categoriesBeforeDelete = gen.CategoryResponseList{}
	s.httpGet(CategoryPath, &categoriesBeforeDelete)
	s.Require().Equal(1, len(categoriesBeforeDelete.Data))
	s.Require().Equal("tech", categoriesBeforeDelete.Data[0].Name)
}

func (s *ApplicationSuite) TestDeleteNonExistingCategory() {
	res := s.httpDelete(fmt.Sprintf("%s/%s", CategoryPath, "tech"))
	s.Require().Equal(404, res)
}
