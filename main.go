package main

import (
	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/routes"
)

func main(){	 
	router := gin.Default()
	//District routes 
	router.GET("/District",routes.GetDistricts)
	router.GET("/District/:region",routes.GetDistrictByRegion)
	router.GET("/District/Search/:search",routes.Search)

	//Constituency routes 
	router.GET("/Constituency",routes.GetConstituencies) 
	router.GET("/Constituency/Region/:region",routes.GetConstituenciesWithRegion) 
	router.Run("localhost:8000")
}