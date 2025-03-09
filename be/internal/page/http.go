package page

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

func (h *Http) PageCreate(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	var req gen.PageCreateJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return fmt.Errorf("failed to decode body: %w", err)
	}

	urlId, err := h.service.createPage(ctx, &req)
	if err != nil {
		return fmt.Errorf("failed to create new page: %w", err)
	}

	w.Header().Add("Location", urlId)
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *Http) PageGet(w http.ResponseWriter, r *http.Request, urlId gen.UrlId, ctx context.Context) error {
	page, err := h.service.getPage(ctx, urlId)
	if err != nil {
		return fmt.Errorf("failed to retrieve page: %w", err)
	}

	if page == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	err = json.NewEncoder(w).Encode(page)
	if err != nil {
		return fmt.Errorf("failed to encode page: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *Http) PageList(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	pages, err := h.service.listPages(ctx)
	if err != nil {
		return fmt.Errorf("failed to list pages: %w", err)
	}

	err = json.NewEncoder(w).Encode(&pages)
	if err != nil {
		return fmt.Errorf("failed to encode pages: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
