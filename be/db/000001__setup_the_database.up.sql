CREATE TABLE IF NOT EXISTS category
(
    id    UUID NOT NULL,
    url_id TEXT NOT NULL,
    name  VARCHAR(100) NOT NULL,
    CONSTRAINT pk_category PRIMARY KEY (id),
    CONSTRAINT uq_category_url_id UNIQUE (url_id)
);

CREATE TABLE IF NOT EXISTS page
(
    id      UUID NOT NULL,
    url_id   TEXT NOT NULL,
    title   VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    pos INTEGER NOT NULL,
    CONSTRAINT pk_page PRIMARY KEY (id),
    CONSTRAINT uq_page_url_id UNIQUE (url_id)
);

CREATE TABLE IF NOT EXISTS article
(
    id UUID NOT NULL,
    url_id TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL,
    edited TIMESTAMPTZ,
    title TEXT NOT NULL,
    description VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,
    category_id UUID,
    CONSTRAINT fk_article_category FOREIGN KEY (category_id) REFERENCES category(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT pk_article PRIMARY KEY (id),
    CONSTRAINT uq_article_url_id UNIQUE (url_id)
)