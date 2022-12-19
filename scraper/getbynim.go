package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/anaskhan96/soup"
	j "github.com/mkamadeus/nicscraper/json"
	"github.com/mkamadeus/nicscraper/models"
)

func (s Scraper) GetByNIM(nim string) (models.Student, error) {

	// CSRF token, can be anything.
	const csrfToken string = "banana"

	// Instantiate HTTP client
	client := &http.Client{}
	formData := fmt.Sprintf("NICitb=%s&uid=%s", csrfToken, nim)
	request, err := http.NewRequest(
		"POST",
		"https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user",
		strings.NewReader(formData),
	)

	if err != nil {
		return models.Student{}, errors.Wrap(err, "failed to make new request")
	}

	// Set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", fmt.Sprintf("ci_session=%s;ITBnic=%s", s.Args.Token, csrfToken))

	// Do request
	var response *http.Response

	response, err = client.Do(request)
	for err != nil || response.StatusCode >= 500 {
		response, err = client.Do(request)
	}

	// Read HTML body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.Student{}, errors.Wrap(err, "failed to do read response body")
	}

	// Parse HTML and get all required data
	document := soup.HTMLParse(string(data))
	inputs := document.FindAll("input", "class", "form-control")

	if len(inputs) != 10 {
		return models.Student{}, errors.New("possibly invalid student/NIM")
	}

	ids := strings.Split(inputs[2].Attrs()["placeholder"], ", ")
	facultyID := ids[0]

	majorID := ""
	if len(ids) > 1 {
		majorID = ids[1]
	}

	email := inputs[7].Attrs()["placeholder"]
	email = strings.ReplaceAll(email, "(at)", "@")
	email = strings.ReplaceAll(email, "(dot)", ".")
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)

	student := models.Student{
		Username:  inputs[1].Attrs()["placeholder"],
		Name:      inputs[3].Attrs()["placeholder"],
		FacultyID: facultyID,
		MajorID:   majorID,
		Email:     email,
	}

	// Input to channel
	return student, nil
}

func (s Scraper) GetByNIMTeams(nim string) (models.TeamsStudent, error) {
	/* Create http client */
	client := &http.Client{}

	/* Body */
	query := models.Query{
		QueryString:        nim,
		DisplayQueryString: nim,
	}

	fieldsToFetch := [...]string{
		"Id",
		"DisplayName",
		"EmailAddresses",
		"JobTitle",
		"ImAddress",
		"UserPrincipalName",
		"ExternalDirectoryObjectId",
		"Phones",
		"MRI",
	}
	entityRequests := models.EntityRequests{
		Query:      query,
		EntityType: models.PeopleEntity,
		Fields:     fieldsToFetch[:],
	}

	body := models.RequestBody{
		EntityRequests: []models.EntityRequests{entityRequests},
		Cvid:           s.Args.CVID,
	}

	json_body, err := json.Marshal(body)

	if err != nil {
		return models.TeamsStudent{}, errors.Wrap(err, "Failed to create request body")
	}

	request, err := http.NewRequest(
		"POST",
		"https://substrate.office.com/search/api/v1/suggestions",
		bytes.NewBuffer(json_body),
	)

	if err != nil {
		return models.TeamsStudent{}, errors.Wrap(err, "Failed to create new request")
	}

	/* Set Headers */
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Args.JWT))
	request.Header.Set("x-client-flights", "enableAutosuggestTopHits,enableAutosuggestTopHitChannels,EnableSelfSuggestion")

	/* Making request */
	var response *http.Response

	response, err = client.Do(request)
	for err != nil || response.StatusCode >= 500 {
		response, err = client.Do(request)
	}

	if response.StatusCode >= 400 {
		return models.TeamsStudent{}, errors.Wrap(err, "Bad request")
	}

	defer request.Body.Close()

	/* Format Data */
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return models.TeamsStudent{}, errors.Wrap(err, "Failed to parse response")
	}

	var dataJson models.TeamsResponse

	if err := json.Unmarshal(data, &dataJson); err != nil {
		return models.TeamsStudent{}, errors.Wrap(err, "Failed to parse response")
	}

	if len(dataJson.Groups) <= 0 {
		return models.TeamsStudent{}, errors.Wrap(err, "possibly invalid student/NIM")
	}

	if len(dataJson.Groups[0].Suggestions) <= 0 {
		return models.TeamsStudent{}, errors.Wrap(err, "possibly invalid student/NIM")
	}

	person := dataJson.Groups[0].Suggestions[0]

	NIM := person.UserPrincipalName[:8]
	NIMPrefixes := NIM[:3]

	var phoneNumber = "Not found"

	if len(person.Phones) > 0 {
		phoneNumber = person.Phones[len(person.Phones)-1].Number
	}

	/* TODO need to reconsider if need exact matching because of weird teams searching algorithm */
	// if len(s.Args.Prefixes.Arr) == 1 && NIMPrefixes != s.Args.Prefixes.Arr[0] {
	// 	return models.TeamsStudent{}, nil
	// }

	student := models.TeamsStudent{
		Name:  person.DisplayName,
		NIM:   NIM,
		Email: person.UserPrincipalName,
		Major: j.NIMToString[NIMPrefixes],
		Phone: phoneNumber,
	}

	return student, nil
}
