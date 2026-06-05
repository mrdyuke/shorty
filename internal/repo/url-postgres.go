package repo

import "github.com/jackc/pgx/v5/pgxpool"

type URLPostgresRepo struct {
	DB *pgxpool.Pool
}

func NewURLPostgresRepo(db *pgxpool.Pool) *URLPostgresRepo {
	return &URLPostgresRepo{DB: db}
}
