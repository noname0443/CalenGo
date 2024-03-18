CREATE DATABASE IF NOT EXISTS calen_go;

CREATE TABLE IF NOT EXISTS calen_go.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS calen_go.notes (
    id SERIAL PRIMARY KEY,

    name TEXT NOT NULL,
    description TEXT,

    start_date DATE,
    end_date DATE
);