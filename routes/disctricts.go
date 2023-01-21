package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

// Gets all the country districts
func GetDistricts(client *gin.Context) {
	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer rows.Close()

	district = nil
	var getdistrict structs.District

	for rows.Next() {
		err = rows.Scan(&newdecoder)
		json.Unmarshal([]byte(newdecoder), &getdistrict)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		district = append(district, getdistrict)
	}

	err = rows.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, district)
}

// Gets districts by country region
func GetDistrictByRegion(client *gin.Context) {

	region := client.Param("region")

	var district []structs.District
	var err error

	if strings.ToLower(region) == "northern" {
		district, err = districtsappend("Northern")

	} else if strings.ToLower(region) == "central" {

		district, err = districtsappend("Central")

	} else if strings.ToLower(region) == "southern" {

		district, err = districtsappend("Southern")

	} else {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Unknown route."})
		return
	}

	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}

	client.IndentedJSON(http.StatusOK, district)
}

func districtsappend(region string) ([]structs.District, error) {
	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d where d.region='" + region + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	district = nil
	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&newdecoder)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(newdecoder), &getdistrict)

		district = append(district, getdistrict)

		err = rows.Err()
		if err != nil {
			return nil, err
		}
	}

	return district, nil
}

// Searches for country districts based on the search parameter
func Search(client *gin.Context) {
	search := client.Param("search")

	rows, err := db.Query("select json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(t.name) FROM traditional_authorities t WHERE t.district_id=d.id)) as district from districts d where d.name LIKE'" + search + "' or d.code LIKE '" + search + "' ")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}

	defer rows.Close()
	district = nil
	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&newdecoder)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		json.Unmarshal([]byte(newdecoder), &getdistrict)
		district = append(district, getdistrict)
		err = rows.Err()
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
	}

	if len(district) == 0 {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "District not found or does not exist"})
	} else {
		client.IndentedJSON(http.StatusOK, district)
	}

}
