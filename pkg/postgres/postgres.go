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

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Conn != nil {
		p.Conn.Close(context.Background())
	}
}
