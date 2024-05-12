package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewDB() *bun.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PUBLISHED_PORT"), os.Getenv("DB_DATABASE"))

	sqlDB, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln(err)
	}

	db := bun.NewDB(sqlDB, pgdialect.New())

	fmt.Println("Connected")
	return db
}

func CloseDB(db *bun.DB) {
	sqlDB := db.DB
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
