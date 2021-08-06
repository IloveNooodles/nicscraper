package scraper

import (
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
