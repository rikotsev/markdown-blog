package article

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
)

type Http struct {
	service *Service
}

func NewHttp(service *Service) *Http {
	return &Http{
		service: service,
	}
}

func (h *Http) CreateArticle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	var createArticleBody gen.ArticleCreateJSONBody
	err := json.NewDecoder(r.Body).Decode(&createArticleBody)
	if err != nil {
		return fmt.Errorf("failed to decode body: %w", err)
	}

	//TODO JSON validations

	urlId, err := h.service.CreateArticle(ctx, &createArticleBody)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	w.Header().Add("Location", urlId)
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h *Http) GetArticle(w http.ResponseWriter, r *http.Request, urlId gen.UrlId, ctx context.Context) error {
	resp, err := h.service.GetArticle(ctx, urlId)
	if err != nil {
		return fmt.Errorf("failed to get article: %w", err)
	}

	if resp == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Http) ArticleList(w http.ResponseWriter, r *http.Request, params gen.ArticleListParams, ctx context.Context) error {
	resp, err := h.service.ListArticles(ctx, params.Category)
	if err != nil {
		return fmt.Errorf("failed to list articles: %w", err)
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *Http) ArticleDelete(w http.ResponseWriter, r *http.Request, urlId gen.UrlId, ctx context.Context) error {
	resp, err := h.service.DeleteArticle(ctx, urlId)
	if err != nil {
		return fmt.Errorf("failed to delete article: %w", err)
	}

	if !resp {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *Http) ArticleEdit(w http.ResponseWriter, r *http.Request, urlId gen.UrlId, ctx context.Context) error {
	var articleCore gen.ArticleCore
	err := json.NewDecoder(r.Body).Decode(&articleCore)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	location, err := h.service.UpdateArticle(ctx, urlId, &articleCore)
	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}

	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusOK)

	return nil
}
