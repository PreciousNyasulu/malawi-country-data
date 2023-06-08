package main

import (
	"log"
	"malawi-country-data/src/routes"
	_"malawi-country-data/docs"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Malawi Country Data API
// @version 1.0
// @description An API for the malawian country details 
// @host localhost:8080
// @Basepath /api
func main() {
	godotenv.Load()
	//localhost := os.Getenv("localhost")
	router := gin.Default()

	// Documentation
	router.GET("/docs/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	//District routes
	router.GET("/api/District", routes.GetDistricts)
	router.GET("/api/District/Region/:region", routes.GetDistrictByRegion)
	router.GET("/api/District/Search/:search", routes.Search)

	//Constituency routes
	router.GET("/api/Constituency", routes.GetConstituencies)
	router.GET("/api/Constituency/Region/:region", routes.GetConstituenciesWithRegion)

	//Village routes
	router.GET("/api/Village", routes.GetVillages)
	router.GET("/api/Village/District/:search", routes.SearchVillageWithDistrict)

	//Traditional Authority Routes
	router.GET("/api/TraditionalAuthorities", routes.GetTraditionalAuthorities)
	router.GET("/api/TraditionalAuthority/:search", routes.SearchTraditionalAuthorities)

	// Residential Areas Routes
	router.GET("/api/ResidentialAreas", routes.GetResidentialAreas)
	router.GET("/api/ResidentialArea/:search", routes.SearchResidentialArea)

	//Wards Routes
	router.GET("/api/Wards", routes.GetWards)
	router.GET("/api/Wards/Region/:search", routes.SearchWardWithRegion)
	router.GET("/api/Wards/District/:search", routes.SearchWardWithDistrict)

	// Healthcheck endpoint
	router.GET("/HealthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"healthy": true,
		})
	})

	address := os.Getenv("server_address")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server at %s got error %v", address, err)
	}
}
