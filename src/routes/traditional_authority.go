package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"malawi-country-data/src/structs"

	"github.com/gin-gonic/gin"
)

var traditionalauthority []structs.District


func GetTraditionalAuthorities(client *gin.Context) {
	query = `SELECT json_build_object('id', d.id, 'name', d.name, 'code', d.code, 'region', d.region, 'traditional_authorities', (SELECT json_agg(ta.name) FROM traditional_authorities ta WHERE ta.district_id = d.id)) as district 
	FROM districts d 
	JOIN traditional_authorities t ON t.district_id = d.id`
	getData(client, query)
}

func SearchTraditionalAuthorities(client *gin.Context) {
	search := client.Param("search")
		query = fmt.Sprintf(`SELECT json_build_object('id', d.id, 'name', d.name, 'code', d.code, 'region', d.region, 'traditional_authorities', (SELECT json_agg(ta.name) FROM traditional_authorities ta WHERE ta.district_id = d.id)) as district 
		FROM districts d 
		JOIN traditional_authorities t ON t.district_id = d.id 
		WHERE t.name ILIKE '%%%s%%'
							`, search)
		getData(client, query)
}

func getData(client *gin.Context, query string) {	
	rows, err := db.Query(query)
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
