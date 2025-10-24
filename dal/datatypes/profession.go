package datatypes

type Profession struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Logo         []byte `json:"logo"`
	IsEnabled    bool   `json:"is_enabled"`
	Description  string `json:"description"`
}
