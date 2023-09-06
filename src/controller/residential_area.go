package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"malawi-country-data/src/structs"

	"github.com/gin-gonic/gin"
)

var residential_areas []structs.ResidentialArea

// @Summary 		Searches for Residential Areas
// @Description 	Searches though Residential Area records
// @Tags 			Residential Area
// @Param 			search path string true "Residential Area"
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.ResidentialArea{}
// @Router 			/ResidentialAreas/Search/{search} [get]
func SearchResidentialArea(client *gin.Context){
	search := client.Param("search")
	query = fmt.Sprintf(`SELECT  json_build_object('district', jsonb_build_object('id', d.id, 'name', d.name, 'code', d.code), 'region', d.region, 'residential_areas', (SELECT json_agg(ra.name) FROM residential_areas ra WHERE ra.district_id=d.id)) as district 
	FROM districts d, residential_areas r 
	WHERE r.district_id=d.id AND r.name ilike '%s'`,search)
	getResidentialArea(client,query)
}

// @Summary 		Gets all Residential Areas
// @Description 	Gets all Residential Area records
// @Tags 			Residential Area
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.ResidentialArea{}
// @Router 			/ResidentialAreas [get]
func GetResidentialAreas(client *gin.Context){
	query := `SELECT json_build_object( 'district', jsonb_build_object('id', d.id, 'name', d.name, 'code', d.code), 'region', d.region, 'residential_areas', (SELECT json_agg(ra.name) FROM residential_areas ra WHERE ra.district_id=d.id)) as district 
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