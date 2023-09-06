package controller

import (
	"fmt"
	"net/http"

	"malawi-country-data/src/structs"

	"github.com/gin-gonic/gin"
)

var constituency []structs.Constituency
var getconstituency structs.Constituency

// Gets all the country constituencies
// @Summary 		Gets all the country constituencies
// @Description 	Gets all the constituencies
// @Tags 			Constituencies
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Constituency{}
// @Router 			/Constituencies [get]
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

// Gets all the constituencies in the northern region
// @Summary 		Gets all northern region constituencies
// @Description 	Gets all the constituencies in the northern region
// @Tags 			Constituencies
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Constituency{}
// @Router 			/Constituencies/Region/Northern [get]
func GetNorthernConstituencies(client *gin.Context) {
	data, err := getConstituencyByRegion("northern")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.BadRequestProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, data)
}

// Gets all the constituencies in the southern region
// @Summary 		Gets all southern region constituencies
// @Description 	Gets all the constituencies in the southern region
// @Tags 			Constituencies
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Constituency{}
// @Router 			/Constituencies/Region/Southern [get]
func GetSouthernConstituencies(client *gin.Context) {
	data, err := getConstituencyByRegion("southern")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.BadRequestProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, data)
}

// Gets all the constituencies in the Central region
// @Summary 		Gets all Central region constituencies
// @Description 	Gets all the constituencies in the Central region
// @Tags 			Constituencies
// @Accept 			json
// @Produce			application/json
// @Success 		200 {object} structs.Constituency{}
// @Router 			/Constituencies/Region/Central [get]
func GetCentralConstituencies(client *gin.Context) {
	data, err := getConstituencyByRegion("central")
	if err != nil {
		client.JSON(http.StatusInternalServerError, structs.BadRequestProblemDetail(err.Error()))
		return
	}
	client.IndentedJSON(http.StatusOK, data)
}

// func GetConstituenciesWithRegion(client *gin.Context) {
// 	region := strings.ToLower(client.Param("region"))

// 	var data []structs.Constituency
// 	var err error
// 	switch region {
// 	case "central":
// 		data, err = getConstituencyByRegion("central")
// 	case "northern":
// 		data, err = getConstituencyByRegion("northern")
// 	case "southern":
// 		data, err = getConstituencyByRegion("southern")
// 	default:
// 		client.JSON(http.StatusBadRequest, structs.BadRequestProblemDetail("Unknown region name."))
// 		return
// 	}

// 	if err != nil {
// 		client.JSON(http.StatusInternalServerError, structs.BadRequestProblemDetail(err.Error()))
// 		return
// 	}

// 	client.IndentedJSON(http.StatusOK, data)
// }

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
