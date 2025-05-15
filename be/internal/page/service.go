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
		Pos:     req.Position,
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
			Position:   foundEntity.Pos,
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
			UrlId:    page.UrlId,
			Title:    page.Title,
			Position: page.Pos,
		})
	}

	return &result, nil
}

func (s *Service) updatePage(ctx context.Context, urlId string, data gen.PageCore) (bool, error) {
	success, err := s.repository.update(ctx, urlId, EntityModification{
		Title:   data.Title,
		Content: data.Content,
		Pos:     data.Position,
	})
	if err != nil {
		return false, fmt.Errorf("failed to perform update: %w", err)
	}

	return success, nil
}

func (s *Service) deletePage(ctx context.Context, urlId string) (bool, error) {
	deleted, err := s.repository.delete(ctx, urlId)
	if err != nil {
		return false, fmt.Errorf("failed to delete: %w", err)
	}

	return deleted, nil
}
