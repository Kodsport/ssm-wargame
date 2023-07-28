package skolverket

type MiniSchoolUnitResp struct {
	SchoolUnits []*MiniUnit `json:"Skolenheter"`
}

type MiniUnit struct {
	SchoolUnitCode   string `json:"Skolenhetskod"`
	MunicipalityCode string `json:"Kommunkod"`
	Status           string `json:"Status"`
	SchoolUnitName   string `json:"Skolenhetsnamn"`
	PeOrgNr          string `json:"PeOrgNr"`
}

type Unit struct {
	SkolenhetInfo struct {
		Namn           string `json:"Namn"`
		Inriktningstyp string `json:"Inriktningstyp"`
		Skolenhetstyp  string `json:"Skolenhetstyp"`
		SkolaNamn      string `json:"SkolaNamn"`
		Besoksadress   struct {
			Postnr  string `json:"Postnr"`
			GeoData struct {
				KoordinatWGS84Lat string `json:"Koordinat_WGS84_Lat"`
				KoordinatWGS84Lng string `json:"Koordinat_WGS84_Lng"`
			} `json:"GeoData"`
		} `json:"Besoksadress"`
		Skolformer []struct {
			Type string `json:"type"`
		} `json:"Skolformer"`
		Status string `json:"Status"`
	} `json:"SkolenhetInfo"`
}
