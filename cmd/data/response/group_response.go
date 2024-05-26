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

type StatisticsResponse struct {
	StatId  string `json:"statistics_id"`
	GroupId string `json:"group_id"`
	TeacherId uint32 `json:"teacher_id"`
	StudentId uint32 `json:"student_id"`
	Words []uint32 `json:"words"`
}

type AddWordToUserResponse struct {
	WordId uint32 `json:"word_id"`
}
