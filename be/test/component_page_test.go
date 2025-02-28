package test

import (
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

const (
	PagePath = "page"
)

func (s *ApplicationSuite) TestCreatePage_Valid() {
	payload := gen.PageCore{
		Title:   ptr("Home"),
		Content: ptr("This is a markdown content of a page."),
	}
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
