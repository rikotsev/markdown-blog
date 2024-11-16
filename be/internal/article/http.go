package article

import (
	"context"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
)

type Http struct {
	repository       *Repository
	urlIdTransformer *urlid.Transformer
}

func NewHttp(repository *Repository, transformer *urlid.Transformer) *Http {
	return &Http{
		repository:       repository,
		urlIdTransformer: transformer,
	}
}
func (h *Http) CreateArticle(ctx context.Context, req *api.ArticleCreateReq) (api.ArticleCreateRes, error) {
	urlId := h.urlIdTransformer.Process(req.Title)
	categoryUrlId := h.urlIdTransformer.Process(req.Category.Name.Value)

	entity, err := h.repository.create(ctx, Entity{
		UrlId:       urlId,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Category: category.Entity{
			UrlId: categoryUrlId,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create article: %w", err)
	}

	return &api.ArticleCreateCreated{Location: api.NewOptString(entity.UrlId)}, nil
}

func (h *Http) GetArticle(ctx context.Context, params api.ArticleGetParams) (*api.Article, error) {
	entity, err := h.repository.get(ctx, string(params.UrlId))
	if err != nil {
		return nil, fmt.Errorf("failed to get article from DB: %w", err)
	}

	return &api.Article{
		ID:          api.NewOptString(entity.Id),
		UrlId:       api.NewOptString(entity.UrlId),
		Title:       api.NewOptString(entity.Title),
		Description: api.NewOptString(entity.Description),
		Content:     api.NewOptString(entity.Content),
		CreatedAt:   api.NewOptDateTime(entity.Created),
		EditedAt:    api.NewOptDateTime(entity.Edited),
		Category: api.NewOptCategory(api.Category{
			ID:    api.NewOptString(entity.Category.Id),
			Name:  api.NewOptString(entity.Category.Name),
			UrlId: api.NewOptString(entity.Category.UrlId),
		}),
	}, nil
}
