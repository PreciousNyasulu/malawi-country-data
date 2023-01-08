package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	conn "github.com/preciousnyasulu/malawi-country-data/databaseconnection"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var db = conn.Connect()
var district []structs.District

func GetDistricts(client *gin.Context) {
	rows, err := db.Query("SELECT id, name, code, region FROM districts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	district = nil

	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&getdistrict.Id, &getdistrict.Name, &getdistrict.Code, &getdistrict.Region)
		if err != nil {
			panic(err.Error())
		}
		district = append(district, getdistrict)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	client.IndentedJSON(http.StatusOK, district)
}

func GetDistrictByRegion(client *gin.Context) {

	region := client.Param("region")

	if strings.ToLower(region) == "northern" {
		
		district = districtsappend("Northern")
		client.IndentedJSON(http.StatusOK, district)


	} else if strings.ToLower(region) == "central" {
		
		district = districtsappend("Central")
		client.IndentedJSON(http.StatusOK, district)

	} else if strings.ToLower(region) == "southern" {

		district = districtsappend("Southern")
		client.IndentedJSON(http.StatusOK, district)

	} else {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Unknown route."})
		return
	}
}

func districtsappend(region string) []structs.District {
	rows, err := db.Query("SELECT id, name, code, region FROM districts WHERE region='"+region+"'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	district = nil
	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&getdistrict.Id, &getdistrict.Name, &getdistrict.Code, &getdistrict.Region)
		if err != nil {
			panic(err.Error())
		}

		district = append(district, getdistrict)

		err = rows.Err()
		if err != nil {
			panic(err.Error())
		}
	}

	return district
}
