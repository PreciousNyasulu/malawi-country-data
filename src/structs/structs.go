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
	District	 struct{
		Id                int      `json:"id"`
		Name              string   `json:"name"`
		Code              string   `json:"code"`
	}	`json:"district"`	
	Region       string   `json:"region"`
	Constituency string   `json:"constituency"`
	Wards        []string `json:"wards"`
}

type ResidentialArea struct {
	District struct{
		Id                int      `json:"id"`
		Name              string   `json:"name"`
		Code              string   `json:"code"`
	} `json:"district"`
	Region            string   `json:"region"`
	Residential_Areas []string `json:"residential_areas"`
}

type TraditionalAuthority struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
}

type Village struct {
	Id           int    `json:"id"`
	Name string `json:"name"`
	Region string `json:"region"`
	District    struct {
		Id                int      `json:"id"`
		Name              string   `json:"name"`
		Code              string   `json:"code"`
	} `json:"district"`
}

type ProblemDetail struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func InternalServerProblemDetail(detail string) ProblemDetail {
	return ProblemDetail{
		Type:   "server_error",
		Title:  "Internal Server Error",
		Detail: detail,
	}
}

func BadRequestProblemDetail(detail string) ProblemDetail {
	return ProblemDetail{
		Type:   "bad_request",
		Title:  "Bad Request",
		Detail: detail,
	}
}
