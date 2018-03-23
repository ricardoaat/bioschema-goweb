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

		specification, err = ParseSpecificationCSV(csv.NewReader(f))
		if err != nil {
			return specification, err
		}

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

	specification, err = ParseSpecificationCSV(csv.NewReader(res.Body))
	if err != nil {
		return specification, err
	}

	return specification, nil
}

/*
ParseSpecificationCSV reads the CSV file
	parses it into an Specification struct. The CSV second row has the Specification Info,
	and from the fourth row the mapping.
*/
func ParseSpecificationCSV(r *csv.Reader) (Specification, error) {
	var specification Specification

	r.FieldsPerRecord = 10

	csv, err := r.ReadAll()
	if err != nil {
		log.Error("An error has occurred trying to parsing the CVS reader ", err)
		return specification, err
	}

	if len(csv) <= 1 {
		log.Error("Empty file, please check its content")
		return specification, err
	}
	log.Debug("CSV lenght ", len(csv))

	specParams := make([]SpecificationParam, 0)
	var specInfo SpecificationInfo

	for i, row := range csv {

		if i == 1 {
			specInfo = SpecificationInfo{
				Title:        strings.Replace(row[0], "\n", "", -1),
				Subtitle:     strings.Replace(row[1], "\n", "", -1),
				Description:  strings.Replace(row[2], "\n", "", -1),
				Version:      strings.Replace(row[3], "\n", "", -1),
				OfficialType: strings.Replace(row[4], "\n", "", -1),
				FullExample:  strings.Replace(row[5], "\n", "", -1),
			}
		}

		if i <= 3 {
			log.Debug("Skipped ", row)
			continue
		}
		s := SpecificationParam{
			Property:             strings.Replace(row[0], "\n", "", -1),
			ExpectedType:         strings.Replace(row[1], "\n", "", -1),
			Description:          strings.Replace(row[2], "\n", "", -1),
			Type:                 strings.Replace(row[3], "\n", "", -1),
			TypeURL:              strings.Replace(row[4], "\n", "", -1),
			BscDescription:       strings.Replace(row[5], "\n", "", -1),
			Marginality:          strings.Replace(row[6], "\n", "", -1),
			Cardinality:          strings.Replace(row[7], "\n", "", -1),
			ControlledVocabulary: strings.Replace(row[8], "\n", "", -1),
			Example:              strings.Replace(row[9], "\n", "", -1),
		}
		specParams = append(specParams, s)
	}
	specification = Specification{specInfo, specParams}

	return specification, nil
}
