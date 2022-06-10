--подсчёт количества фильмов, имеющих дубли по названию.
select count(1)
from 
	(select "name"   
	from movies m
	group by "name"
	having count(id) > 1) q;
	
---выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;
select 'actor' as role, a.first_name || ' ' || a.last_name as name 
from movies_actors ma join actors a on ma.actor_id = a.id 
group by ma.actor_id, a.first_name , a.last_name 
having count(movie_id) > 2
union
select 'director', d.first_name || ' ' || d.last_name 
from movies_directors md join directors d on md.director_id = d.id 
group by md.director_id, d.first_name, d.last_name 
having count(movie_id) > 2
order by name;

--выборка фильмов для нескольких режиссёров из списка (подзапрос);
select m.id, m."name" movie, m."year", m.box_office, m.rating, c."name" company 
from movies m join companies c on m.company_id = c.id
			  join movies_directors md on m.id  = md.movie_id 
where md.director_id in (select id from directors d where last_name in ('Кэмерон', 'Земекис', 'Крэйвен', 'Хичкок'))	