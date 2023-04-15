package databaseconnection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	password := os.Getenv("password")
	host := os.Getenv("host")
	user := os.Getenv("user")
	dbname := os.Getenv("database")
	sslmode := os.Getenv("sslmode")

	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 5432
	}

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",user, password, host, port, dbname,sslmode))
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Printf("Connected!")
	return db
}
