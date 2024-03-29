package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile string = "db.db"

const createUsers string = `
CREATE TABLE IF NOT EXISTS users (
id INTEGER NOT NULL PRIMARY KEY,
first_name TEXT NOT NULL,
last_name TEXT NOT NULL,
email TEXT UNIQUE NOT NULL,
password TEXT NOT NULL
);`

const createSessions string = `
CREATE TABLE IF NOT EXISTS sessions (
id INTEGER NOT NULL PRIMARY KEY,
session_id TEXT NOT NULL,
user_id INTEGER NOT NULL,
expires INTEGER NOT NULL,
FOREIGN KEY(user_id) REFERENCES users(id)
);`

func Setup(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, createUsers)
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, createSessions)
	if err != nil {
		return nil, err
	}

	return db, nil
}
