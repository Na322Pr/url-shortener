package service

import (
	"url-shortener/internal/repo"
)

type UrlService struct {
	urlRepo repo.Url
}

func NewUrlService(urlRepo repo.Url) *UrlService {
	return &UrlService{urlRepo: urlRepo}
}

func (s *UrlService) CreateURL(urlToSave string, alias string) (int, error) {
	return s.urlRepo.CreateURL(urlToSave, alias)
}

func (s *UrlService) GetURL(alias string) (string, error) {
	return s.urlRepo.GetURL(alias)
}

func (s *UrlService) DeleteURLbyID(url_id int) error {
	return s.urlRepo.DeleteURLbyID(url_id)
}
