package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectionDb() *sqlx.DB {
	dsn := "postgres://koala:koala1234@localhost:5432/koala?sslmode=disable"

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("Successfully connected to PostgreSQL!")

	return db
}