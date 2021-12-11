package postgres

import (
	"context"
	"fmt"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DDL = `
		DROP TABLE IF EXISTS shortened_urls;
		
		CREATE TABLE IF NOT EXISTS shortened_urls 
		(
			slug  varchar(10)  PRIMARY KEY,
			url   text 	   NOT NULL UNIQUE
		);
	`
)

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(ctx context.Context, db *pgxpool.Pool) (*Store, error) {
	s := new(Store)
	s.pool = db

	if err := s.initSchema(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) initSchema(ctx context.Context) error {
	if _, err := s.pool.Exec(ctx, DDL); err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateShortenedURL(ctx context.Context, url string) (string, error) {
	sql := `
		WITH row AS (
			INSERT INTO shortened_urls  (slug, url)
			VALUES ($1, $2)
			ON CONFLICT (url) DO NOTHING
			RETURNING *
		)
		SELECT slug FROM row
		UNION
		SELECT slug FROM shortened_urls 
		WHERE url = $2;
	`

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)

	slug := utils.GenerateSlug()
	if err = tx.QueryRow(ctx, sql, slug, url).Scan(&slug); err != nil {
		fmt.Println(err)

		return "", err
	}

	if err = tx.Commit(ctx); err != nil {
		return "", err
	}

	return slug, nil
}

func (s *Store) GetURL(ctx context.Context, slug string) (string, error) {
	var url string

	sql := `
		SELECT url FROM shortened_urls 
		WHERE slug = $1;
	`

	if err := s.pool.QueryRow(ctx, sql, slug).Scan(&url); err != nil {
		return "", err
	}

	return url, nil
}
