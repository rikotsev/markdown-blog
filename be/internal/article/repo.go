package article

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rikotsev/markdown-blog/be/internal/category"
	"time"
)

const (
	createArticle = `
WITH inserted_article AS (
	INSERT INTO
		article (id, url_id, title, description, content, created, edited, category_id)
	SELECT
    	$1, $2, $3, $4, $5, $6, $6, (SELECT id FROM category WHERE url_id = $7)
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
		item.Title, item.Description, item.Content, currentTimestamp, item.Category.UrlId)
	if err != nil {
		return nil, fmt.Errorf("failed to insert article: %w", err)
	}

	var article Entity
	for rows.Next() {
		err := rows.Scan(&article.Id, &article.UrlId, &article.Title,
			&article.Description, &article.Content, &article.Created, &article.Edited,
			&article.Category.Id, &article.Category.UrlId, &article.Category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan from rows: %w", err)
		}
		break
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
	var article Entity
	for rows.Next() {
		err := rows.Scan(&article.Id, &article.UrlId, &article.Title,
			&article.Description, &article.Content, &article.Created, &article.Edited,
			&article.Category.Id, &article.Category.UrlId, &article.Category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan from rows: %w", err)
		}
		break
	}
	return &article, nil
}
