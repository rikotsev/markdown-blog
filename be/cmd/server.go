package main

import (
	"github.com/rikotsev/markdown-blog/be/internal/config"
	"github.com/rikotsev/markdown-blog/be/internal/server"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		slog.Error("failed to init configuration", "err", err)
		os.Exit(1)
	}

	appCtx, err := server.BuildContext(cfg)
	if err != nil {
		slog.Error("failed to spin up application context", "err", err)
	}
	//TODO either move AuthProvider to context or pass db,

	applicationServer, err := server.New(appCtx, server.Okta(cfg))
	if err != nil {
		slog.Error("failed to create markdown blog server", "err", err)
		os.Exit(1)
	}

	if err := applicationServer.Start(); err != nil {
		slog.Error("failed to start markdown blog server", "err", err)
		os.Exit(1)
	}
}
