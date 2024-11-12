package test

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
)

const (
	CategoryPath = "category"
)

func (s *ApplicationSuite) TestCategoryGetWithNoResults() {
	allCategories := api.CategoryListOK{}
	s.httpGet(CategoryPath, &allCategories)

	s.Require().Equal(0, len(allCategories.Categories))
}

func (s *ApplicationSuite) TestCategoryGetWithNewlyCreatedCategory() {
	req := api.CategoryCreateReq{Name: "tech"}
	res := api.Category{}

	responseCode := s.httpPost(CategoryPath, req, &res)

	s.Require().Equal(201, responseCode)
	s.Require().Equal("tech", res.Name.Value)
	s.Require().Equal("tech", res.UrlId.Value)

	allCategories := api.CategoryListOK{}
	s.httpGet(CategoryPath, &allCategories)

	s.Require().Equal(1, len(allCategories.Categories))
}

func (s *ApplicationSuite) TestCreateMultipleCategories() {
	req := api.CategoryCreateReq{Name: "tech"}
	_ = s.httpPost(CategoryPath, req, &api.Category{})
	req.Name = "travel"
	_ = s.httpPost(CategoryPath, req, &api.Category{})
	req.Name = "philosophy"
	_ = s.httpPost(CategoryPath, req, &api.Category{})

	res := api.CategoryListOK{}
	s.httpGet(CategoryPath, &res)

	s.Require().Equal(3, len(res.Categories))
	s.Require().Equal("tech", res.Categories[0].GetName().Value)
	s.Require().Equal("tech", res.Categories[0].GetUrlId().Value)
	s.Require().Equal("travel", res.Categories[1].GetName().Value)
	s.Require().Equal("travel", res.Categories[1].GetUrlId().Value)
	s.Require().Equal("philosophy", res.Categories[2].GetName().Value)
	s.Require().Equal("philosophy", res.Categories[2].GetUrlId().Value)
}

func (s *ApplicationSuite) TestCreateExistingCategory() {
	req := api.CategoryCreateReq{Name: "tech"}
	_ = s.httpPost(CategoryPath, req, &api.Category{})
	res := api.Problem{}
	responseCode := s.httpPost(CategoryPath, req, &res)

	s.Require().Equal(409, responseCode)
	s.Require().Equal(409, res.Status.Value)
	s.Require().Equal("category.exists.title", res.Title.Value)
	s.Require().NotNil(res.Detail.Value)
}

func (s *ApplicationSuite) TestDeleteExistingCategory() {
	_ = s.httpPost(CategoryPath, api.CategoryCreateReq{Name: "tech"}, &api.Category{})
	toBeDeleted := api.Category{}
	_ = s.httpPost(CategoryPath, api.CategoryCreateReq{Name: "philosophy"}, &toBeDeleted)

	categoriesBeforeDelete := api.CategoryListOK{}
	s.httpGet(CategoryPath, &categoriesBeforeDelete)
	s.Require().Equal(2, len(categoriesBeforeDelete.Categories))
	s.Require().Equal("tech", categoriesBeforeDelete.Categories[0].Name.Value)
	s.Require().Equal("philosophy", categoriesBeforeDelete.Categories[1].Name.Value)

	res := s.httpDelete(fmt.Sprintf("%s/%s", CategoryPath, toBeDeleted.UrlId.Value))
	s.Require().Equal(200, res)

	categoriesBeforeDelete = api.CategoryListOK{}
	s.httpGet(CategoryPath, &categoriesBeforeDelete)
	s.Require().Equal(1, len(categoriesBeforeDelete.Categories))
	s.Require().Equal("tech", categoriesBeforeDelete.Categories[0].Name.Value)
}

func (s *ApplicationSuite) TestDeleteNonExistingCategory() {
	res := s.httpDelete(fmt.Sprintf("%s/%s", CategoryPath, "tech"))
	//TODO fix when API is updated
	s.Require().Equal(401, res)
}
