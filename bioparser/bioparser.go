//Package bioparser parses a bioschemas's CSV file into an YAML document for jekyll
package bioparser

import (
	"errors"
	"fmt"
	"os"

	"github.com/ricardoaat/bioschemas-goweb/config"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

//SpecificationParam Represents a parameter of a bioschemas spec
type SpecificationParam struct {
	Property             string   `yaml:"property"`
	ExpectedTypes        []string `yaml:"expected_types"`
	Description          string   `yaml:"description"`
	Type                 string   `yaml:"type"`
	TypeURL              string   `yaml:"type_url"`
	BscDescription       string   `yaml:"bsc_description"`
	Marginality          string   `yaml:"marginality"`
	Cardinality          string   `yaml:"cardinality"`
	ControlledVocabulary string   `yaml:"controlled_vocab"`
	Example              string   `yaml:"example"`
}

//SpecificationInfo Details for an specification struct
type SpecificationInfo struct {
	Title        string `yaml:"title"`
	Subtitle     string `yaml:"subtitle"`
	Description  string `yaml:"description"`
	Version      string `yaml:"version"`
	VersionDate  string `yaml:"version_date"`
	OfficialType string `yaml:"official_type"`
	FullExample  string `yaml:"full_example"`
}

//Specification is an object representation of the Bioschemas specification spreadsheet
//the "Bioschemas fields" sheet
type Specification struct {
	SpecificationInfo   SpecificationInfo    `yaml:"spec_info"`
	SpecificationParams []SpecificationParam `yaml:"mapping"`
}

var (
	//Specifications stores an slice of Specification structs
	Specifications []Specification
)

//Start receive a file path and an URL path for a CSV, that could be empty,
//and prints its content formated for Jekyl (YAML)
func Start(u, f string) {
	err := processCsv(u, f)
	if err != nil {
		log.Error("CSV processor failed ", err)
		return
	}
	yamlSpec, err := specsToYAML(Specifications[0])
	if err != nil {
		log.Error("Fail to get YAML string ", err)
	}
	fmt.Println(yamlSpec)

	err = specYAMLtoFile(Specifications[0])
	if err != nil {
		log.Error("Fail to create YAML file from specification ", err)
	}
}

func processCsv(u, f string) error {
	if f != "" {
		log.Info("Validating a File " + f)
		s, err := LoadFile(f)
		if err != nil {
			log.Error(err)
			return err
		}
		Specifications = append(Specifications, s)

	} else if u != "" {
		log.Info("Validating an URL " + u)
		s, err := LoadURL(u)
		if err != nil {
			log.Error(err)
			return err
		}
		Specifications = append(Specifications, s)

	} else {
		log.Warn("No valid param")
		return errors.New("Empty parameters")
	}

	return nil
}

func specYAMLtoFile(s Specification) error {
	fn := fmt.Sprintf("%s specification.yaml", s.SpecificationInfo.Title)
	fout, err := os.Create(config.Conf.Path.OutPath + fn)
	if err != nil {
		log.Error("Fail to create file. Check your file path and permissions")
		return err
	}
	defer fout.Close()

	yamlSpec, err := specsToYAML(s)
	if err != nil {
		log.Error("Fail to get YAML string")
		return err
	}

	_, err = fout.WriteString(yamlSpec)
	if err != nil {
		log.Error("Fail to write YAML file")
		return err
	}

	return nil
}

func specsToYAML(s Specification) (string, error) {
	var spec string

	d, err := yaml.Marshal(s)
	if err != nil {
		log.Error("Error marshalling specification object")
		return spec, err
	}

	spec = string(d)
	return spec, nil
}
