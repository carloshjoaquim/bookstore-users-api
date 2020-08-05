use users_db;

create table if not exists users
(
	id bigint auto_increment
		primary key,
	first_name varchar(200) null,
	last_name varchar(200) null,
	email varchar(200) not null,
	date_created varchar(45) null,
	status varchar(45) not null,
	password varchar(35) not null,
	constraint users_email_uindex
		unique (email)
);
