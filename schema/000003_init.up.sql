CREATE TABLE parks
(
    id serial not null unique,
	title varchar(255) not null,
	description varchar(255),
	coords int[]
);

CREATE TABLE users_parks
(
    id serial not null unique,
	user_id int references users(id) on delete cascade not null,
	park_id int references parks(id)
);