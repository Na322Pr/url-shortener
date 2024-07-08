package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Conn    *pgx.Conn
}

func Connection(url string) (*Postgres, error) {
	pg := &Postgres{}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	var err error
	pg.Conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("error database connecting: %w", err)
	}

	_, err = pg.Conn.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		alias TEXT NOT NULL UNIQUE);
		CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);`)

	if err != nil {
		return nil, fmt.Errorf("error in database: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Conn != nil {
		p.Conn.Close(context.Background())
	}
}
