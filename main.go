package main 

import (
	"fmt" 
	conn "github.com/preciousnyasulu/malawi-country-data/databaseconnection"
	"github.com/preciousnyasulu/malawi-country-data/structs"


)

func main(){
	 var db  = conn.Connect()
	 var district []structs.District

	rows, err := db.Query("SELECT id, name, code, region FROM districts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var getdistrict structs.District
		err = rows.Scan(&getdistrict.Id,&getdistrict.Name,&getdistrict.Code,&getdistrict.Region)
		if err != nil {
			panic(err.Error())
		}
		district = append(district, getdistrict)
		fmt.Println(getdistrict.Id, getdistrict.Code,getdistrict.Name,getdistrict.Region)
	}
	
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}