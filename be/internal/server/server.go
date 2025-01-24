package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/rikotsev/markdown-blog/be/gen"
	"github.com/rikotsev/markdown-blog/be/internal/article"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type ApplicationServer struct {
	appCtx       ApplicationContext
	articleHttp  *article.Http
	categoryHttp *category.Http
	Listener     net.Listener
}

var _ gen.ServerInterface = (*ApplicationServer)(nil)

func New(appCtx ApplicationContext) (ApplicationServer, error) {
	queryTimeout := 5 * time.Second
	transformer := urlid.NewTransformerWith(urlid.Slug)
	listener, err := net.Listen("tcp", appCtx.Cfg().Server.Address)
	if err != nil {
		return ApplicationServer{}, err
	}

	return ApplicationServer{
		appCtx:       appCtx,
		articleHttp:  article.NewHttp(article.NewService(article.NewRepository(appCtx.Pool(), queryTimeout), transformer, category.NewMapper())),
		categoryHttp: category.NewHttp(category.NewService(category.NewRepository(appCtx.Pool(), queryTimeout), transformer)),
		Listener:     listener,
	}, nil
}

func (a ApplicationServer) Start() error {
	mux := http.NewServeMux()
	handler := gen.HandlerFromMux(a, mux)
	return http.Serve(a.Listener, handler)
}

func handleError(w http.ResponseWriter, r *http.Request, err error, code int) {
	defaultErrorMessage := "something.went.wrong"
	defaultErrorDetail := "follow.error.instance.id"
	defaultType := "about:blank"
	errorInstanceId := uuid.New().String()

	problem := gen.Problem{
		Title:           &defaultErrorMessage,
		ErrorInstanceId: &errorInstanceId,
		Instance:        &r.RequestURI,
		Status:          &code,
		Type:            &defaultType,
		Detail:          &defaultErrorDetail,
	}

	if errors.Is(err, category.DuplicateRecordErr) {
		errorMessage := "category.exists.title"
		errorDetail := "category.exists.detail"
		problem.Title = &errorMessage
		problem.Detail = &errorDetail
		code = 409
	}

	slog.Error("server failure when handling request", "err", err, "errorInstanceId", errorInstanceId)

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(problem)
}

func (a ApplicationServer) ArticleList(w http.ResponseWriter, r *http.Request, params gen.ArticleListParams) {
	err := a.articleHttp.ArticleList(w, r, params, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	err := a.articleHttp.CreateArticle(w, r, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) ArticleDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.ArticleDelete(w, r, urlId, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) ArticleGet(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.GetArticle(w, r, urlId, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) ArticleEdit(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.ArticleEdit(w, r, urlId, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) CategoryList(w http.ResponseWriter, r *http.Request) {
	err := a.categoryHttp.ListCategories(w, r, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) CategoryCreate(w http.ResponseWriter, r *http.Request) {
	err := a.categoryHttp.CreateCategory(w, r, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) CategoryDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.categoryHttp.DeleteCategory(w, r, urlId, context.Background())
	if err != nil {
		handleError(w, r, err, http.StatusInternalServerError)
	}
}

func (a ApplicationServer) PageList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationServer) PageCreate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationServer) PageDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationServer) PageGet(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationServer) PageEdit(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	//TODO implement me
	panic("implement me")
}
