package file

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mkamadeus/nicscraper/models"
)

func OutputJSON(filename string, data []models.Student) error {
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
