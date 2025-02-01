package test

import (
	"encoding/json"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

func (s *ApplicationSuite) TestUnauthorizedEndpoints() {
	tests := []struct {
		endpoint       string
		request        func() *http.Response
		httpStatusCode int
		problem        gen.Problem
	}{
		{
			endpoint: "POST /category",
			request: func() *http.Response {
				req := gen.CategoryCore{Name: ptr("tech")}
				return s.httpPostRawWithHeaders(CategoryPath, req, map[string]string{})
			},
			httpStatusCode: http.StatusUnauthorized,
			problem: gen.Problem{
				Title:    ptr("auth.failed"),
				Status:   ptr(http.StatusUnauthorized),
				Detail:   ptr("auth.failed"),
				Instance: ptr("/" + CategoryPath),
			},
		},
		{
			endpoint: "DELETE /category/{urlId}",
			request: func() *http.Response {
				path := fmt.Sprintf("%s/%s", CategoryPath, "tech")
				return s.httpDeleteWithHeaders(path, map[string]string{})
			},
			httpStatusCode: http.StatusUnauthorized,
			problem: gen.Problem{
				Title:    ptr("auth.failed"),
				Status:   ptr(http.StatusUnauthorized),
				Detail:   ptr("auth.failed"),
				Instance: ptr(fmt.Sprintf("/%s/%s", CategoryPath, "tech")),
			},
		},
		{
			endpoint: "POST /article",
			request: func() *http.Response {
				req := gen.ArticleCore{}
				return s.httpPostRawWithHeaders(ArticlePath, req, map[string]string{})
			},
			httpStatusCode: http.StatusUnauthorized,
			problem: gen.Problem{
				Title:    ptr("auth.failed"),
				Status:   ptr(http.StatusUnauthorized),
				Detail:   ptr("auth.failed"),
				Instance: ptr("/" + ArticlePath),
			},
		},
		{
			endpoint: "DELETE /article/{urlId}",
			request: func() *http.Response {
				return s.httpDeleteWithHeaders(fmt.Sprintf("%s/%s", ArticlePath, "yet-another-js-framework"), map[string]string{})
			},
			httpStatusCode: http.StatusUnauthorized,
			problem: gen.Problem{
				Title:    ptr("auth.failed"),
				Status:   ptr(http.StatusUnauthorized),
				Detail:   ptr("auth.failed"),
				Instance: ptr(fmt.Sprintf("/%s/%s", ArticlePath, "yet-another-js-framework")),
			},
		},
		{
			endpoint: "PATCH /article/{urlId}",
			request: func() *http.Response {
				req := gen.ArticleCore{}
				return s.httpPatchRawWithHeaders(fmt.Sprintf("%s/%s", ArticlePath, "yet-another-js-framework"), req, map[string]string{})
			},
			httpStatusCode: http.StatusUnauthorized,
			problem: gen.Problem{
				Title:    ptr("auth.failed"),
				Status:   ptr(http.StatusUnauthorized),
				Detail:   ptr("auth.failed"),
				Instance: ptr(fmt.Sprintf("/%s/%s", ArticlePath, "yet-another-js-framework")),
			},
		},
	}

	for idx, testCase := range tests {
		s.Run(fmt.Sprintf("[%d / %d] Testing secure endpoint: %s", idx+1, len(tests), testCase.endpoint), func() {
			resp := testCase.request()
			actual := gen.Problem{}
			err := json.NewDecoder(resp.Body).Decode(&actual)
			s.Require().NoError(err)
			s.Require().Equal(testCase.httpStatusCode, resp.StatusCode)
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
