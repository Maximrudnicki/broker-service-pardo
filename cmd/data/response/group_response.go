package response

type GroupResponse struct {
	GroupId  string   `json:"group_id"`
	Title    string   `json:"title"`
	Students []uint32 `json:"students"`
}
