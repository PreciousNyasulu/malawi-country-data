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

// Gets all the country districts
func GetDistricts(client *gin.Context) {
	rows, err := db.Query("SELECT id, name, code, region FROM districts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	district = nil
	var getdistrict structs.District
	for rows.Next() {
		
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

// Gets districts by country region
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
	rows, err := db.Query("SELECT id, name, code, region FROM districts WHERE region='" + region + "'")
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

// Searches for country districts based on the search parameter
func Search(client *gin.Context) {
	search := client.Param("search")

	rows, err := db.Query("SELECT id,name,code,region FROM districts WHERE name LIKE '%" + search + "%' OR code='" + search + "'")
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

	if len(district) == 0 {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "District not found or does not exist"})
	}else{
		client.IndentedJSON(http.StatusOK, district)
	}

	
}
