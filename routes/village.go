package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var village []structs.Village

func GetVillages(client *gin.Context) {
	rows, err := db.Query("SELECT v.id,v.village_name,d.name FROM villages v, districts d WHERE d.id=v.district_id")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer rows.Close()
	var getvillage structs.Village
	for rows.Next() {
		err = rows.Scan(&getvillage.Id, &getvillage.Village_Name, &getvillage.District)
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
