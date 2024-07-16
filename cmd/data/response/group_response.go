package response

type GroupResponse struct {
	UserId   uint32   `json:"user_id"`
	GroupId  string   `json:"group_id"`
	Title    string   `json:"title"`
	Students []uint32 `json:"students"`
}

type StudentResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type TeacherResponse struct {
	TeacherId uint32 `json:"teacher_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type StatisticsResponse struct {
	StatId    string   `json:"statistics_id"`
	GroupId   string   `json:"group_id"`
	TeacherId uint32   `json:"teacher_id"`
	StudentId uint32   `json:"student_id"`
	Words     []uint32 `json:"words"`
}

type AddWordToUserResponse struct {
	WordId uint32 `json:"word_id"`
}

type StudentInfo struct {
	StudentId uint32 `json:student_id`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type StudentInformation struct {
	StudentId uint32          `json:student_id`
	Email     string          `json:"email"`
	Username  string          `json:"username"`
	Words     []VocabResponse `json:"words"`
}
