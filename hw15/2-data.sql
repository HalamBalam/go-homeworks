truncate companies cascade;
truncate actors cascade;
truncate directors cascade;
truncate movies cascade;
truncate movies_actors;
truncate movies_directors;

insert into companies (id, name) values (1, '20th Century Fox');
insert into companies (id, name) values (2, 'Warner Bros.');
insert into companies (id, name) values (3, 'Paramount Pictures');
insert into companies (id, name) values (4, 'Orion Pictures');
insert into companies (id, name) values (5, 'New Line Cinema');
insert into companies (id, name) values (6, 'Universal Pictures');

insert into actors (id, first_name, last_name, date_of_birth) values (1, 'Сэм', 'Уортингтон', '1976-08-02');
insert into actors (id, first_name, last_name, date_of_birth) values (2, 'Зои', 'Салдана', '1978-06-19');
insert into actors (id, first_name, last_name, date_of_birth) values (3, 'Сигурни', 'Уивер', '1949-10-08');
insert into actors (id, first_name, last_name, date_of_birth) values (4, 'Том', 'Хэнкс', '1956-07-09');
insert into actors (id, first_name, last_name, date_of_birth) values (5, 'Дэвид', 'Морс', '1953-10-11');
insert into actors (id, first_name, last_name, date_of_birth) values (6, 'Бонни', 'Хант', '1961-09-22');
insert into actors (id, first_name, last_name, date_of_birth) values (7, 'Робин', 'Райт', '1966-04-08');
insert into actors (id, first_name, last_name, date_of_birth) values (8, 'Салли', 'Филд', '1946-11-06');
insert into actors (id, first_name, last_name, date_of_birth) values (9, 'Киану', 'Ривз', '1964-09-02');
insert into actors (id, first_name, last_name, date_of_birth) values (10, 'Лоренс', 'Фишберн', '1961-07-30');
insert into actors (id, first_name, last_name, date_of_birth) values (11, 'Кэрри-Энн', 'Мосс', '1967-08-21');
insert into actors (id, first_name, last_name, date_of_birth) values (12, 'Арнольд', 'Шварценеггер', '1947-07-30');
insert into actors (id, first_name, last_name, date_of_birth) values (13, 'Майкл', 'Бин', '1956-07-31');
insert into actors (id, first_name, last_name, date_of_birth) values (14, 'Линда', 'Хэмилтон', '1956-09-26');
insert into actors (id, first_name, last_name, date_of_birth) values (15, 'Хьюго', 'Уивинг', '1960-04-04');
insert into actors (id, first_name, last_name, date_of_birth) values (16, 'Хэзер', 'Лэнгенкэмп', '1964-07-17');
insert into actors (id, first_name, last_name, date_of_birth) values (17, 'Роберт', 'Инглунд', '1947-06-06');
insert into actors (id, first_name, last_name, date_of_birth) values (18, 'Джонни', 'Депп', '1963-06-09');
insert into actors (id, first_name, last_name, date_of_birth) values (19, 'Джеки Эрл', 'Хейли', '1961-07-14');
insert into actors (id, first_name, last_name, date_of_birth) values (20, 'Кайл', 'Галлнер', '1986-10-22');
insert into actors (id, first_name, last_name, date_of_birth) values (21, 'Руни', 'Мара', '1985-04-17');
insert into actors (id, first_name, last_name, date_of_birth) values (22, 'Элайджа', 'Вуд', '1981-01-28');
insert into actors (id, first_name, last_name, date_of_birth) values (23, 'Иэн', 'Маккеллен', '1939-05-25');
insert into actors (id, first_name, last_name, date_of_birth) values (24, 'Шон', 'Эстин', '1971-02-25');
insert into actors (id, first_name, last_name, date_of_birth) values (25, 'Майкл Дж.', 'Фокс', '1961-06-09');
insert into actors (id, first_name, last_name, date_of_birth) values (26, 'Кристофер', 'Ллойд', '1938-10-22');
insert into actors (id, first_name, last_name, date_of_birth) values (27, 'Лиа', 'Томпсон', '1961-05-31');

