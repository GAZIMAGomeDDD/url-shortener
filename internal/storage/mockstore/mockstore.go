package mockstore

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type Store struct {
	mock.Mock
}

func NewStore() *Store {
	store := new(Store)

	return store
}

func (s *Store) CreateShortenedURL(ctx context.Context, url string) (string, error) {
	args := s.Called(url)

	arg0 := args.Get(0)

	if arg0 == "" {
		return "", args.Error(1)
	}

	return arg0.(string), args.Error(1)
}

func (s *Store) GetURL(ctx context.Context, slug string) (string, error) {
	args := s.Called(slug)

	arg0 := args.Get(0)

	if arg0 == "" {
		return "", args.Error(1)
	}

	return arg0.(string), args.Error(1)
}
