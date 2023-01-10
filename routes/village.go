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
		panic(err.Error())
	}
	defer rows.Close()
	var getvillage structs.Village
	for rows.Next() {
		err = rows.Scan(&getvillage.Id, &getvillage.Village_Name, &getvillage.District)
		if err != nil {
			panic(err.Error())
		}
		village = append(village, getvillage)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	client.IndentedJSON(http.StatusOK, village)
}
