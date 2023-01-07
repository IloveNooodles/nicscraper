package models

type Student struct {
	Username  string `json:"username" csv:"username"`
	Name      string `json:"name" csv:"name"`
	FacultyID string `json:"tpb" csv:"tpb"`
	MajorID   string `json:"major" csv:"major"`
	Email     string `json:"email" csv:"email"`
}

type TeamsStudent struct {
	Name  string `json:"name" csv:"name"`
	NIM   string `json:"nim" csv:"nim"`
	Major string `json:"major" csv:"major"`
	Email string `json:"email" csv:"email"`
	Phone string `json:"phone" csv:"phone"`
}
