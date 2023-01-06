package structs

import(

)

type District struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Region string `json:"region"`
}

type Constituency struct{
	Id int `json:"id"`
	Name string `json:"name"`	
}

type Ward struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

type ResidentialArea struct{
	Id int `json:"id"`
	Name  string `json:"name"`
}

type TraditionalAuthority struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

