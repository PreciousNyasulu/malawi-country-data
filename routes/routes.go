package routes

import (
	"database/sql"

	"github.com/joho/godotenv"
	conn "malawi-country-data/databaseconnection"
	"malawi-country-data/structs"
)

// just the init
var db *sql.DB
var district []structs.District
var newdecoder string

func init() {
	godotenv.Load()
	db = conn.Connect()
}
