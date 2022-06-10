package storage

import (
	"context"
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
	AddMovies(context.Context, []Movie) error
	DeleteMovie(context.Context, Movie) error
	UpdateMovie(context.Context, Movie) error
	GetMovies(context.Context, ...int) ([]Movie, error)
	AddCompanies(context.Context, []Company) error
	ClearDB(ctx context.Context) error
}
