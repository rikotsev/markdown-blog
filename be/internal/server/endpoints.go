package server

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/rikotsev/markdown-blog/be/gen/api"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"github.com/rikotsev/markdown-blog/be/internal/common"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
	"log/slog"
	"time"
)

func NewEndpointsHandler(ctx ApplicationContext) (api.Handler, error) {
	return &endpointsHandler{
		requestTimeout:  time.Second * 3,
		appCtx:          ctx,
		categoryHandler: category.NewHttp(category.NewRepository(ctx.Pool()), urlid.NewTransformerWith(urlid.Slug)),
	}, nil
}

type endpointsHandler struct {
	requestTimeout  time.Duration
	appCtx          ApplicationContext
	categoryHandler *category.Http
}

func (handler *endpointsHandler) ArticleCreate(ctx context.Context, req *api.ArticleCreateReq) (api.ArticleCreateRes, error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) ArticleList(ctx context.Context, params api.ArticleListParams) (api.ArticleListRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) CategoryCreate(ctx context.Context, req *api.CategoryCreateReq) (api.CategoryCreateRes, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	defer cancelFunc()

	newCategory, err := handler.categoryHandler.CreateCategory(ctx, req)
	if err != nil {
		if errors.Is(err, category.DuplicateRecordErr) {
			return &api.Problem{
				Title: api.NewOptString("category.exists.title"),
				Code:  api.NewOptInt(409),
			}, nil
		}
		errorId := uuid.New().String()
		slog.Error("failed to create new category", "id", errorId, "err", err)
		return common.Problem("Internal Error", "Failed to create category", errorId, 500), nil
	}

	return newCategory, nil
}

func (handler *endpointsHandler) CategoryDelete(ctx context.Context, params api.CategoryDeleteParams) (api.CategoryDeleteRes, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *endpointsHandler) CategoryList(ctx context.Context) (api.CategoryListRes, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, handler.requestTimeout)
	defer cancelFunc()
	categories, err := handler.categoryHandler.ListCategories(ctx)
	if err != nil {
		errorId := uuid.New().String()
		slog.Error("failed to list categories", "id", errorId, "err", err)
		return common.Problem("Internal Error", "Failed to list categories", errorId, 500), nil
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

func (handler *endpointsHandler) PageList(ctx context.Context) (api.PageListRes, error) {
	//TODO implement me
	panic("implement me")
}
