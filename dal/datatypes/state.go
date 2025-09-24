package datatypes

type State struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Abbreviation  string `json:"abbreviation"`
	ParentCountry string `json:"parent_country"`
}
