CREATE TABLE users
(
    id serial not null unique,
	name varchar(255)  not null,
	username varchar(255)  not null unique,
	password_hash varchar(255)  not null
);

CREATE TABLE parks
(
    id serial not null unique,
	title varchar(255) not null,
	coords JSON not null,
	description varchar(500),
    address varchar(255)
);

CREATE TABLE users_parks
(
    id serial not null unique,
	user_id int references users(id) on delete cascade not null,
	park_id int references parks(id) on delete cascade not null
);
