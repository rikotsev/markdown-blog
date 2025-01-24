package category

import (
	"context"
	"encoding/json"
	"github.com/rikotsev/markdown-blog/be/gen"
	"net/http"
	"time"
)

type Http struct {
	service *Service
	timeout time.Duration
}

func NewHttp(service *Service) *Http {
	return &Http{
		service: service,
		timeout: 5 * time.Second,
	}
}

func (h *Http) ListCategories(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	response, err := h.service.ListCategories(ctx)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}

func (h *Http) CreateCategory(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	var createCategoryRequest gen.CategoryCreateJSONBody
	err := json.NewDecoder(r.Body).Decode(&createCategoryRequest)
	if err != nil {
		return err
	}

	newCategory, err := h.service.CreateCategory(ctx, &createCategoryRequest)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newCategory)
	if err != nil {
		return err
	}

	return nil
}

func (h *Http) DeleteCategory(w http.ResponseWriter, r *http.Request, urlId gen.UrlId, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	deleted, err := h.service.DeleteCategory(ctx, urlId)
	if err != nil {
		return err
	}

	if !deleted {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
