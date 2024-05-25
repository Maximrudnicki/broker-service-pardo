package response

type GroupResponse struct {
	GroupId  string   `json:"group_id"`
	Title    string   `json:"title"`
	Students []uint32 `json:"students"`
}

type StudentResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type AddWordToUserResponse struct {
	WordId uint32 `json:"word_id"`
}
