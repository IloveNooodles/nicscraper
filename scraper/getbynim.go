package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/mkamadeus/nicscraper/models"
)

func (s Scraper) GetByNIM(nim string) {

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
		s.Failed <- nim
		return
	}

	// Set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", fmt.Sprintf("ci_session=%s;ITBnic=%s", s.Token, csrfToken))

	// Do request
	response, err := client.Do(request)
	if err != nil {
		s.Failed <- nim
		return
	}

	// Read HTML body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		s.Failed <- nim
		return
	}

	// Parse HTML and get all required data
	document := soup.HTMLParse(string(data))
	inputs := document.FindAll("input", "class", "form-control")

	ids := strings.Split(inputs[2].Attrs()["placeholder"], ",")
	facultyID := ids[0]

	majorID := ""
	if len(ids) > 1 {
		majorID = ids[1]
	}

	email := inputs[7].Attrs()["placeholder"]
	email = strings.ReplaceAll(email, "(at)", "@")
	email = strings.ReplaceAll(email, "(dot)", ".")
	email = strings.ToLower(email)

	student := &models.Student{
		Username:  inputs[1].Attrs()["placeholder"],
		Name:      inputs[3].Attrs()["placeholder"],
		FacultyID: facultyID,
		MajorID:   majorID,
		Email:     email,
	}

	log.Printf("Scraped %s", nim)

	// Input to channel
	s.Students <- student
}
