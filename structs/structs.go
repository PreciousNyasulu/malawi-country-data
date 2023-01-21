package structs

type District struct {
	Id                      int      `json:"id"`
	Name                    string   `json:"name"`
	Code                    string   `json:"code"`
	Region                  string   `json:"region"`
	Traditional_Authorities []string `json:"traditional_authorities"`
}

type Constituency struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
}

type Ward struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Constituency string `json:"constituency"`
}

type ResidentialArea struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TraditionalAuthority struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
}

type Village struct {
	Id           int    `json:"id"`
	Village_Name string `json:"village_name"`
	District     string `json:"district"`
}
