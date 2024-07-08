package pgdb

import (
	"context"
	"fmt"
	"url-shortener/pkg/postgres"
)

type UrlRepo struct {
	*postgres.Postgres
}

func NewUrlRepo(pg *postgres.Postgres) *UrlRepo {
	return &UrlRepo{pg}
}

func (r *UrlRepo) CreateURL(urlToSave string, alias string) (int, error) {
	const op = "repo.pgdb.URLRepo.CreateURL"

	sql, args, _ := r.Builder.Insert("urls").
		Columns("url", "alias").
		Values(urlToSave, alias).
		Suffix("RETURNING id").
		ToSql()

	var id int
	err := r.Conn.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s - r.Conn.Exec: %w", op, err)
	}

	return id, nil
}

func (r *UrlRepo) GetURL(alias string) (string, error) {
	const op = "repo.pgdb.URLRepo.GetURL"

	sql, args, _ := r.Builder.Select("url").
		From("urls").
		Where("alias = ?", alias).
		ToSql()

	var urlToReturn string
	err := r.Conn.QueryRow(context.Background(), sql, args...).Scan(&urlToReturn)
	if err != nil {
		return "", fmt.Errorf("%s - r.Conn.Exec: %w", op, err)
	}

	return urlToReturn, nil
}

func (r *UrlRepo) DeleteURLbyID(url_id int) error {
	const op = "repo.pgdb.URLRepo.DeleteURLbyID"

	sql, args, _ := r.Builder.Delete("urls").Where("id = ?", url_id).ToSql()

	_, err := r.Conn.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s - r.Conn.Exec: %w", op, err)
	}

	return nil
}
