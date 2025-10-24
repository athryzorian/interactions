package datatypes

type Locality struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ParentCity int    `json:"parent_city"`
}
