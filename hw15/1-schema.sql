DROP TABLE IF EXISTS movies_actors;
DROP TABLE IF EXISTS movies_directors;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS companies;

CREATE TABLE companies (
    id SERIAL PRIMARY KEY, -- первичный ключ
    name TEXT NOT NULL DEFAULT ''
);

CREATE TABLE actors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    date_of_birth DATE
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    date_of_birth DATE
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY, -- первичный ключ
    name TEXT NOT NULL DEFAULT '',
    year INTEGER NOT null check (year >= 1800),
    box_office BIGINT default 0,
    rating TEXT CHECK (rating IN ('PG-10', 'PG-13', 'PG-18')),
    company_id INTEGER NOT NULL REFERENCES companies(id),
    UNIQUE(year, name)
);

CREATE TABLE movies_actors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    movie_id INTEGER NOT NULL REFERENCES movies(id),
    actor_id INTEGER NOT NULL REFERENCES actors(id),
    UNIQUE(movie_id, actor_id)
);

CREATE TABLE movies_directors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    movie_id INTEGER NOT NULL REFERENCES movies(id),
    director_id INTEGER NOT NULL REFERENCES directors(id),
    UNIQUE(movie_id, director_id)
);