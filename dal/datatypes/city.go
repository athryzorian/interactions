package datatypes

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ParentState string `json:"parent_state"`
}
