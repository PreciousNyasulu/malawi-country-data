package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var traditionalauthority []structs.District

func SearchTraditionalAuthority(client *gin.Context) {
	search := client.Param("search")
	rows, err := db.Query("select DISTINCT json_object('id',d.id,'name',d.name,'code',d.code,'region', d.region ,'traditional_authorities',(SELECT json_arrayagg(ta.name) FROM traditional_authorities ta WHERE ta.district_id=d.id)) as district from districts d, traditional_authorities t WHERE t.district_id=d.id and t.name like '%" + search + "%'")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer rows.Close()

	traditionalauthority = nil
	var getTraditionalAuthority structs.District

	for rows.Next() {
		err = rows.Scan(&newdecoder)
		json.Unmarshal([]byte(newdecoder), &getTraditionalAuthority)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		traditionalauthority = append(traditionalauthority, getTraditionalAuthority)
	}

	err = rows.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, traditionalauthority)
}
