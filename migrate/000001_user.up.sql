CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE users (
    id      serial not null unique,
    name    varchar(255) not null,
    phone   varchar(11) not null unique,
    email   varchar(255) unique,
    gender  gender,
    birthday date,
    password_hash varchar(255),
    verifed boolean default false
);
