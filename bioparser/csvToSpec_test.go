package bioparser_test

import (
	"encoding/csv"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/BioSchemas/bioschemas-goweb/bioparser"
)

func TestParseSpecificationCSVToReturnSpecificationInstance(t *testing.T) {
	mdata := `title,subtitle,description,version,official type,full example,,,,
Beacon,A convention for beacon to self-describe. ,In this document we propose a simple way for a beacons to self-describe their genetic variant cardinality service for better integration with other beacons within the beacon-network. It builds upon the Beacon service API and uses existing schema.org entities and properties.		,0.2,,,,,,
schema.org / pending / External ontologies,,,,,bioschemas,,,,
Property,Expected Type,Description,Type,Type URL,BSC Description,Marginality,Cardinality,Controlled Vocabulary,Example
description,Text,A description of the item.,,,Description of this Beacon,Recommended,ONE,,
identifier,"PropertyValue or 
Text or 
URL","The identifier property represents any kind of identifier for any kind of Thing, such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes for more details.",,,Identifier,Recommended,ONE,,
`
	r := csv.NewReader(strings.NewReader(mdata))

	spec, err := bioparser.ParseSpecificationCSV(r)
	ok(t, err)

	actual := spec.SpecificationParams[0].Property
	expected := "description"

	assert(t, actual == expected, fmt.Sprintf("Expected %s but got %s", expected, actual))
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
