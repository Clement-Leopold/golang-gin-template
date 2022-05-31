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