package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/preciousnyasulu/malawi-country-data/structs"
)

var constituency []structs.Constituency
var getconstituency structs.Constituency

// Gets all the country constituencies
func GetConstituencies(client *gin.Context) {
	row, err := db.Query("SELECT c.id, c.name, d.name FROM constituencies c, districts d WHERE d.id=c.district_id")
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	constituency = nil
	for row.Next() {

		err = row.Scan(&getconstituency.Id, &getconstituency.Name, &getconstituency.District)
		if err != nil {
			panic(err.Error())
		}
		constituency = append(constituency, getconstituency)
	}
	err = row.Err()
	if err != nil {
		panic(err.Error())
	}
	client.IndentedJSON(http.StatusOK, constituency)
}

// Gets all the constituencies in a region
func GetConstituenciesWithRegion(client *gin.Context) {
	region := strings.ToLower(client.Param("region"))

	switch region {
	case "central":
		client.IndentedJSON(http.StatusOK, getConstituencyByRegion("central"))
	case "northern":
		client.IndentedJSON(http.StatusOK, getConstituencyByRegion("northern"))
	case "southern":
		client.IndentedJSON(http.StatusOK, getConstituencyByRegion("southern"))
	default:
		client.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Unknown region name."})
	}
}

func getConstituencyByRegion(region string) []structs.Constituency {
	rows, err := db.Query("SELECT c.id ,c.name ,d.name FROM constituencies c, districts d WHERE d.id=c.district_id AND d.region='" + region + "'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	constituency = nil

	for rows.Next() {
		err = rows.Scan(&getconstituency.Id, &getconstituency.Name, &getconstituency.District)
		if err != nil {
			panic(err.Error())
		}
		constituency = append(constituency, getconstituency)

		err = rows.Err()
		if err != nil {
			panic(err.Error())
		}
	}

	return constituency
}
