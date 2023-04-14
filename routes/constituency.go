package routes

import (
	"fmt"
	"net/http"
	"strings"

	"malawi-country-data/structs"

	"github.com/gin-gonic/gin"
)

var constituency []structs.Constituency
var getconstituency structs.Constituency


// Gets all the country constituencies
func GetConstituencies(client *gin.Context) {
	query = `SELECT c.id, c.name, d.name
			FROM constituencies c
			JOIN districts d ON d.id=c.district_id
	`
	row, err := db.Query(query)
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	defer row.Close()

	constituency = nil
	for row.Next() {

		err = row.Scan(&getconstituency.Id, &getconstituency.Name, &getconstituency.District)
		if err != nil {
			client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
			return
		}
		constituency = append(constituency, getconstituency)
	}
	err = row.Err()
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.InternalServerProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, constituency)
}

// Gets all the constituencies in a region
func GetConstituenciesWithRegion(client *gin.Context) {
	region := strings.ToLower(client.Param("region"))

	var data []structs.Constituency
	var err error
	switch region {
	case "central":
		data, err = getConstituencyByRegion("central")
	case "northern":
		data, err = getConstituencyByRegion("northern")
	case "southern":
		data, err = getConstituencyByRegion("southern")
	default:
		client.JSON(http.StatusBadRequest, structs.BadRequestProblemDetail("Unknown region name."))
		return
	}

	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.BadRequestProblemDetail(err.Error()))
		return
	}

	client.IndentedJSON(http.StatusOK, data)
}

func getConstituencyByRegion(region string) ([]structs.Constituency, error) {
	query = fmt.Sprintf(`SELECT c.id, c.name, d.name
	FROM constituencies c
	JOIN districts d ON d.id = c.district_id
	WHERE d.region ILIKE '%s'`, region)
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	constituency = nil

	for rows.Next() {
		err = rows.Scan(&getconstituency.Id, &getconstituency.Name, &getconstituency.District)
		if err != nil {
			return nil, err
		}
		constituency = append(constituency, getconstituency)

		err = rows.Err()
		if err != nil {
			return nil, err
		}
	}

	return constituency, nil
}
