package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"malawi-country-data/structs"
)

var village []structs.Village

func GetVillages(client *gin.Context) {
	query := `SELECT v.id, v.village_name, d.name
				FROM villages v
				JOIN districts d ON d.id = v.district_id`
	rows, err := db.Query(query)
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
