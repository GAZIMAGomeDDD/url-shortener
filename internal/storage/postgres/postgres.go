package postgres

import (
	"context"

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

func (s *Store) InitSchema(ctx context.Context) error {
	if _, err := s.pool.Exec(ctx, DDL); err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateSlug(ctx context.Context, url string) (string, error) {
	sql := `
		WITH ins AS (
			INSERT INTO shortened_urls  (slug, url)
			VALUES ($1, $2)
			ON CONFLICT (url) DO NOTHING
			RETURNING *
		)
		SELECT slug FROM ins
		UNION
		SELECT slug FROM shortened_urls 
		WHERE url = $1;
	`

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)

	slug := utils.GenerateSlug()
	if err = tx.QueryRow(ctx, sql, slug, url).Scan(&slug); err != nil {
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
