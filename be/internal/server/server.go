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
	"github.com/rs/cors"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type ApplicationServer struct {
	appCtx       ApplicationContext
	authProvider AuthenticationProvider
	articleHttp  *article.Http
	categoryHttp *category.Http
	Listener     net.Listener
}

var _ gen.ServerInterface = (*ApplicationServer)(nil)

func New(appCtx ApplicationContext, authProvider AuthenticationProvider) (ApplicationServer, error) {
	queryTimeout := 5 * time.Second
	transformer := urlid.NewTransformerWith(urlid.Slug)
	listener, err := net.Listen("tcp", appCtx.Cfg().Server.Address)
	if err != nil {
		return ApplicationServer{}, err
	}

	return ApplicationServer{
		appCtx:       appCtx,
		authProvider: authProvider,
		articleHttp:  article.NewHttp(article.NewService(article.NewRepository(appCtx.Pool(), queryTimeout), transformer, category.NewMapper())),
		categoryHttp: category.NewHttp(category.NewService(category.NewRepository(appCtx.Pool(), queryTimeout), transformer)),
		Listener:     listener,
	}, nil
}

func (a ApplicationServer) Start() error {
	mux := http.NewServeMux()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{a.appCtx.Cfg().Cors.AllowedOrigins},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := gen.HandlerWithOptions(a, gen.StdHTTPServerOptions{
		BaseRouter:       mux,
		ErrorHandlerFunc: handleExpectedError,
		Middlewares: []gen.MiddlewareFunc{
			handleUnexpectedError,
			AuthAsMiddleware(a.authProvider, handleExpectedError),
		},
	})
	slog.Info("Starting server", "address", a.Listener.Addr().String())
	return http.Serve(a.Listener, corsHandler.Handler(handler))
}

var handleUnexpectedError = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("recovered from panic", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func handleExpectedError(w http.ResponseWriter, r *http.Request, err error) {
	defaultErrorMessage := "something.went.wrong"
	defaultErrorDetail := "follow.error.instance.id"
	defaultType := "about:blank"
	errorInstanceId := uuid.New().String()
	status := http.StatusInternalServerError

	problem := gen.Problem{
		Title:           &defaultErrorMessage,
		ErrorInstanceId: &errorInstanceId,
		Instance:        &r.RequestURI,
		Type:            &defaultType,
		Detail:          &defaultErrorDetail,
	}

	if errors.Is(err, category.DuplicateRecordErr) {
		errorMessage := "category.exists.title"
		errorDetail := "category.exists.detail"
		problem.Title = &errorMessage
		problem.Detail = &errorDetail
		status = http.StatusConflict
	}

	if errors.Is(err, ErrUnauthorized) {
		errorMessage := "auth.failed"
		errorDetail := "auth.failed"
		problem.Title = &errorMessage
		problem.Detail = &errorDetail
		status = http.StatusUnauthorized
	}

	slog.Error("server failure when handling request", "err", err, "errorInstanceId", errorInstanceId)

	problem.Status = &status

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(problem)
}

func (a ApplicationServer) ArticleList(w http.ResponseWriter, r *http.Request, params gen.ArticleListParams) {
	err := a.articleHttp.ArticleList(w, r, params, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	err := a.articleHttp.CreateArticle(w, r, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) ArticleDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.ArticleDelete(w, r, urlId, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) ArticleGet(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.GetArticle(w, r, urlId, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) ArticleEdit(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.articleHttp.ArticleEdit(w, r, urlId, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) CategoryList(w http.ResponseWriter, r *http.Request) {
	err := a.categoryHttp.ListCategories(w, r, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) CategoryCreate(w http.ResponseWriter, r *http.Request) {
	err := a.categoryHttp.CreateCategory(w, r, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
	}
}

func (a ApplicationServer) CategoryDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId) {
	err := a.categoryHttp.DeleteCategory(w, r, urlId, context.Background())
	if err != nil {
		handleExpectedError(w, r, err)
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
