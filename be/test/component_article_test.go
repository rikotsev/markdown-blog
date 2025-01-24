package test

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
)

const (
	ArticlePath        = "article"
	CategoryTechnology = "Technology"
	CategoryPhilosophy = "Philosophy"
	CategoryTravel     = "Travel"
)

func dummyArticles(categoryToId map[string]string) []gen.ArticleCreateJSONRequestBody {
	return []gen.ArticleCreateJSONRequestBody{
		{
			Title: "Yet Another JS Framework",
			Category: gen.CategoryRef{
				Id: categoryToId[CategoryTechnology],
			},
			Description: "Discussing another JS framework",
			Content:     "This is a very long content about JS",
		},
		{
			Title: "Morocco - a destination worth visiting",
			Category: gen.CategoryRef{
				Id: categoryToId[CategoryTravel],
			},
			Description: "The Morocco you know nothing about...",
			Content:     "This is a very long content about Morocco",
		},
		{
			Title: "Sisyphus and Albert Camus",
			Category: gen.CategoryRef{
				Id: categoryToId[CategoryPhilosophy],
			},
			Description: "This is a positive essay, trust me",
			Content:     "This is a very long content about the myth",
		},
		{
			Title: "Rome - why it's worth it every time",
			Category: gen.CategoryRef{
				Id: categoryToId[CategoryTravel],
			},
			Description: "History with amazing cuisine and more",
			Content:     "This is a very long content about Rome",
		},
	}
}

func (s *ApplicationSuite) setupCategories() map[string]string {
	categoryToId := map[string]string{}
	result := gen.Category{}
	s.httpPost(CategoryPath, gen.CategoryCore{
		Name: ptr(CategoryTechnology),
	}, &result)
	categoryToId[CategoryTechnology] = result.Id
	s.httpPost(CategoryPath, gen.CategoryCore{
		Name: ptr(CategoryPhilosophy),
	}, &result)
	categoryToId[CategoryPhilosophy] = result.Id
	s.httpPost(CategoryPath, gen.CategoryCore{
		Name: ptr(CategoryTravel),
	}, &result)
	categoryToId[CategoryTravel] = result.Id

	return categoryToId
}

func (s *ApplicationSuite) TestCreateValidArticle() {
	categoryToId := s.setupCategories()
	req := dummyArticles(categoryToId)[0]

	respCreate := s.httpPostRaw(ArticlePath, &req)
	actualUrlId := respCreate.Header.Get("Location")

	s.Require().Equal(201, respCreate.StatusCode)
	s.Require().Equal("yet-another-js-framework", actualUrlId)

	responseGet := gen.ArticleResponseGet{}
	s.httpGet(fmt.Sprintf("%s/%s", ArticlePath, actualUrlId), &responseGet)

	s.Require().Equal(req.Title, responseGet.Data.Title)
	s.Require().Equal(req.Description, responseGet.Data.Description)
	s.Require().Equal(req.Content, responseGet.Data.Content)
	s.Require().Equal(actualUrlId, responseGet.Data.UrlId)
	s.Require().NotNil(responseGet.Data.CreatedAt)
	s.Require().NotNil(responseGet.Data.EditedAt)
	s.Require().Equal(req.Category.Id, responseGet.Data.Category.Id)
	s.Require().NotNil(responseGet.Included)
	s.Require().True(len(*responseGet.Included) > 0)

	includedCategory, err := (*responseGet.Included)[0].AsCategory()
	s.Require().NoError(err)
	s.Require().Equal(CategoryTechnology, includedCategory.Name)
}

func (s *ApplicationSuite) TestGetAllArticlesWithoutCategory() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	s.httpPostRaw(ArticlePath, &dummyArticles[0])
	s.httpPostRaw(ArticlePath, &dummyArticles[1])
	s.httpPostRaw(ArticlePath, &dummyArticles[2])
	s.httpPostRaw(ArticlePath, &dummyArticles[3])

	resp := gen.ArticleResponseList{}
	s.httpGet(ArticlePath, &resp)

	for index, expected := range dummyArticles {
		s.Require().Equal(expected.Title, resp.Data[index].Title)
		s.Require().Equal(expected.Category.Id, resp.Data[index].Category.Id)
		s.Require().Equal(expected.Description, resp.Data[index].Description)
		s.Require().Equal(expected.Content, resp.Data[index].Content)
	}

	s.Require().NotNil(resp.Included)
	s.Require().Equal(3, len(*resp.Included))
}

func (s *ApplicationSuite) TestDeleteArticle() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	resp := s.httpPostRaw(ArticlePath, &dummyArticles[0])
	urlId := resp.Header.Get("Location")

	statusCode := s.httpDelete(fmt.Sprintf("%s/%s", ArticlePath, urlId))
	resp = s.httpGetRaw(fmt.Sprintf("%s/%s", ArticlePath, urlId))

	s.Require().Equal(200, statusCode)
	s.Require().Equal(404, resp.StatusCode)
}

