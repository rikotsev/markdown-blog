package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rikotsev/markdown-blog/be/internal/config"
)

type ApplicationContext interface {
	Ctx() context.Context
	Cfg() *config.Config
	Pool() *pgxpool.Pool
}

type runtimeApplicationContext struct {
	ctx  context.Context
	cfg  *config.Config
	pool *pgxpool.Pool
}

func (r *runtimeApplicationContext) Ctx() context.Context {
	return r.ctx
}

func (r *runtimeApplicationContext) Cfg() *config.Config {
	return r.cfg
}

func (r *runtimeApplicationContext) Pool() *pgxpool.Pool {
	return r.pool
}

func BuildContext(cfg *config.Config) (ApplicationContext, error) {
	serverCtx := context.Background()
	poolCfg, err := pgxpool.ParseConfig(cfg.Database.Url)
	if err != nil {
		return nil, fmt.Errorf("could not create a pgx pool config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(serverCtx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("could not create a pgx Pool: %w", err)
	}

	return &runtimeApplicationContext{
		ctx:  serverCtx,
		cfg:  cfg,
		pool: pool,
	}, nil
}
