package bioparser_test

import (
	"encoding/csv"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/ricardoaat/bioschemas-goweb/bioparser"
)

func TestParseSpecificationCSVToReturnSpecificationInstance(t *testing.T) {
	mdata := `schema.org,,,bioschemas,,,
Property,Expected Type,Description,BSC Description,MG,CN,CV
dataset,Dataset,"A dataset contained in this catalog.
Inverse property: includedInDataCatalog.",Dataset(s) served by this Beacon,Minimum,MANY,
`
	r := csv.NewReader(strings.NewReader(mdata))

	actual := bioparser.ParseSpecificationCSV(r).SpecificationParams[0].Property
	expected := "dataset"

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
