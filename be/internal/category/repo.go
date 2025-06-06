package category

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

const listAllCategories = `
	SELECT
		id,
		url_id,
		name
	FROM
	    category;
`

const createCategory = `
	INSERT INTO
		category (id, url_id, name)
	VALUES 
		($1, $2, $3);
`

const deleteCategory = `
	DELETE FROM
		category
	WHERE
	    url_id = $1;
`

var DuplicateRecordErr = errors.New("a category with that url id already exists. please select a different name")

type Repository struct {
	pool         *pgxpool.Pool
	queryTimeout time.Duration
}

type Entity struct {
	Id    string
	UrlId string
	Name  string
}

func NewRepository(pool *pgxpool.Pool, queryTimeout time.Duration) *Repository {
	return &Repository{
		pool:         pool,
		queryTimeout: queryTimeout,
	}
}

func (r *Repository) listCategories(ctx context.Context) ([]*Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire db conection: %w", err)
	}
	defer conn.Release()

	var categories []*Entity
	if err := pgxscan.Select(ctx, r.pool, &categories, listAllCategories); err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}

	return categories, nil
}

func (r *Repository) createCategory(ctx context.Context, entity *Entity) error {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	trx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire db connection: %w", err)
	}
	defer func(trx pgx.Tx, ctx context.Context) {
		_ = trx.Rollback(ctx)
	}(trx, ctx)

	id := uuid.New()

	if res, err := trx.Exec(ctx, createCategory, id, entity.UrlId, entity.Name); err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return DuplicateRecordErr
			}
		}
		return fmt.Errorf("failed to execute insert category: %w", err)
	} else {
		slog.Debug("createCategory", "rows affected", res.RowsAffected())
	}

	entity.Id = id.String()
	err = trx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

func (r *Repository) deleteCategory(ctx context.Context, urlId string) (int, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return 0, err
	}

	tag, err := conn.Exec(ctx, deleteCategory, urlId)
	if err != nil {
		return 0, err
	}

	return int(tag.RowsAffected()), nil
}
