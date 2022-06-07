package pgstorage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"homeworks/hw16/pkg/storage"
)

func AddMovies(ctx context.Context, db *pgxpool.Pool, movies []storage.Movie) error {
	// начало транзакции
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	for _, movie := range movies {
		batch.Queue(`INSERT INTO movies(id, name, year, box_office, rating, company_id) VALUES ($1, $2, $3, $4, $5, $6)`,
			movie.ID, movie.Name, movie.Year, movie.BoxOffice, movie.Rating, movie.CompanyID)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}

func DeleteMovie(ctx context.Context, db *pgxpool.Pool, m storage.Movie) error {
	_, err := db.Exec(ctx, `DELETE FROM movies WHERE id = $1`, m.ID)
	return err
}

func UpdateMovie(ctx context.Context, db *pgxpool.Pool, m storage.Movie) error {
	_, err := db.Exec(ctx, `UPDATE movies
								SET name = $1, year = $2, box_office = $3, rating = $4, company_id = $5
								WHERE id = $6`, m.Name, m.Year, m.BoxOffice, m.Rating, m.CompanyID, m.ID)
	return err
}

func GetMovies(ctx context.Context, db *pgxpool.Pool, companyIds ...int) ([]storage.Movie, error) {
	var rows pgx.Rows
	var err error
	sql := `
		SELECT id, name, year, box_office, rating, company_id 
		FROM movies`
	if len(companyIds) > 0 {
		sql += ` WHERE company_id = $1`
		rows, err = db.Query(ctx, sql, companyIds[0])
	} else {
		rows, err = db.Query(ctx, sql)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []storage.Movie
	for rows.Next() {
		var m storage.Movie
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Year,
			&m.BoxOffice,
			&m.Rating,
			&m.CompanyID,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func AddCompanies(ctx context.Context, db *pgxpool.Pool, comps []storage.Company) error {
	// начало транзакции
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	for _, c := range comps {
		batch.Queue(`INSERT INTO companies(id, name) VALUES ($1, $2)`, c.ID, c.Name)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}

func ClearDB(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, `DELETE FROM movies;
								DELETE FROM companies`)
	return err
}
