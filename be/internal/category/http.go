package category

import (
	"context"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
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

func (h *Http) ListCategories(ctx context.Context) ([]api.Category, error) {
	entities, err := h.repository.listCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories from repository: %w", err)
	}

	apiCategories := make([]api.Category, 0, len(entities))

	for _, entity := range entities {
		apiCategories = append(apiCategories, api.Category{
			ID: api.OptString{
				Value: entity.Id,
				Set:   true,
			},
			UrlId: api.OptString{
				Value: entity.UrlId,
				Set:   true,
			},
			Name: api.OptString{
				Value: entity.Name,
				Set:   true,
			},
		})
	}

	return apiCategories, nil
}

func (h *Http) CreateCategory(ctx context.Context, req *api.CategoryCreateReq) (*api.Category, error) {
	entity := Entity{
		Name:  req.Name,
		UrlId: h.urlIdTransformer.Process(req.Name),
	}
	if err := h.repository.createCategory(ctx, &entity); err != nil {
		return nil, fmt.Errorf("failed to create category in repository: %w", err)
	}

	return &api.Category{
		ID:    api.NewOptString(entity.Id),
		UrlId: api.NewOptString(entity.UrlId),
		Name:  api.NewOptString(entity.Name),
	}, nil
}

func (h *Http) DeleteCategory(ctx context.Context, req api.CategoryDeleteParams) (api.CategoryDeleteRes, error) {
	rowsDeleted, err := h.repository.deleteCategory(ctx, string(req.UrlId))
	if err != nil {
		return nil, err
	}

	if rowsDeleted == 0 {
		return &api.CategoryDeleteUnauthorized{}, nil
	}

	return &api.CategoryDeleteOK{}, nil
}
