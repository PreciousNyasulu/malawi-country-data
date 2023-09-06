package router

import (
	"log"
	"malawi-country-data/src/controller"
	_"malawi-country-data/docs"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)
var server_address string
func init(){
	godotenv.Load()
	server_address = os.Getenv("server_address")
}

func Route() {
	//localhost := os.Getenv("localhost")
	router := gin.Default()

	// Documentation
	router.GET("/docs/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	//District controller
	router.GET("/api/Districts", controller.GetDistricts)
	router.GET("/api/Districts/Region/:region", controller.GetDistrictByRegion)
	router.GET("/api/Districts/Search/:search", controller.Search)

	//Constituency controller
	router.GET("/api/Constituencies", controller.GetConstituencies)
	router.GET("/api/Constituencies/Region/Northern", controller.GetNorthernConstituencies)
	router.GET("/api/Constituencies/Region/Southern", controller.GetSouthernConstituencies)
	router.GET("/api/Constituencies/Region/Central", controller.GetCentralConstituencies)
	// router.GET("/api/Constituencies/Region/:region", controller.GetConstituenciesWithRegion)

	//Village controller
	router.GET("/api/Villages", controller.GetVillages)
	router.GET("/api/Villages/Search/:search", controller.GetVillages)
	router.GET("/api/Villages/District/:search", controller.SearchVillageWithDistrict)
 
	//Traditional Authority controller
	router.GET("/api/TraditionalAuthorities", controller.GetTraditionalAuthorities)
	router.GET("/api/TraditionalAuthorities/Search/:search", controller.SearchTraditionalAuthorities)

	// Residential Areas controller
	router.GET("/api/ResidentialAreas", controller.GetResidentialAreas)
	router.GET("/api/ResidentialAreas/Search/:search", controller.SearchResidentialArea)

	//Wards controller
	router.GET("/api/Wards", controller.GetWards)
	router.GET("/api/Wards/Region/Northern", controller.NorthernRegionWards)
	router.GET("/api/Wards/Region/Central", controller.CentralRegionWards)
	router.GET("/api/Wards/Region/Southern", controller.SouthernRegionWards)
	router.GET("/api/Wards/District/:search", controller.SearchWardWithDistrict)

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
