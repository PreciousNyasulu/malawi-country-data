package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	conn "github.com/preciousnyasulu/malawi-country-data/databaseconnection"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var db = conn.Connect()
var district []structs.District
var newdecoder string

// Gets all the country districts
func GetDistricts(client *gin.Context) {
	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	district = nil
	var getdistrict structs.District
	
	for rows.Next() {	
		err = rows.Scan(&newdecoder)
		json.Unmarshal([]byte(newdecoder),&getdistrict)
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
	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d where d.region='"+region+"'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	district = nil
	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&newdecoder)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(newdecoder),&getdistrict)

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

	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d where d.name LIKE'"+search+"' or d.code LIKE '"+search+"' ")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	district = nil
	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&newdecoder)
		if err != nil {
			panic(err.Error())
			
		}
		json.Unmarshal([]byte(newdecoder),&getdistrict)
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
