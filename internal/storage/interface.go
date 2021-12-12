package storage

import (
	"context"
)

type StorageIface interface {
	CreateShortenedURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, slug string) (string, error)
}
