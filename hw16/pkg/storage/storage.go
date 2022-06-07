package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Movie struct {
	ID        int
	Name      string
	Year      int
	BoxOffice int64
	Rating    string
	CompanyID int
}

type Company struct {
	ID   int
	Name string
}

type Interface interface {
	AddMovies(context.Context, *pgxpool.Pool, []Movie) error
	DeleteMovie(context.Context, *pgxpool.Pool, Movie) error
	UpdateMovie(context.Context, *pgxpool.Pool, Movie) error
	GetMovies(context.Context, *pgxpool.Pool, ...int) ([]Movie, error)
}
