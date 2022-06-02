-- version 1 bussiness table 
CREATE TABLE t_accounts (
    auto_id SERIAL primary key,
	id uuid unique not null,
	name VARCHAR ( 50 ) NOT NULL,
	dob VARCHAR ( 100 ) NOT NULL,
	address VARCHAR ( 255 ) NOT NULL,
    Description varchar(255),
	created_at TIMESTAMP NOT NULL
);

-- following table
CREATE TABLE t_followings (
	auto_id BIGSERIAL primary key,
	u_id uuid not null,
	following_id uuid not null
);
-- create index for following table
CREATE INDEX u_f_id_idx ON t_followings (u_id, following_id);
