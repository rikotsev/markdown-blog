package article

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"time"
)

// TODO refactor this and make a SQL builder
const (
	createArticle = `
WITH inserted_article AS (
	INSERT INTO
		article (id, url_id, title, description, content, created, edited, category_id)
	SELECT
    	$1, $2, $3, $4, $5, $6, $6, $7
	RETURNING *
)
SELECT 
	a.id,
	a.url_id,
	a.title,
	a.description,
	a.content,
	a.created,
	a.edited,
	c.id AS category_id,
	c.url_id AS category_url_id,
	c.name AS category_name
FROM
	inserted_article AS a
	LEFT JOIN category AS c ON a.category_id = c.id;
`

	getArticle = `
SELECT
	a.id,
	a.url_id,
	a.title,
	a.description,
	a.content,
	a.created,
	a.edited,
	c.id AS category_id,
	c.url_id AS category_url_id,
	c.name AS category_name
FROM
    article AS a
	LEFT JOIN category AS c ON a.category_id = c.id
WHERE
    a.url_id = $1
`

	listArticles = `
SELECT
	a.id,
	a.url_id,
	a.title,
	a.description,
	a.content,
	a.created,
	a.edited,
	c.id AS category_id,
	c.url_id AS category_url_id,
	c.name AS category_name
FROM
    article AS a
	LEFT JOIN category AS c ON a.category_id = c.id
WHERE
    c.url_id = COALESCE($1, c.url_id)
ORDER BY
    a.created;
`

	updateArticle = `
UPDATE
	article
SET
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    content = COALESCE($3, content),
    category_id = COALESCE($4, category_id)
WHERE
    url_id = $5;
`
)

type Repository struct {
	pool         *pgxpool.Pool
	queryTimeout time.Duration
}

type Entity struct {
	Id          string
	UrlId       string
	Title       string
	Description string
	Content     string
	Created     time.Time
	Edited      time.Time
	Category    category.Entity
}

type EntityModification struct {
	Title         *string
	Description   *string
	Content       *string
	CategoryUrlId *string
	CategoryId    *string
}

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
	currentTimestamp := time.Now()
	rows, err := conn.Query(ctx, createArticle, id.String(), item.UrlId,
		item.Title, item.Description, item.Content, currentTimestamp, item.Category.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert article: %w", err)
	}

	if !rows.Next() {
		return nil, nil
	}

	var article Entity
	err = r.scan(rows, &article)
	if err != nil {
		return nil, fmt.Errorf("failed to scan from rows: %w", err)
	}

	return &article, nil
}

func (r *Repository) get(ctx context.Context, urlId string) (*Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, fmt.Errorf("failed to get conn with: %w", err)
	}

	rows, err := conn.Query(ctx, getArticle, urlId)
	if err != nil {
		return nil, fmt.Errorf("failed to get article with url_id: %s with error: %w", urlId, err)
	}

	if !rows.Next() {
		return nil, nil
	}

	var article Entity
	err = r.scan(rows, &article)
	if err != nil {
		return nil, fmt.Errorf("failed to scan from rows: %w", err)
	}

	return &article, nil
}

func (r *Repository) list(ctx context.Context, categoryValue *string) ([]Entity, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, fmt.Errorf("failed to get conn with: %w", err)
	}

	rows, err := conn.Query(ctx, listArticles, categoryValue)
	if err != nil {
		return nil, fmt.Errorf("failed to scan rows: %w", err)
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

func (r *Repository) scan(rows pgx.Rows, article *Entity) error {
	err := rows.Scan(&article.Id, &article.UrlId, &article.Title,
		&article.Description, &article.Content, &article.Created, &article.Edited,
		&article.Category.Id, &article.Category.UrlId, &article.Category.Name)
	return err
}

func (r *Repository) delete(ctx context.Context, urlId string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return false, fmt.Errorf("failed to get conn with: %w", err)
	}

	tag, err := conn.Exec(ctx, "DELETE FROM article WHERE url_id = $1", urlId)
	if err != nil {
		return false, fmt.Errorf("failed to delete article with: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}

func (r *Repository) update(ctx context.Context, urlId string, modification EntityModification) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.queryTimeout)
	defer cancelFunc()

	conn, err := r.pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return false, fmt.Errorf("failed to get conn with: %w", err)
	}

	tag, err := conn.Exec(ctx, updateArticle,
		modification.Title,
		modification.Description,
		modification.Content,
		modification.CategoryId,
		urlId,
	)
	if err != nil {
		return false, fmt.Errorf("failed to delete article with: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}
