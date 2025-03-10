package test

import (
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

const (
	PagePath = "page"
)

func (s *ApplicationSuite) dummyPages() []gen.PageCore {
	return []gen.PageCore{
		{
			Title:   ptr("Home"),
			Content: ptr("This is a home page"),
		},
		{
			Title:   ptr("About me"),
			Content: ptr("This is an about me page"),
		},
		{
			Title:   ptr("Contacts"),
			Content: ptr("this is a contacts page"),
		},
	}
}

func (s *ApplicationSuite) TestCreatePage_Valid() {
	payload := s.dummyPages()[0]
	getResult := gen.PageResponseGet{}

	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &getResult)

	s.Require().Equal(http.StatusCreated, response.StatusCode)
	s.Require().NotNil(getResult.Data)
	s.Require().Equal("home", urlId)
	s.Require().Equal(*payload.Title, getResult.Data.Title)
	s.Require().Equal(*payload.Content, getResult.Data.Content)
}

func (s *ApplicationSuite) TestListPages() {
	dummyPages := s.dummyPages()
	s.httpPostRaw(PagePath, dummyPages[0])
	s.httpPostRaw(PagePath, dummyPages[1])
	s.httpPostRaw(PagePath, dummyPages[2])

	var result gen.PageResponseList
	s.httpGet(PagePath, &result)

	s.Require().Equal(3, len(result.Data))
	s.Require().Equal(*dummyPages[0].Title, *result.Data[0].Title)
	s.Require().Equal("home", *result.Data[0].UrlId)
	s.Require().Equal(*dummyPages[1].Title, *result.Data[1].Title)
	s.Require().Equal("about-me", *result.Data[1].UrlId)
	s.Require().Equal(*dummyPages[2].Title, *result.Data[2].Title)
	s.Require().Equal("contacts", *result.Data[2].UrlId)
}

func (s *ApplicationSuite) TestEditPage_Valid_TitleAndContent() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = ptr(*payload.Title + "-modified title")
	payload.Content = ptr(*payload.Content + "-modified content")
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Title, afterEdit.Data.Title)
	s.Require().Equal(*payload.Content, afterEdit.Data.Content)
}

func (s *ApplicationSuite) TestEditPage_Valid_Title() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = ptr(*payload.Title + "-modified title")
	payload.Content = nil
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Title, afterEdit.Data.Title)
	s.Require().Equal(beforeEdit.Data.Content, afterEdit.Data.Content)
}

func (s *ApplicationSuite) TestEditPage_Valid_Content() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = nil
	payload.Content = ptr(*payload.Content + "-modified content")
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Content, afterEdit.Data.Content)
	s.Require().Equal(beforeEdit.Data.Title, afterEdit.Data.Title)
}

func (s *ApplicationSuite) TestEditPage_NotFound() {
	payload := s.dummyPages()[0]

	payload.Title = nil
	payload.Content = ptr(*payload.Content + "-modified content")
	response := s.httpPatchRaw(PagePath+"/not-existing", &payload)

	s.Require().Equal(http.StatusNotFound, response.StatusCode)
}

func (s *ApplicationSuite) TestDeletePage_Existing() {
	payload := s.dummyPages()[0]

	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	statusCode := s.httpDelete(PagePath + "/" + urlId)
	response = s.httpGetRaw(PagePath + "/" + urlId)

	s.Require().Equal(http.StatusOK, statusCode)
	s.Require().Equal(http.StatusNotFound, response.StatusCode)
}

func (s *ApplicationSuite) TestDeletePage_NotFound() {
	statusCode := s.httpDelete(PagePath + "/not-existing")

	s.Require().Equal(http.StatusNotFound, statusCode)
}
