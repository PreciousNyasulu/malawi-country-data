package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"malawi-country-data/src/routes"
)

func main() {
	godotenv.Load()
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
	router.GET("/Village/District/:search", routes.SearchVillageWithDistrict)

	//Traditional Authority Routes
	router.GET("/TraditionalAuthorities", routes.GetTraditionalAuthorities)
	router.GET("/TraditionalAuthority/:search", routes.SearchTraditionalAuthorities)

	// Residential Areas Routes
	router.GET("/ResidentialAreas", routes.GetResidentialAreas)
	router.GET("/ResidentialArea/:search", routes.SearchResidentialArea)

	//Wards Routes
	router.GET("/Wards", routes.GetWards)
	router.GET("/Wards/Region/:search", routes.SearchWardWithRegion)
	router.GET("/Wards/District/:search", routes.SearchWardWithDistrict)

	address := os.Getenv("server_address")
	err := router.Run(address)
	if err != nil {
		log.Fatalf("failed to start server at %s got error %v", address, err)
	}
}
