package models

type Query struct {
	QueryString        string `json:"QueryString"`
	DisplayQueryString string `json:"DisplayQueryString"`
}

type EntityRequests struct {
	Query      Query      `json:"Query"`
	EntityType EntityType `json:"EntityType"`
	Fields     []string   `json:"Fields"`
}

type RequestBody struct {
	EntityRequests []EntityRequests `json:"EntityRequests"`
	Cvid           string           `json:"Cvid"`
}

type EntityType string

const (
	PeopleEntity  EntityType = "People"
	FileEntity    EntityType = "File"
	TeamEntity    EntityType = "Team"
	ChannelEntity EntityType = "Channel"
	ChatEntity    EntityType = "Chat"
)
