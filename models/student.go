package models

type Student struct {
	Username  string `json:"username" csv:"username"`
	Name      string `json:"name" csv:"name"`
	FacultyID string `json:"tpb" csv:"tpb"`
	MajorID   string `json:"major" csv:"major"`
	Email     string `json:"email" csv:"email"`
}
