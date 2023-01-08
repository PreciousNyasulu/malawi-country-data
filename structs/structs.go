package structs

import(

)

//
type District struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Region string `json:"region"`
}

type Constituency struct{
	Id int `json:"id"`
	Name string `json:"name"`	
	District string `json:"district"`	
}

type Ward struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Constituency string `json:"constituency"`
}

type ResidentialArea struct{
	Id int `json:"id"`
	Name  string `json:"name"`
}

type TraditionalAuthority struct {
	Id int `json:"id"`
	Name string `json:"name"`
	District string `json:"district"`
}

