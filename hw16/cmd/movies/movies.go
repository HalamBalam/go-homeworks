package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"homeworks/hw16/pkg/pgstorage"
	"homeworks/hw16/pkg/storage"
	"log"
)

func main() {
	ctx := context.Background()
	pwd := "12345678"
	pool, err := pgxpool.Connect(ctx, "postgres://postgres:"+pwd+"@localhost/movies")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	db := pgstorage.New(pool)
	err = db.ClearDB(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	comps := []storage.Company{
		{ID: 0, Name: "<не указана>"},
		{ID: 1, Name: "20th Century Fox"},
		{ID: 2, Name: "Warner Bros."},
		{ID: 3, Name: "Paramount Pictures"},
	}
	err = db.AddCompanies(ctx, comps)
	if err != nil {
		fmt.Println(err)
		return
	}

	movies := []storage.Movie{
		{ID: 1, Name: "Аватар", Year: 2009, BoxOffice: 2847379794, Rating: "PG-13", CompanyID: 1},
		{ID: 2, Name: "Зеленая миля", Year: 1999, BoxOffice: 286801374, Rating: "PG-18", CompanyID: 2},
		{ID: 3, Name: "Форрест Гамп", Year: 1994, BoxOffice: 677387716, Rating: "PG-13", CompanyID: 3},
	}
	err = db.AddMovies(ctx, movies)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := db.GetMovies(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Добавлено 3 фильма:")
	fmt.Printf("%+v\n", data)

	err = db.DeleteMovie(ctx, movies[len(movies)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = db.GetMovies(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Удален 1 фильм:")
	fmt.Printf("%+v\n", data)

	m := movies[0]
	m.Name = "Матрица"
	m.Year = 1999
	err = db.UpdateMovie(ctx, m)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = db.GetMovies(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Изменен 1 фильм:")
	fmt.Printf("%+v\n", data)

	data, err = db.GetMovies(ctx, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Фильмы одной компании:")
	fmt.Printf("%+v\n", data)
}
