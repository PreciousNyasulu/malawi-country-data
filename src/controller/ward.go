package controller

import (
	"encoding/json"
	"fmt"
	"malawi-country-data/src/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var wards []structs.Ward

// @Summary 		Gets all Wards
// @Description 	Gets all Ward records
// @Tags 			Wards
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Ward{}
// @Router 			/Wards [get]
func GetWards(client *gin.Context) {
	query := `SELECT 
				json_build_object( 'district', json_build_object('id', d.id, 'name', d.name, 'code', d.code), 'region', d.region, 'constituency', c.name, 'wards', (SELECT json_agg(wa.name) FROM wards wa WHERE wa.constituency_id = c.id))
				AS district
				FROM districts d
				INNER JOIN constituencies c ON c.district_id = d.id`
	SearchWard(client, query)
}

// @Summary 		Gets all wards in the northern region
// @Description 	Gets all northern region wards
// @Tags 			Wards
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Ward{}
// @Router 			/Wards/Region/Northern [get]
func NorthernRegionWards(client *gin.Context) {
	SearchWardWithRegion("Northern", client)
}

// @Summary 		Gets all wards in the southern region
// @Description 	Gets all southern region wards
// @Tags 			Wards
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Ward{}
// @Router 			/Wards/Region/Southern [get]
func SouthernRegionWards(client *gin.Context) {
	SearchWardWithRegion("Southern", client)
}

// @Summary 		Gets all wards in the Central region
// @Description 	Gets all central region wards
// @Tags 			Wards
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Ward{}
// @Router 			/Wards/Region/Central [get]
func CentralRegionWards(client *gin.Context) {
	SearchWardWithRegion("Central", client)
}

func SearchWardWithRegion(Region string, client *gin.Context) {
	query := fmt.Sprintf(`SELECT json_build_object('district', json_build_object('id', d.id, 'name', d.name, 'code', d.code), 'region', d.region, 'constituency', c.name, 'wards', (SELECT json_agg(wa.name) FROM wards wa WHERE wa.constituency_id = c.id)) as district 
							FROM districts d 
							JOIN constituencies c ON c.district_id = d.id
							WHERE d.region ILIKE '%s'`, Region)
	SearchWard(client, query)
}

// @Summary 		Searches wards by district
// @Description 	Gets all wards by district name
// @Param 			search path string true "District name"
// @Tags 			Wards
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Ward{}
// @Router 			/Wards/District/{search} [get]
func SearchWardWithDistrict(client *gin.Context) {
	search := client.Param("search")
	query = fmt.Sprintf(`select json_build_object(
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
  `, search)
	SearchWard(client, query)
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
