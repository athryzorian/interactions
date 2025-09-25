package datatypes

type Country struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CountryCode  int8   `json:"country_code"`
}
