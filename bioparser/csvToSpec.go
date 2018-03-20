package bioparser

import (
	"encoding/csv"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

//LoadFile loads CSV file into an spec object
func LoadFile(fp string) (Specification, error) {
	var specification Specification

	re := regexp.MustCompile(`.*\.csv$`)
	isCsv := re.MatchString(fp)
	if isCsv {
		f, err := os.Open(fp)
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()

		specification = ParseSpecificationCSV(csv.NewReader(f))

		return specification, nil
	}

	return specification, errors.New("Error that path doesn't looks like a CSV file")
}

//LoadURL loads CSV from url into an spec object
func LoadURL(u string) (Specification, error) {
	var specification Specification

	res, err := http.Get(u)
	if err != nil {
		return specification, err
	}
	defer res.Body.Close()
	specification = ParseSpecificationCSV(csv.NewReader(res.Body))

	return specification, nil
}

func ParseSpecificationCSV(r *csv.Reader) Specification {

	r.FieldsPerRecord = 7

	csv, err := r.ReadAll()
	if err != nil {
		log.Panic("An error has occurred trying to parsing the CVS reader ", err)
	}

	if len(csv) <= 1 {
		log.Error("Empty file, please check its content")
	}
	log.Debug("CSV lenght ", len(csv))

	specificationParams := make([]SpecificationParam, 0)
	for i, row := range csv {
		if i <= 1 {
			log.Debug("Skipped ", row)
			continue
		}
		s := SpecificationParam{
			strings.Replace(row[0], "\n", "", -1),
			strings.Replace(row[1], "\n", "", -1),
			strings.Replace(row[2], "\n", "", -1),
			strings.Replace(row[3], "\n", "", -1),
			strings.Replace(row[4], "\n", "", -1),
			strings.Replace(row[5], "\n", "", -1),
			strings.Replace(row[6], "\n", "", -1),
		}
		specificationParams = append(specificationParams, s)
	}

	return Specification{"le name", specificationParams}

}
