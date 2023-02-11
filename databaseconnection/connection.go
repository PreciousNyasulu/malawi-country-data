package databaseconnection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	password := os.Getenv("password")
	host := os.Getenv("server")
	user := os.Getenv("user")
	database := os.Getenv("database")

	// port, err := strconv.Atoi(os.Getenv("port"))
	// if err != nil {
	// 	port = 3306
	// }

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database))
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
