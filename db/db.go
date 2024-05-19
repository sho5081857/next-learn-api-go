package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewDB() *bun.DB {

	if os.Getenv("GO_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PUBLISHED_PORT"), os.Getenv("DB_DATABASE"))

	sqlDB, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln(err)
	}

	db := bun.NewDB(sqlDB, pgdialect.New())

	// SQLファイルを開く
	file, err := os.Open("_tools/first.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// SQLファイルの内容を読み込む
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// SQLファイルの内容をクエリとして実行
	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	return db
}

func CloseDB(db *bun.DB) {
	sqlDB := db.DB
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
