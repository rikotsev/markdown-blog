package page

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type (
	Repository struct {
		pool         *pgxpool.Pool
		queryTimeout time.Duration
	}

	Entity struct {
		Id      string
		UrlId   string
		Title   string
		Content string
	}

	EntityModification struct {
		Title   *string
		Content *string
	}
)

const (
	createPageSql = `
	INSERT INTO page (id, url_id, title, content)
	VALUES ($1, $2, $3, $4)
	RETURNING *;
`
	getPageSql = `
	SELECT
		id, url_id, title, content
	FROM
	    page
	WHERE
	    url_id = $1;
`
	listPagesSql = `
	SELECT
		'' AS id,
		url_id,
		title,
		'' AS content
	FROM
	    page
`
	updatePageSql = `
	UPDATE
		page
	SET
	    title = COALESCE($1, title),
	    content = COALESCE($2, content)
	WHERE 
	    url_id = $3;
`
	deletePageSql = `
	DELETE FROM
	    page
	WHERE
	    url_id = $1;
`
)

func NewRepository(pool *pgxpool.Pool, queryTimeout time.Duration) *Repository {
	return &Repository{
		pool:         pool,
		queryTimeout: queryTimeout,
	}
}

func (r *Repository) create(ctx context.Context, item Entity) (*Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, fmt.Errorf("failed to get conn with: %w", err)
	}

	id := uuid.New()
	rows, err := conn.Query(ctx, createPageSql, id.String(), item.UrlId,
		item.Title, item.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to insert page: %w", err)
	}

	if !rows.Next() {
		return nil, fmt.Errorf("failed to retrieve inserted page: %w", err)
	}

	var page Entity
	err = r.scan(rows, &page)
	if err != nil {
		return nil, fmt.Errorf("failed to map sql columns to fields: %w", err)
	}

	return &page, nil
}

func (r *Repository) get(ctx context.Context, urlId string) (*Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, fmt.Errorf("failed to get conn with: %w", err)
	}

	rows, err := conn.Query(ctx, getPageSql, urlId)
	if err != nil {
		return nil, fmt.Errorf("failed to get page with urlId: %s with error: %w", urlId, err)
	}

	if !rows.Next() {
		return nil, nil
	}

	var page Entity
	err = r.scan(rows, &page)
	if err != nil {
		return nil, fmt.Errorf("failed to map sql columns to fields: %w", err)
	}

	return &page, nil
}

func (r *Repository) list(ctx context.Context) ([]Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, fmt.Errorf("failed to get conn with: %w", err)
	}

	rows, err := conn.Query(ctx, listPagesSql)
	if err != nil {
		return nil, fmt.Errorf("failed to list pages with error: %w", err)
	}

	var result []Entity
	for rows.Next() {
		entity := Entity{}
		err := r.scan(rows, &entity)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}

	return result, nil
}

func (r *Repository) update(ctx context.Context, urlId string, modifications EntityModification) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return false, fmt.Errorf("failed to get conn with: %w", err)
	}

	tag, err := conn.Exec(ctx, updatePageSql, modifications.Title, modifications.Content, urlId)
	if err != nil {
		return false, fmt.Errorf("failed to perform update: %w", err)
	}

	return tag.RowsAffected() > 0, nil
}

func (r *Repository) delete(ctx context.Context, urlId string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return false, fmt.Errorf("failed to get conn with: %w", err)
	}

	tag, err := conn.Exec(ctx, deletePageSql, urlId)
	if err != nil {
		return false, fmt.Errorf("failed to perform update: %w", err)
	}

	return tag.RowsAffected() > 0, nil
}

func (r *Repository) scan(rows pgx.Rows, target *Entity) error {
	return rows.Scan(&target.Id, &target.UrlId, &target.Title, &target.Content)
}
