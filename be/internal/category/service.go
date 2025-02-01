package category

import (
	"context"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
)

type Service struct {
	repository       *Repository
	urlIdTransformer *urlid.Transformer
}

func NewService(repository *Repository, transformer *urlid.Transformer) *Service {
	return &Service{
		repository:       repository,
		urlIdTransformer: transformer,
	}
}

func (h *Service) ListCategories(ctx context.Context) (*gen.CategoryResponseList, error) {
	entities, err := h.repository.listCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories from repository: %w", err)
	}

	apiCategories := make([]gen.Category, 0, len(entities))

	for _, entity := range entities {
		apiCategories = append(apiCategories, gen.Category{
			Id:    entity.Id,
			UrlId: entity.UrlId,
			Name:  entity.Name,
		})
	}

	return &gen.CategoryResponseList{
		Data: apiCategories,
	}, nil
}

func (h *Service) CreateCategory(ctx context.Context, req *gen.CategoryCreateJSONBody) (*gen.Category, error) {
	entity := Entity{
		Name:  req.Name,
		UrlId: h.urlIdTransformer.Process(req.Name),
	}
	if err := h.repository.createCategory(ctx, &entity); err != nil {
		return nil, fmt.Errorf("failed to create category in repository: %w", err)
	}

	return &gen.Category{
		EntityType: gen.CategoryTypeCategory,
		Id:         entity.Id,
		UrlId:      entity.UrlId,
		Name:       entity.Name,
	}, nil
}

func (h *Service) DeleteCategory(ctx context.Context, urlId string) (bool, error) {
	rowsDeleted, err := h.repository.deleteCategory(ctx, urlId)
	if err != nil {
		return false, err
	}

	if rowsDeleted == 0 {
		return false, nil
	}

	return true, nil
}
