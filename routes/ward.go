package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"malawi-country-data/structs"

	"github.com/gin-gonic/gin"
)

var wards []structs.Ward

func GetWards(client *gin.Context){
	query := `SELECT 
				json_build_object('id', d.id, 'district_name', d.name, 'district_code', d.code, 'region', d.region, 'constituency', c.name, 'wards', (SELECT json_agg(wa.name) FROM wards wa WHERE wa.constituency_id = c.id))
				AS district
				FROM districts d
				INNER JOIN constituencies c ON c.district_id = d.id`
	SearchWard(client,query)
}

func SearchWardWithRegion(client *gin.Context){
	search := client.Param("search")
	query := fmt.Sprintf(`SELECT json_build_object('id', d.id, 'district_name', d.name, 'district_code', d.code, 'region', d.region, 'constituency', c.name, 'wards', (SELECT json_agg(wa.name) FROM wards wa WHERE wa.constituency_id = c.id)) as district 
							FROM districts d 
							JOIN constituencies c ON c.district_id = d.id
							WHERE d.region ILIKE '%s'`, search)
	SearchWard(client,query)
}

func SearchWardWithDistrict(client *gin.Context){
	search := client.Param("search")
	query := fmt.Sprintf(`SELECT 
	json_build_array(
	  'id', d.id, 
	  'district_name', d.name, 
	  'district_code', d.code, 
	  'region', d.region, 
	  'constituency', c.name, 
	  'wards', (
		SELECT json_agg(wa.name) 
		FROM wards wa 
		WHERE wa.constituency_id=c.id
	  )
	) as district 
  FROM 
	districts d, constituencies c 
  WHERE 
	c.district_id=d.id 
	AND d.name ILIKE '%s'
  `,search)
	SearchWard(client,query)
}

func SearchWard(client *gin.Context, query string) {
	
	rows, err := db.Query(query)
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}

	defer rows.Close()
	wards = nil
	var getwards structs.Ward

	for rows.Next() {
		err = rows.Scan(&newdecoder)
		json.Unmarshal([]byte(newdecoder), &getwards)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		wards = append(wards, getwards)
	}

	err = rows.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, wards)
}
