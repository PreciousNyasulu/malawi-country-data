package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"malawi-country-data/structs"
	"github.com/gin-gonic/gin"
)

var query = ""

// Gets all the country districts
func GetDistricts(client *gin.Context) {
	query = `SELECT 
				json_build_object(
					'id', d.id,
					'name', d.name,
					'code', d.code,
					'region', d.region,
					'traditional_authorities', (
						SELECT json_agg(t.name)
						FROM traditional_authorities t
						WHERE t.district_id = d.id
					)
				) AS district
			FROM districts d`
	rows, err := db.Query(query)
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

	query = fmt.Sprintf(`
		SELECT json_build_object(
			'id', d.id,
			'name', d.name,
			'code', d.code,
			'region', d.region,
			'traditional_authorities', (
				SELECT json_agg(t.name)
				FROM traditional_authorities t
				WHERE t.district_id = d.id
			)
		) AS district
		FROM districts d
		WHERE d.region = '%s'
	`, region)

	rows, err := db.Query(query)
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

	query = fmt.Sprintf(`SELECT 
							json_build_object(
								'id', d.id,
								'name', d.name,
								'code', d.code,
								'region', d.region,
								'traditional_authorities', (
									SELECT json_agg(t.name)
									FROM traditional_authorities t
									WHERE t.district_id = d.id
								)
							) AS district
						FROM districts d
						WHERE d.name ILIKE '%%%s%%' OR d.code ILIKE '%%%s%%'
						`, search,search)
	rows, err := db.Query(query)
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
