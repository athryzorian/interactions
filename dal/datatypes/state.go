package datatypes

type State struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Abbreviation  string `json:"abbreviation"`
	ParentCountry int    `json:"parent_country"`
}
