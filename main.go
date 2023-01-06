package main

import (
	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/routes"
)

func main(){	 
	router := gin.Default()
	router.GET("/District",routes.GetDistricts)
	router.GET("/District/:region",routes.GetDistrictByRegion)
	router.Run("localhost:8000")
}