insert into directors (id, first_name, last_name, date_of_birth) values (1, 'Джеймс', 'Кэмерон', '1954-08-16');
insert into directors (id, first_name, last_name, date_of_birth) values (2, 'Фрэнк', 'Дарабонт', '1959-01-28');
insert into directors (id, first_name, last_name, date_of_birth) values (3, 'Роберт', 'Земекис', '1951-05-14');
insert into directors (id, first_name, last_name, date_of_birth) values (4, 'Лана', 'Вачовски', '1965-06-21');
insert into directors (id, first_name, last_name, date_of_birth) values (5, 'Лилли', 'Вачовски', '1967-12-29');
insert into directors (id, first_name, last_name, date_of_birth) values (6, 'Уэс', 'Крэйвен', '1939-08-02');
insert into directors (id, first_name, last_name, date_of_birth) values (7, 'Сэмюэл', 'Байер', '1962-02-17');
insert into directors (id, first_name, last_name, date_of_birth) values (8, 'Питер', 'Джексон', '1961-10-31');

insert into movies (id, name, "year", box_office, rating, company_id)
			values (1, 'Аватар', 2009, 2847379794, 'PG-13', 1);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (2, 'Зеленая миля', 1999, 286801374, 'PG-18', 2);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (3, 'Форрест Гамп', 1994, 677387716, 'PG-13', 3);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (4, 'Матрица', 1999, 463517383, 'PG-18', 2);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (5, 'Терминатор', 1984, 78371200, 'PG-18', 4);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (6, 'Матрица: Перезагрузка', 2003, 742128461, 'PG-18', 2);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (7, 'Кошмар на улице Вязов', 1984, 25563260, 'PG-18', 5);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (8, 'Кошмар на улице Вязов', 2010, 115664037, 'PG-18', 5);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (9, 'Властелин колец: Братство Кольца', 2001, 880839846, 'PG-13', 5);
insert into movies (id, name, "year", box_office, rating, company_id)
			values (10, 'Назад в будущее', 1985, 381109762, 'PG-10', 6);
			
insert into movies_actors (movie_id, actor_id) values (1, 1);
insert into movies_actors (movie_id, actor_id) values (1, 2);
insert into movies_actors (movie_id, actor_id) values (1, 3);
insert into movies_actors (movie_id, actor_id) values (2, 4);
insert into movies_actors (movie_id, actor_id) values (2, 5);
insert into movies_actors (movie_id, actor_id) values (2, 6);
insert into movies_actors (movie_id, actor_id) values (3, 4);
insert into movies_actors (movie_id, actor_id) values (3, 7);
insert into movies_actors (movie_id, actor_id) values (3, 8);
insert into movies_actors (movie_id, actor_id) values (4, 9);
insert into movies_actors (movie_id, actor_id) values (4, 10);
insert into movies_actors (movie_id, actor_id) values (4, 11);
insert into movies_actors (movie_id, actor_id) values (5, 12);
insert into movies_actors (movie_id, actor_id) values (5, 13);
insert into movies_actors (movie_id, actor_id) values (5, 14);
insert into movies_actors (movie_id, actor_id) values (6, 9);
insert into movies_actors (movie_id, actor_id) values (6, 10);
insert into movies_actors (movie_id, actor_id) values (6, 11);
insert into movies_actors (movie_id, actor_id) values (6, 15);
insert into movies_actors (movie_id, actor_id) values (7, 16);
insert into movies_actors (movie_id, actor_id) values (7, 17);
insert into movies_actors (movie_id, actor_id) values (7, 18);
insert into movies_actors (movie_id, actor_id) values (8, 19);
insert into movies_actors (movie_id, actor_id) values (8, 20);
insert into movies_actors (movie_id, actor_id) values (8, 21);
insert into movies_actors (movie_id, actor_id) values (9, 22);
insert into movies_actors (movie_id, actor_id) values (9, 23);
insert into movies_actors (movie_id, actor_id) values (9, 24);
insert into movies_actors (movie_id, actor_id) values (10, 25);
insert into movies_actors (movie_id, actor_id) values (10, 26);
insert into movies_actors (movie_id, actor_id) values (10, 27);

insert into movies_directors (movie_id, director_id) values (1, 1);
insert into movies_directors (movie_id, director_id) values (2, 2);
insert into movies_directors (movie_id, director_id) values (3, 3);
insert into movies_directors (movie_id, director_id) values (4, 4);
insert into movies_directors (movie_id, director_id) values (4, 5);
insert into movies_directors (movie_id, director_id) values (5, 1);
insert into movies_directors (movie_id, director_id) values (6, 4);
insert into movies_directors (movie_id, director_id) values (6, 5);
insert into movies_directors (movie_id, director_id) values (7, 6);
insert into movies_directors (movie_id, director_id) values (8, 7);
insert into movies_directors (movie_id, director_id) values (9, 8);
insert into movies_directors (movie_id, director_id) values (10, 3);