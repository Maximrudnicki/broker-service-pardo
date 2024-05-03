package request

type AddStudentRequest struct {
	Token string `json:"token"`
	GroupId string `json:"group_id"`
}

type CreateGroupRequest struct {
	Token string `json:"token"`
	Title string `json:"title"`
}

type DeleteGroupRequest struct {
	Token string `json:"token"`
	GroupId string `json:"group_id"`
}

type FindGroupRequest struct {
	Token string `json:"token"`
	GroupId string `json:"group_id"`
}

type FindGroupsTeacherRequest struct {
	Token string `json:"token"`
}

type FindGroupsStudentRequest struct {
	Token string `json:"token"`
}

type RemoveStudentRequest struct {
	Token string `json:"token"`
	GroupId string `json:"group_id"`
	UserId uint32 `json:"user_id"`
}

