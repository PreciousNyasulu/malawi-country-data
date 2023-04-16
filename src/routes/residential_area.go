package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"malawi-country-data/src/structs"

	"github.com/gin-gonic/gin"
)

var residential_areas []structs.ResidentialArea

func SearchResidentialArea(client *gin.Context){
	search := client.Param("search")
	query = fmt.Sprintf(`SELECT  json_build_object('id', d.id, 'name', d.name, 'code', d.code, 'region', d.region, 'residential_areas', (SELECT json_agg(ra.name) FROM residential_areas ra WHERE ra.district_id=d.id)) as district 
						FROM districts d, residential_areas r 
						WHERE r.district_id=d.id AND r.name ILIKE '%s'`,search)
	getResidentialArea(client,query)
}

func GetResidentialAreas(client *gin.Context){
	query := `SELECT json_build_object('id', d.id, 'name', d.name, 'code', d.code, 'region', d.region, 'residential_areas', (SELECT json_agg(ra.name) FROM residential_areas ra WHERE ra.district_id=d.id)) as district 
				FROM districts d, residential_areas r 
				WHERE r.district_id=d.id`
	getResidentialArea(client,query)
}

func getResidentialArea(client *gin.Context,query string) {	
	rows, err := db.Query(query)
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer rows.Close()

	residential_areas = nil
	var getResidential_area structs.ResidentialArea

	for rows.Next() {
		err = rows.Scan(&newdecoder)
		json.Unmarshal([]byte(newdecoder), &getResidential_area)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		residential_areas = append(residential_areas, getResidential_area)
	}

	err = rows.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, residential_areas)
}