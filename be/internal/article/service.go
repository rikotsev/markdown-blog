package article

import (
	"context"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"github.com/rikotsev/markdown-blog/be/internal/urlid"
)

type Service struct {
	repository       *Repository
	urlIdTransformer *urlid.Transformer
	categoryMapper   *category.Mapper
}

func NewService(repository *Repository, transformer *urlid.Transformer, categoryMapper *category.Mapper) *Service {
	return &Service{
		repository:       repository,
		urlIdTransformer: transformer,
		categoryMapper:   categoryMapper,
	}
}
func (s *Service) CreateArticle(ctx context.Context, req *gen.ArticleCreateJSONBody) (string, error) {
	urlId := s.urlIdTransformer.Process(req.Title)

	entity, err := s.repository.create(ctx, Entity{
		UrlId:       urlId,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Category: category.Entity{
			Id: req.Category.Id,
		},
	})
	if err != nil {
		return "about:blank", fmt.Errorf("failed to create article: %w", err)
	}

	return entity.UrlId, nil
}

func (s *Service) GetArticle(ctx context.Context, urlId string) (*gen.ArticleResponseGet, error) {
	entity, err := s.repository.get(ctx, urlId)
	if err != nil {
		return nil, fmt.Errorf("failed to get article from DB: %w", err)
	}

	if entity == nil {
		return nil, nil
	}

	result := s.mapToApiResource(entity)
	includedItem := gen.Included_Item{}
	err = includedItem.FromCategory(s.categoryMapper.ToHttpLayer(entity.Category))
	if err != nil {
		return nil, err
	}

	return &gen.ArticleResponseGet{
		Data: &result,
		Included: &gen.Included{
			includedItem,
		},
	}, nil
}

func (s *Service) ListArticles(ctx context.Context) (*gen.ArticleResponseList, error) {
	entities, err := s.repository.list(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list articles from DB: %w", err)
	}

	articles := make([]gen.Article, 0, len(entities))
	categories := make([]gen.Included_Item, 0)
	addedCategories := map[string]struct{}{}

	for _, entity := range entities {
		articles = append(articles, s.mapToApiResource(&entity))
		if _, ok := addedCategories[entity.Category.Id]; !ok {
			item := gen.Included_Item{}
			err = item.FromCategory(gen.Category{
				Id:    entity.Category.Id,
				UrlId: entity.Category.UrlId,
				Name:  entity.Category.Name,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to map to category: %w", err)
			}
			categories = append(categories, item)
			addedCategories[entity.Category.Id] = struct{}{}
		}
	}

	result := gen.ArticleResponseList{
		Data:     articles,
		Included: &categories,
	}

	return &result, nil
}

func (s *Service) DeleteArticle(ctx context.Context, urlId string) (bool, error) {
	deleted, err := s.repository.delete(ctx, urlId)
	if err != nil {
		return false, fmt.Errorf("failed to delete article in the DB: %w", err)
	}
	if !deleted {
		return false, nil
	}

	return true, nil
}

func (s *Service) UpdateArticle(ctx context.Context, urlId string, req *gen.ArticleCore) (string, error) {
	modification := EntityModification{}

	if req.Title != nil {
		modification.Title = req.Title
	}

	if req.Description != nil {
		modification.Description = req.Description
	}

	if req.Content != nil {
		modification.Content = req.Content
	}

	if req.Category != nil {
		modification.CategoryId = &req.Category.Id
	}

	_, err := s.repository.update(ctx, urlId, modification)
	if err != nil {
		return "", fmt.Errorf("failed to update article in the DB: %w", err)
	}

	return urlId, nil
}

func (s *Service) mapToApiResource(entity *Entity) gen.Article {
	return gen.Article{
		EntityType:  gen.ArticleEntityTypeArticle,
		Id:          entity.Id,
		UrlId:       entity.UrlId,
		Title:       entity.Title,
		Description: entity.Description,
		Content:     entity.Content,
		CreatedAt:   &entity.Created,
		EditedAt:    &entity.Edited,
		Category: gen.CategoryRef{
			EntityType: gen.CategoryTypeCategory,
			Id:         entity.Category.Id,
		},
	}
}
