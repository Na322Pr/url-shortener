package service

import "url-shortener/internal/repo"

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=Url
type Url interface {
	CreateURL(urlToSave string, alias string) (int, error)
	GetURL(alias string) (string, error)
	DeleteURLbyID(id int) error
}

type Service struct {
	Url Url
}

type ServiceDependencies struct {
	Repo *repo.Repository
}

func NewService(deps ServiceDependencies) *Service {
	return &Service{
		Url: NewUrlService(deps.Repo.Url),
	}
}
