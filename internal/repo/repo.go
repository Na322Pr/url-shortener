package repo

import (
	"url-shortener/internal/repo/pgdb"
	"url-shortener/pkg/postgres"
)

type Url interface {
	CreateURL(urlToSave string, alias string) (int, error)
	GetURL(alias string) (string, error)
	DeleteURLbyID(url_id int) error
}

type Repository struct {
	Url
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{
		Url: pgdb.NewUrlRepo(pg),
	}
}
