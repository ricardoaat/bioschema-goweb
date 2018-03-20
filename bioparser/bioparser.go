package bioparser

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

//SpecificationParam Represents a parameter of a bioschemas spec
type SpecificationParam struct {
	Property             string `yaml:"name"`
	ExpectedType         string `yaml:"expected_type"`
	Description          string `yaml:"bsc_dec"`
	BscDescription       string `yaml:"sdo_desc"`
	Marginality          string `yaml:"marginality"`
	Cardinality          string `yaml:"cardinality"`
	ControlledVocabulary string `yaml:"controlled_vocab"`
}

//Specification A Bioschemas specification
type Specification struct {
	Name                string               `yaml:"name"`
	SpecificationParams []SpecificationParam `yaml:"properties"`
}

var (
	//Specifications All the parsed specifications
	Specifications []Specification
)

//Validate - Starts Validation
func Validate(f, u string) {
	if f != "" {
		log.Info("Validating a File " + f)
		s, err := LoadFile(f)
		if err != nil {
			log.Error(err)
		}
		Specifications = append(Specifications, s)
		printYaml()
	} else if u != "" {
		log.Info("Validating an URL " + u)
		s, err := LoadURL(u)
		if err != nil {
			log.Error(err)
		}
		Specifications = append(Specifications, s)
		printYaml()
	} else {
		log.Warn("Please enter valid shit")
	}
}

func printYaml() {
	d, err := yaml.Marshal(Specifications[0].SpecificationParams)
	if err != nil {
		log.Panic(err)
	}
	log.Info(string(d))
}
