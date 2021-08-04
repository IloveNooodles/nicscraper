package file

import (
	"io/ioutil"

	"github.com/jszwec/csvutil"
	"github.com/mkamadeus/nicscraper/models"
)

func OutputCSV(filename string, data []models.Student) error {
	file, err := csvutil.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
