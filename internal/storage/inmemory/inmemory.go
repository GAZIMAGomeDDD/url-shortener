package inmemory

import (
	"context"
	"sync"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/utils"
	"github.com/jackc/pgx/v4"
)

type shortenedUrl struct {
	slug string
	url  string
}

type Store struct {
	sync.Mutex
	m map[string]shortenedUrl
}

func NewStore() *Store {
	return &Store{
		m: map[string]shortenedUrl{},
	}
}

func (s *Store) CreateSlug(ctx context.Context, url string) (string, error) {
	s.Lock()
	defer s.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	surl, ok := s.m[url]
	if ok {
		return surl.slug, nil
	}

	surl = shortenedUrl{
		slug: utils.GenerateSlug(),
		url:  url,
	}

	s.m[url] = surl

	return surl.slug, nil
}

func (s *Store) GetURL(ctx context.Context, slug string) (string, error) {
	s.Lock()
	defer s.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	for _, surl := range s.m {
		if surl.slug == slug {
			return surl.url, nil
		}
	}

	return "", pgx.ErrNoRows
}
