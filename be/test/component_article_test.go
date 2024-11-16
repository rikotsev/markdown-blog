package test

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
)

const (
	ArticlePath        = "article"
	CategoryTechnology = "Technology"
	CategoryPhilosophy = "Philosophy"
	CategoryTravel     = "Travel"
)

func (s *ApplicationSuite) setupCategories() {
	s.httpPost(CategoryPath, api.CategoryCreateReq{
		Name: CategoryTechnology,
	}, &api.Category{})
	s.httpPost(CategoryPath, api.CategoryCreateReq{
		Name: CategoryPhilosophy,
	}, &api.Category{})
	s.httpPost(CategoryPath, api.CategoryCreateReq{
		Name: CategoryTravel,
	}, &api.Category{})
}

func (s *ApplicationSuite) TestCreateValidArticle() {
	s.setupCategories()
	req := api.ArticleCreateReq{
		Title: "Another Javascript Framework",
		Category: api.Category{
			Name: api.NewOptString(CategoryTechnology),
		},
		Description: "Exploring yet another JS framework!",
		Content:     "This is a long and cumbersome content with complaints about JavaScript.",
	}

	resp := s.httpPostRaw(ArticlePath, &req)
	actualUrlId := resp.Header.Get("Location")

	s.Require().Equal(201, resp.StatusCode)
	s.Require().Equal("another-javascript-framework", actualUrlId)

	actualArticle := api.Article{}
	s.httpGet(fmt.Sprintf("%s/%s", ArticlePath, actualUrlId), &actualArticle)

	s.Require().Equal(req.Title, actualArticle.Title.Value)
	s.Require().Equal(req.Description, actualArticle.Description.Value)
	s.Require().Equal(req.Content, actualArticle.Content.Value)
	s.Require().Equal(actualUrlId, actualArticle.UrlId.Value)
	s.Require().True(actualArticle.ID.Set)
	s.Require().True(actualArticle.CreatedAt.Set)
	s.Require().True(actualArticle.EditedAt.Set)
}
