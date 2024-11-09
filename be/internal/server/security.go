package server

import (
	"context"
	"github.com/rikotsev/markdown-blog/be/gen/api"
)

func NewSecurityHandler() (api.SecurityHandler, error) {
	return &securityHandler{}, nil
}

type securityHandler struct {
}

func (s securityHandler) HandleMainAuth(ctx context.Context, operationName string, t api.MainAuth) (context.Context, error) {
	return ctx, nil
}
