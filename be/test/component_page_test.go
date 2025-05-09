package test

import (
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

const (
	PagePath               = "page"
	modifiedContentPostfix = "-modified content"
	modifiedTitlePostfix   = "-modified title"
)

func (s *ApplicationSuite) dummyPages() []gen.PageCore {
	return []gen.PageCore{
		{
			Title:    ptr("Home"),
			Content:  ptr("This is a home page"),
			Position: ptr(100),
		},
		{
			Title:    ptr("About me"),
			Content:  ptr("This is an about me page"),
			Position: ptr(200),
		},
		{
			Title:    ptr("Contacts"),
			Content:  ptr("this is a contacts page"),
			Position: ptr(300),
		},
	}
}

func (s *ApplicationSuite) TestCreatePageValid() {
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
	s.Require().Equal(*dummyPages[0].Title, result.Data[0].Title)
	s.Require().Equal("home", result.Data[0].UrlId)
	s.Require().Equal(*dummyPages[0].Position, result.Data[0].Position)
	s.Require().Equal(*dummyPages[1].Title, result.Data[1].Title)
	s.Require().Equal("about-me", result.Data[1].UrlId)
	s.Require().Equal(*dummyPages[1].Position, result.Data[1].Position)
	s.Require().Equal(*dummyPages[2].Title, result.Data[2].Title)
	s.Require().Equal("contacts", result.Data[2].UrlId)
	s.Require().Equal(*dummyPages[2].Position, result.Data[2].Position)
}

func (s *ApplicationSuite) TestEditPageValidTitleAndContent() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = ptr(*payload.Title + modifiedTitlePostfix)
	payload.Content = ptr(*payload.Content + modifiedContentPostfix)
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Title, afterEdit.Data.Title)
	s.Require().Equal(*payload.Content, afterEdit.Data.Content)
}

func (s *ApplicationSuite) TestEditPageValidTitle() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = ptr(*payload.Title + modifiedTitlePostfix)
	payload.Content = nil
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Title, afterEdit.Data.Title)
	s.Require().Equal(beforeEdit.Data.Content, afterEdit.Data.Content)
}

func (s *ApplicationSuite) TestEditPageValidContent() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = nil
	payload.Content = ptr(*payload.Content + modifiedContentPostfix)
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Content, afterEdit.Data.Content)
	s.Require().Equal(beforeEdit.Data.Title, afterEdit.Data.Title)
}

func (s *ApplicationSuite) TestEditPageValidPosition() {
	beforeEdit := gen.PageResponseGet{}
	afterEdit := gen.PageResponseGet{}
	payload := s.dummyPages()[0]
	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	s.httpGet(PagePath+"/"+urlId, &beforeEdit)

	payload.Title = nil
	payload.Content = ptr(*payload.Content + modifiedContentPostfix)
	response = s.httpPatchRaw(PagePath+"/"+urlId, &payload)
	s.httpGet(PagePath+"/"+urlId, &afterEdit)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	s.Require().Equal(beforeEdit.Data.Id, afterEdit.Data.Id)
	s.Require().Equal(beforeEdit.Data.UrlId, afterEdit.Data.UrlId)
	s.Require().Equal(*payload.Content, afterEdit.Data.Content)
	s.Require().Equal(beforeEdit.Data.Title, afterEdit.Data.Title)
}

func (s *ApplicationSuite) TestEditPageNotFound() {
	payload := s.dummyPages()[0]

	payload.Title = nil
	payload.Content = ptr(*payload.Content + modifiedContentPostfix)
	response := s.httpPatchRaw(PagePath+"/not-existing", &payload)

	s.Require().Equal(http.StatusNotFound, response.StatusCode)
}

func (s *ApplicationSuite) TestDeletePageExisting() {
	payload := s.dummyPages()[0]

	response := s.httpPostRaw(PagePath, payload)
	urlId := response.Header.Get("Location")
	statusCode := s.httpDelete(PagePath + "/" + urlId)
	response = s.httpGetRaw(PagePath + "/" + urlId)

	s.Require().Equal(http.StatusOK, statusCode)
	s.Require().Equal(http.StatusNotFound, response.StatusCode)
}

func (s *ApplicationSuite) TestDeletePageNotFound() {
	statusCode := s.httpDelete(PagePath + "/not-existing")

	s.Require().Equal(http.StatusNotFound, statusCode)
}
