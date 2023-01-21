package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/preciousnyasulu/malawi-country-data/routes"
)

func main() {
	godotenv.Load(".env")
	//localhost := os.Getenv("localhost")
	router := gin.Default()
	//District routes
	router.GET("/District", routes.GetDistricts)
	router.GET("/District/Region/:region", routes.GetDistrictByRegion)
	router.GET("/District/Search/:search", routes.Search)

	//Constituency routes
	router.GET("/Constituency", routes.GetConstituencies)
	router.GET("/Constituency/Region/:region", routes.GetConstituenciesWithRegion)

	//Village routes
	router.GET("/Village", routes.GetVillages)

	address := os.Getenv("server_address")
	err := router.Run(address)
	if err != nil {
		log.Fatalf("failed to start server at %s got error %v", address, err)
	}
}
