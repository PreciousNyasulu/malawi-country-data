package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var wards []structs.Ward

func GetWards(client *gin.Context) {
	rows, err := db.Query("select DISTINCT json_object('id',d.id,'district_name',d.name,'district_code',d.code,'region', d.region ,'constituency',c.name,'wards',(SELECT json_arrayagg(wa.name) FROM wards wa WHERE wa.constituency_id=c.id)) as district from districts d, constituencies c WHERE c.district_id=d.id")
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

func SearchWardWithRegion(client *gin.Context) {
	search := client.Param("search")
	rows, err := db.Query("select DISTINCT json_object('id',d.id,'district_name',d.name,'district_code',d.code,'region', d.region ,'constituency',c.name,'wards',(SELECT json_arrayagg(wa.name) FROM wards wa WHERE wa.constituency_id=c.id)) as district from districts d, constituencies c WHERE c.district_id=d.id and d.region='" + search + "'")
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
