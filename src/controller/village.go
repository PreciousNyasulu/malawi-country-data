package controller

import (
	"fmt"
	"net/http"

	"malawi-country-data/src/structs"

	"github.com/gin-gonic/gin"
)

var village []structs.Village

// Gets all the country villages
// @Summary 		Gets all the country villages
// @Description 	Gets all the villages
// @Tags 			villages
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Village{}
// @Router 			/Villages [get]
func GetVillages(client *gin.Context) {
	query = `SELECT json_build_object('id',v.id, 'name',v.village_name, 'district',json_build_object('id', d.id, 'name', d.name, 'code', d.code ),'region', d.region)
				FROM villages v
				JOIN districts d ON d.id = v.district_id`
	_search(client, query)
}

// Gets all the villages in a district
// @Summary 		Gets all the country villages in a district
// @Description 	Gets all the villages in district
// @Param 			district path string true "District"
// @Tags 			villages
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Village{}
// @Router 			/Villages/District/{search} [get]
func SearchVillageWithDistrict(client *gin.Context) {
	search := client.Param("search")
	query = fmt.Sprintf(`SELECT json_build_object('id',v.id, 'name',v.village_name, 'district',json_build_object('id', d.id, 'name', d.name, 'code', d.code ),'region', d.region)
	FROM villages v
	JOIN districts d ON d.id = v.district_id 
	WHERE d.name ilike '%%%s%%'`, search)
	_search(client, query)
}

// Searches for the villages
// @Summary 		Searches village name
// @Description 	Gets all the villages in district
// @Param 			village path string true "village"
// @Tags 			villages
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Village{}
// @Router 			/Villages/Search/{search} [get]
func VillageSearch(client *gin.Context){
	search := client.Param("search")
	query = fmt.Sprintf( `SELECT json_build_object('id',v.id, 'name',v.village_name, 'district',json_build_object('id', d.id, 'name', d.name, 'code', d.code ),'region', d.region)
	FROM villages v
	JOIN districts d ON d.id = v.district_id
				WHERE v.name ilike '%%%s%%'`,search)
	_search(client, query)
}

func _search(client *gin.Context, query string) {
	rows, err := db.Query(query)
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer rows.Close()
	var getvillage structs.Village
	for rows.Next() {
		err = rows.Scan(&getvillage.Id , &getvillage.Name, &getvillage.Region, &getvillage.District)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		village = append(village, getvillage)
	}
	err = rows.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, village)
}