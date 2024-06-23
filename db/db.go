package db

import (
	"context"
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
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database:", err)
	}

	// テーブルが存在するかどうかを確認するクエリを実行
	var exists bool
	err = db.NewSelect().
		ColumnExpr("EXISTS (SELECT FROM pg_tables WHERE tablename = ?) AS exists", "users").
		Scan(context.Background(), &exists)
	if err != nil {
		log.Println(err)
	}

	// テーブルが存在しない場合のみ選択されたコードを実行
	if !exists {
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
	}

	fmt.Println("Connected")
	return db
}
