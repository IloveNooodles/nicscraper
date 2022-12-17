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

func (s Scraper) GetByNIMTeams(nim string, cvid string) (models.Student, error) {
	/* Create http client */
	client := &http.Client{}

	/* Body */
	query := map[string]string{"QueryString": nim, "DisplayQueryString": nim}

	/* FIXME: Maybe create the struct for the option */
	fieldsToFetch := [...]string{
		"Id",
		"DisplayName",
		"EmailAddresses",
		"JobTitle",
		"ImAddress",
		"UserPrincipalName",
		"ExternalDirectoryObjectId",
		"MRI",
	}

	/* FIXME: Maybe create the struct for the request object :D */
	entityRequests := map[string]interface{}{"Query": query, "EntityType": "People", "Fields": fieldsToFetch}
	body := map[string]interface{}{"EntityRequests": [...]interface{}{entityRequests}, "Cvid": cvid}

	json_body, err := json.Marshal(body)

	if err != nil {
		return models.Student{}, errors.Wrap(err, "Failed to create request body")
	}

	request, err := http.NewRequest(
		"POST",
		"https://substrate.office.com/search/api/v1/suggestions",
		bytes.NewBuffer(json_body),
	)

	if err != nil {
		return models.Student{}, errors.Wrap(err, "Failed to create new request")
	}

	/* Set Headers */
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Args.Token))
	request.Header.Set("x-client-flights", "enableAutosuggestTopHits,enableAutosuggestTopHitChannels,EnableSelfSuggestion")

	/* Making request */
	var response *http.Response

	response, err = client.Do(request)
	for err != nil || response.StatusCode >= 500 {
		response, err = client.Do(request)
	}

	defer request.Body.Close()

	/* TODO If you want to get faculty code need to do another request :/, Major needs mapping */
	/* Format data
			{
	  Groups: [
	    {
	      Suggestions: [
	        {
	          Id: "737519ec-3b0e-4350-9cfd-d32a4434d5f1@db6e1183-4c65-405c-82ce-7cd53fa6e9dc",
	          DisplayName: "Muhammad Garebaldhie Er Rahman",
	          EmailAddresses: ["13520029@mahasiswa.itb.ac.id"],
	          Phones: [
	            {
	              Number: "2500935",
	              Type: "Business",
	            },
	            {
	              Number: "82216612992",
	              Type: "Mobile",
	            },
	          ],
	          JobTitle: "Mahasiswa",
	          ImAddress: "sip:13520029@mahasiswa.itb.ac.id",
	          MRI: "8:orgid:737519ec-3b0e-4350-9cfd-d32a4434d5f1",
	          UserPrincipalName: "13520029@mahasiswa.itb.ac.id",
	          ExternalDirectoryObjectId: "737519ec-3b0e-4350-9cfd-d32a4434d5f1",
	          Text: "Muhammad Garebaldhie Er Rahman",
	          QueryText: "13520029@mahasiswa.itb.ac.id",
	          PropertyHits: ["EmailAddresses"],
	          ReferenceId: "f96e7bd4-89ea-e492-ab4e-9bd03718ffc5.2000.1",
	        },
	      ],
	      Type: "People",
	    },
	  ],
	  Instrumentation: {
	    TraceId: "f96e7bd4-89ea-e492-ab4e-9bd03718ffc5",
	  },
	}
		}
	*/

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return models.Student{}, errors.Wrap(err, "Failed to parse response")
	}

	var dataJson models.TeamsResponse

	if err := json.Unmarshal(data, &dataJson); err != nil {
		return models.Student{}, errors.Wrap(err, "Failed to parse response")
	}

	if len(dataJson.Groups[0].Suggestions) <= 0 {
		return models.Student{}, errors.Wrap(err, "possibly invalid student/NIM")
	}

	person := dataJson.Groups[0].Suggestions[0]

	student := models.Student{
		Username:  person.DisplayName,
		Name:      person.DisplayName,
		FacultyID: "placeholder",
		MajorID:   nim,
		Email:     person.ImAddress,
	}

	return student, nil
}
