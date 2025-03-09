package page

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

func NewService(repository *Repository, urlIdTransformer *urlid.Transformer) *Service {
	return &Service{
		repository:       repository,
		urlIdTransformer: urlIdTransformer,
	}
}

func (s *Service) createPage(ctx context.Context, req *gen.PageCreateJSONBody) (string, error) {
	newEntity, err := s.repository.create(ctx, Entity{
		Title:   req.Title,
		Content: req.Content,
		UrlId:   s.urlIdTransformer.Process(req.Title),
	})

	if err != nil {
		return "", fmt.Errorf("failed to create db record: %w", err)
	}

	return newEntity.UrlId, nil
}

func (s *Service) getPage(ctx context.Context, urlId string) (*gen.PageResponseGet, error) {
	foundEntity, err := s.repository.get(ctx, urlId)
	if err != nil {
		return nil, fmt.Errorf("failed to get entity: %w", err)
	}
	//does not exist
	if foundEntity == nil {
		return nil, nil
	}

	return &gen.PageResponseGet{
		Data: &gen.Page{
			EntityType: gen.PageEntityTypePage,
			Id:         foundEntity.Id,
			UrlId:      foundEntity.UrlId,
			Title:      foundEntity.Title,
			Content:    foundEntity.Content,
		},
	}, nil
}

func (s *Service) listPages(ctx context.Context) (*gen.PageResponseList, error) {
	entities, err := s.repository.list(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve entities: %w", err)
	}

	result := gen.PageResponseList{
		Data: make([]gen.PageUrlIdAndTitle, 0, len(entities)),
	}

	for _, page := range entities {
		result.Data = append(result.Data, gen.PageUrlIdAndTitle{
			UrlId: &page.UrlId,
			Title: &page.Title,
		})
	}

	return &result, nil
}
