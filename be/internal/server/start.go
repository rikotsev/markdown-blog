package server

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
	"net"
	"net/http"
)

type ApplicationServer struct {
	Ctx      ApplicationContext
	srv      *api.Server
	Listener net.Listener
}

func New(ctx ApplicationContext, contextBuildError error) (*ApplicationServer, error) {
	if contextBuildError != nil {
		return nil, fmt.Errorf("failed to build markdown blog context: %w", contextBuildError)
	}

	listener, err := net.Listen("tcp", ctx.Cfg().Server.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s with %w", ctx.Cfg().Server.Address, err)
	}

	endpointsHandler, err := NewEndpointsHandler(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint handler: %w", err)
	}
	securityHandler, err := NewSecurityHandler()
	if err != nil {
		return nil, fmt.Errorf("failed to create security handler: %w", err)
	}
	srv, err := api.NewServer(endpointsHandler, securityHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	return &ApplicationServer{
		Ctx:      ctx,
		srv:      srv,
		Listener: listener,
	}, nil
}

func (m *ApplicationServer) Start() error {
	return http.Serve(m.Listener, m.srv)
}
