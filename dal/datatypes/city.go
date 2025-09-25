package datatypes

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ParentState int    `json:"parent_state"`
}
