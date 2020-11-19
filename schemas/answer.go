package schemas 

type Matching struct {
	GeoList []string `json:"geoList"`
	Cost int `json:"cost"`
	Cities []string `json:"cities"`
}