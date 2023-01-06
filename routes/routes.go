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
		rows, err := db.Query("SELECT id, name, code, region FROM districts WHERE region='Northern'")
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
	} else if strings.ToLower(region) == "central" {
		rows, err := db.Query("SELECT id, name, code, region FROM districts WHERE region='Central'")
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
	} else if strings.ToLower(region) == "southern" {
		rows, err := db.Query("SELECT id, name, code, region FROM districts WHERE region='Southern'")
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
		client.IndentedJSON(http.StatusOK, district)
	} else {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Unknown route."})
		return
	}

}
