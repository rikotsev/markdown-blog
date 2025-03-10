package test

import (
	"encoding/json"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

func (s *ApplicationSuite) TestUnauthorizedEndpoints() {
	const (
		TechCategoryPath       = CategoryPath + "/tech"
		AnotherJsFrameworkPath = ArticlePath + "/yet-another-js-framework"
		HomePagePath           = PagePath + "/home"
	)

	tests := []struct {
		endpoint string
		request  func() *http.Response
		problem  gen.Problem
	}{
		{
			endpoint: "POST /category",
			request: func() *http.Response {
				req := gen.CategoryCore{Name: ptr("tech")}
				return s.httpPostRawWithHeaders(CategoryPath, req, map[string]string{})
			},
			problem: standardProblem(CategoryPath),
		},
		{
			endpoint: "DELETE /category/{urlId}",
			request: func() *http.Response {
				return s.httpDeleteWithHeaders(TechCategoryPath, map[string]string{})
			},
			problem: standardProblem(TechCategoryPath),
		},
		{
			endpoint: "POST /article",
			request: func() *http.Response {
				req := gen.ArticleCore{}
				return s.httpPostRawWithHeaders(ArticlePath, req, map[string]string{})
			},
			problem: standardProblem(ArticlePath),
		},
		{
			endpoint: "DELETE /article/{urlId}",
			request: func() *http.Response {
				return s.httpDeleteWithHeaders(AnotherJsFrameworkPath, map[string]string{})
			},
			problem: standardProblem(AnotherJsFrameworkPath),
		},
		{
			endpoint: "PATCH /article/{urlId}",
			request: func() *http.Response {
				req := gen.ArticleCore{}
				return s.httpPatchRawWithHeaders(AnotherJsFrameworkPath, req, map[string]string{})
			},
			problem: standardProblem(AnotherJsFrameworkPath),
		},
		{
			endpoint: "POST /page",
			request: func() *http.Response {
				req := gen.PageCreateJSONBody{}
				return s.httpPostRawWithHeaders(PagePath, req, map[string]string{})
			},
			problem: standardProblem(PagePath),
		},
		{
			endpoint: "PATCH /page/{urlId}",
			request: func() *http.Response {
				req := gen.PageCore{}
				return s.httpPatchRawWithHeaders(HomePagePath, req, map[string]string{})
			},
			problem: standardProblem(HomePagePath),
		},
		{
			endpoint: "DELETE /page/{urlId}",
			request: func() *http.Response {
				return s.httpDeleteWithHeaders(HomePagePath, map[string]string{})
			},
			problem: standardProblem(HomePagePath),
		},
	}

	for idx, testCase := range tests {
		s.Run(fmt.Sprintf("[%d / %d] Testing secure endpoint: %s", idx+1, len(tests), testCase.endpoint), func() {
			resp := testCase.request()
			actual := gen.Problem{}
			err := json.NewDecoder(resp.Body).Decode(&actual)
			s.Require().NoError(err)
			s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
			s.Require().NotNil(actual.Title)
			s.Require().Equal(*testCase.problem.Title, *actual.Title)
			s.Require().NotNil(actual.Status)
			s.Require().Equal(*testCase.problem.Status, *actual.Status)
			s.Require().NotNil(actual.Detail)
			s.Require().Equal(*testCase.problem.Detail, *actual.Detail)
			s.Require().NotNil(actual.Instance)
			s.Require().Equal(*testCase.problem.Instance, *actual.Instance)
			s.Require().NotNil(actual.ErrorInstanceId)
		})
	}
}

func standardProblem(customInstanceValue string) gen.Problem {
	return gen.Problem{
		Title:    ptr("auth.failed"),
		Status:   ptr(http.StatusUnauthorized),
		Detail:   ptr("auth.failed"),
		Instance: ptr("/" + customInstanceValue),
	}
}
