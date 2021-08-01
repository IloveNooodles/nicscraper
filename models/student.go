package models

type Student struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	FacultyID string `json:"tpb"`
	MajorID   string `json:"major"`
	Email     string `json:"email"`
}
