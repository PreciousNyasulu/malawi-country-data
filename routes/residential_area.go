package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var residential_areas []structs.ResidentialArea

func SearchResidentialArea(client *gin.Context) {
	search := client.Param("search")
	rows, err := db.Query("select DISTINCT json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'residential_areas',(SELECT json_arrayagg(ta.name) FROM residential_areas ta WHERE ta.district_id=d.id)) as district from districts d, residential_areas r WHERE r.district_id=d.id and r.name like '%" + search + "%'")
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


