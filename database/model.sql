-- MODEL
create database wallet;

create extension pgcrypto;

drop table if exists users cascade;
drop table if exists expenditures cascade;
drop table if exists cash cascade;

create table users (
	user_id serial not null primary key,
	full_name character varying(64) not null,
	username character varying(16) not null,
	password character varying(60) not null,
	total_money decimal(10, 2) default 0
);

create table cash ( /*cashes is true name*/
	cash_id serial not null primary key,
	user_id int not null references users (user_id),	
	amount decimal(10, 2) not null,
	summary character varying(256) not null,
	received_at timestamp with time zone default current_timestamp
);

create table expenditures (
	expenditure_id serial not null primary key,
	user_id int not null references users (user_id),
	amount decimal(10, 2) not null,
	summary character varying(256) not null,
	spent_at timestamp with time zone default current_timestamp
);

	