package structs

type StreetData struct {
	Places []Places `json:"places"`
}

type Places struct {
	PlaceCode       string   `json:"placecode"`
	Place           string   `json:"place"`
	PlaceAKA        string   `json:"place_aka"`
	PostalCodeRange []string `json:"postalcode_range"`
	PostalCodes     []string `json:"postlcodes"`
	Municipality    string   `json:"municipality"`
	MunicipalityURL string   `json:"municipality_url"`
	Province        string   `json:"province"`
}
