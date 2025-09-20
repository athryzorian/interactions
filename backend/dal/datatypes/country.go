package datatypes

type Country struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CountryCode  string `json:"country_code"`
}
