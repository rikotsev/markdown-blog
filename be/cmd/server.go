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
	}
	markdownBlog, err := server.New(server.BuildContext(cfg))
	if err != nil {
		slog.Error("failed to create markdown blog server", "err", err)
		os.Exit(1)
	}

	if err := markdownBlog.Start(); err != nil {
		slog.Error("failed to start markdown blog server", "err", err)
		os.Exit(1)
	}
}
