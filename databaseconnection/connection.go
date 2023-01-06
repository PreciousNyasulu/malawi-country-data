package databaseconnection

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() sql.DB {
	godotenv.Load(".env")
	//localhost := os.Getenv("localhost")
	password := os.Getenv("password")
	user := os.Getenv("user")
	database := os.Getenv("database")

	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Printf("Connected!")
	return *db
}
