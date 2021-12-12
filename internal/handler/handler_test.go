package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/handler"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/storage/mockstore"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/utils"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/suite"
)

type unitTestSuite struct {
	suite.Suite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(unitTestSuite))
}

func (s *unitTestSuite) Test_createShortenedURL() {
	store := mockstore.NewStore()
	slug := utils.GenerateSlug()
	store.On("CreateShortenedURL", "https://yandex.ru").Return(slug, nil)
	h := handler.NewHandler(store)

	testSrv := httptest.NewServer(h.Init())
	c := testSrv.Client()

	body, _ := json.Marshal(map[string]string{"url": "https://yandex.ru"})
	req, _ := http.NewRequest("POST", testSrv.URL, bytes.NewReader(body))
	resp, _ := c.Do(req)

	s.Equal(http.StatusCreated, resp.StatusCode)
}

func (s *unitTestSuite) Test_createShortenedURLInternalServerError() {
	store := mockstore.NewStore()
	store.On("CreateShortenedURL", "https://yandex.ru").Return("", fmt.Errorf("some error"))
	h := handler.NewHandler(store)

	testSrv := httptest.NewServer(h.Init())
	c := testSrv.Client()

	body, _ := json.Marshal(map[string]string{"url": "https://yandex.ru"})
	req, _ := http.NewRequest("POST", testSrv.URL, bytes.NewReader(body))
	resp, _ := c.Do(req)

	s.Equal(http.StatusInternalServerError, resp.StatusCode)
}

func (s *unitTestSuite) Test_redirect() {
	store := mockstore.NewStore()
	slug := utils.GenerateSlug()
	store.On("GetURL", slug).Return("https://ozon.ru", nil)
	h := handler.NewHandler(store)

	testSrv := httptest.NewServer(h.Init())
	c := testSrv.Client()

	req, _ := http.NewRequest("GET", testSrv.URL+"/"+slug, nil)
	resp, _ := c.Do(req)

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *unitTestSuite) Test_redirectNotFound() {
	store := mockstore.NewStore()
	slug := utils.GenerateSlug()
	store.On("GetURL", slug).Return("", pgx.ErrNoRows)
	h := handler.NewHandler(store)

	testSrv := httptest.NewServer(h.Init())
	c := testSrv.Client()

	req, _ := http.NewRequest("GET", testSrv.URL+"/"+slug, nil)
	resp, _ := c.Do(req)

	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *unitTestSuite) Test_redirectInternalServerError() {
	store := mockstore.NewStore()
	slug := utils.GenerateSlug()
	store.On("GetURL", slug).Return("", fmt.Errorf("some error"))
	h := handler.NewHandler(store)

	testSrv := httptest.NewServer(h.Init())
	c := testSrv.Client()

	req, _ := http.NewRequest("GET", testSrv.URL+"/"+slug, nil)
	resp, _ := c.Do(req)

	s.Equal(http.StatusInternalServerError, resp.StatusCode)
}
