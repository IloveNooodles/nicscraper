package models

type TeamsResponse struct {
	Groups          []GroupsTeamsResponse `json:"groups"`
	Instrumentation interface{}           `json:"instrumentation"`
}

type GroupsTeamsResponse struct {
	Suggestions []SuggestionsTeamsResponse `json:"suggestions"`
	Type        string                     `json:"type"`
}

type Phone struct {
	Number string `json:"Number"`
	Type   string `json:"Type"`
}

type SuggestionsTeamsResponse struct {
	Id                        string   `json:"Id"`
	DisplayName               string   `json:"DisplayName"`
	EmailAddresses            []string `json:"EmailAddresses"`
	JobTitle                  string   `json:"JobTitle"`
	ImAddress                 string   `json:"ImAddress"`
	MRI                       string   `json:"MRI"`
	UserPrincipalName         string   `json:"UserPrincipalName"`
	ExternalDirectoryObjectId string   `json:"ExternalDirectoryObjectId"`
	Text                      string   `json:"Text"`
	QueryText                 string   `json:"QueryText"`
	ReferenceId               string   `json:"ReferenceId"`
	Phones                    []Phone  `json:"Phones"`
}