func (s *ApplicationSuite) TestUpdateArticleTitle() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	resp := s.httpPostRaw(ArticlePath, &dummyArticles[0])
	urlId := resp.Header.Get("Location")
	getResp := gen.ArticleResponseGet{}
	articlePathWithId := fmt.Sprintf("%s/%s", ArticlePath, urlId)
	newTitle := "New Title"

	patchResp := s.httpPatchRaw(articlePathWithId, &gen.ArticleCore{
		Title: &newTitle,
	})
	s.httpGet(articlePathWithId, &getResp)

	s.Require().NotNil(getResp.Data)
	s.Require().Equal(200, patchResp.StatusCode)
	s.Require().Equal(urlId, patchResp.Header.Get("Location"))
	s.assertArticle(gen.ArticleCreateJSONBody{
		Title:       newTitle,
		Description: dummyArticles[0].Description,
		Content:     dummyArticles[0].Content,
		Category:    dummyArticles[0].Category,
	}, *getResp.Data)
}

func (s *ApplicationSuite) TestUpdateArticleDescription() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	resp := s.httpPostRaw(ArticlePath, &dummyArticles[0])
	urlId := resp.Header.Get("Location")
	getResp := gen.ArticleResponseGet{}
	articlePathWithId := fmt.Sprintf("%s/%s", ArticlePath, urlId)
	newDescription := "New Description"

	patchResp := s.httpPatchRaw(articlePathWithId, &gen.ArticleCore{
		Description: &newDescription,
	})
	s.httpGet(articlePathWithId, &getResp)

	s.Require().NotNil(getResp.Data)
	s.Require().Equal(200, patchResp.StatusCode)
	s.Require().Equal(urlId, patchResp.Header.Get("Location"))
	s.assertArticle(gen.ArticleCreateJSONBody{
		Title:       dummyArticles[0].Title,
		Description: newDescription,
		Content:     dummyArticles[0].Content,
		Category:    dummyArticles[0].Category,
	}, *getResp.Data)
}

func (s *ApplicationSuite) TestUpdateArticleContent() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	resp := s.httpPostRaw(ArticlePath, &dummyArticles[0])
	urlId := resp.Header.Get("Location")
	getResp := gen.ArticleResponseGet{}
	articlePathWithId := fmt.Sprintf("%s/%s", ArticlePath, urlId)
	newContent := "New Content"

	patchResp := s.httpPatchRaw(articlePathWithId, &gen.ArticleCore{
		Content: &newContent,
	})
	s.httpGet(articlePathWithId, &getResp)

	s.Require().NotNil(getResp.Data)
	s.Require().Equal(200, patchResp.StatusCode)
	s.Require().Equal(urlId, patchResp.Header.Get("Location"))
	s.assertArticle(gen.ArticleCreateJSONBody{
		Title:       dummyArticles[0].Title,
		Description: dummyArticles[0].Description,
		Content:     newContent,
		Category:    dummyArticles[0].Category,
	}, *getResp.Data)
}

func (s *ApplicationSuite) TestUpdateArticleCategory() {
	categoryToId := s.setupCategories()
	dummyArticles := dummyArticles(categoryToId)
	resp := s.httpPostRaw(ArticlePath, &dummyArticles[0])
	urlId := resp.Header.Get("Location")
	getResp := gen.ArticleResponseGet{}
	articlePathWithId := fmt.Sprintf("%s/%s", ArticlePath, urlId)

	patchResp := s.httpPatchRaw(articlePathWithId, &gen.ArticleCore{
		Category: &gen.CategoryRef{Id: categoryToId[CategoryPhilosophy]},
	})
	s.httpGet(articlePathWithId, &getResp)

	s.Require().NotNil(getResp.Data)
	s.Require().Equal(200, patchResp.StatusCode)
	s.Require().Equal(urlId, patchResp.Header.Get("Location"))
	s.assertArticle(gen.ArticleCreateJSONBody{
		Title:       dummyArticles[0].Title,
		Description: dummyArticles[0].Description,
		Content:     dummyArticles[0].Content,
		Category:    gen.CategoryRef{Id: categoryToId[CategoryPhilosophy]},
	}, *getResp.Data)
}

func (s *ApplicationSuite) assertArticle(expected gen.ArticleCreateJSONBody, actual gen.Article) {
	s.Require().Equal(expected.Title, actual.Title)
	s.Require().Equal(expected.Description, actual.Description)
	s.Require().Equal(expected.Content, actual.Content)
	s.Require().Equal(expected.Category.Id, actual.Category.Id)
}
