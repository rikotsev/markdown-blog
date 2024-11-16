package server

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rikotsev/markdown-blog/be/gen/api"
	"github.com/rikotsev/markdown-blog/be/internal/article"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
	"log/slog"
	"time"
)

func NewEndpointsHandler(ctx ApplicationContext) (api.Handler, error) {
	queryTimeout := time.Second * 5
	transformer := urlid.NewTransformerWith(urlid.Slug)

	return &endpointsHandler{
		requestTimeout:  time.Second * 3,
		appCtx:          ctx,
		categoryHandler: category.NewHttp(category.NewRepository(ctx.Pool(), queryTimeout), transformer),
		articleHandler:  article.NewHttp(article.NewRepository(ctx.Pool(), queryTimeout), transformer),
	}, nil
}

type endpointsHandler struct {
	requestTimeout  time.Duration
	appCtx          ApplicationContext
	categoryHandler *category.Http
	articleHandler  *article.Http
}

func (handler *endpointsHandler) ArticleCreate(ctx context.Context, req *api.ArticleCreateReq) (api.ArticleCreateRes, error) {
	//ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	//defer cancelFunc()

	return handler.articleHandler.CreateArticle(ctx, req)
}

func (handler *endpointsHandler) ArticleDelete(ctx context.Context, params api.ArticleDeleteParams) (api.ArticleDeleteRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) ArticleEdit(ctx context.Context, req *api.ArticleCore, params api.ArticleEditParams) (api.ArticleEditRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) ArticleGet(ctx context.Context, params api.ArticleGetParams) (api.ArticleGetRes, error) {
	return handler.articleHandler.GetArticle(ctx, params)
}

func (handler *endpointsHandler) ArticleList(ctx context.Context, params api.ArticleListParams) (*api.ArticleListOK, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) CategoryCreate(ctx context.Context, req *api.CategoryCreateReq) (api.CategoryCreateRes, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	defer cancelFunc()

	newCategory, err := handler.categoryHandler.CreateCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (handler *endpointsHandler) CategoryDelete(ctx context.Context, params api.CategoryDeleteParams) (api.CategoryDeleteRes, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	defer cancelFunc()

	res, err := handler.categoryHandler.DeleteCategory(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (handler *endpointsHandler) CategoryList(ctx context.Context) (*api.CategoryListOK, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	defer cancelFunc()
	categories, err := handler.categoryHandler.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &api.CategoryListOK{
		Categories: categories,
	}, nil
}

func (handler *endpointsHandler) PageCreate(ctx context.Context, req *api.PageCreateReq) (api.PageCreateRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) PageDelete(ctx context.Context, params api.PageDeleteParams) (api.PageDeleteRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) PageEdit(ctx context.Context, req *api.PageCore, params api.PageEditParams) (api.PageEditRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) PageGet(ctx context.Context, params api.PageGetParams) (api.PageGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) PageList(ctx context.Context) (*api.PageListOK, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) NewError(ctx context.Context, err error) *api.ProblemStatusCode {
	if errors.Is(err, category.DuplicateRecordErr) {
		return &api.ProblemStatusCode{
			StatusCode: 409,
			Response: api.Problem{
				Title:    api.NewOptString("category.exists.title"),
				Status:   api.NewOptInt(409),
				Detail:   api.NewOptString("category.exists.detail"),
				Instance: api.NewOptString("/category"),
				Type:     api.NewOptString("about:blank"),
			},
		}
	}

	errorId := uuid.New().String()
	slog.Error("failed to perform task", "id", errorId, "err", err)
	return &api.ProblemStatusCode{
		StatusCode: 500,
		Response: api.Problem{
			Title:           api.NewOptString("internal.server.error.title"),
			Status:          api.NewOptInt(500),
			Detail:          api.NewOptString("internal.server.error.detail"),
			Type:            api.NewOptString("about:blank"),
			ErrorInstanceId: api.NewOptString(errorId),
		},
	}
}
