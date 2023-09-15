package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

// County models US Counties by their FIPS number and name.
// Child nodes are linked to their county with a LOCATED_IN edge
// that is created when the child's county FIPS field matches the
// parent's county FIPS ID.
type county struct {
	StateFIPS  string `csv:"STATE_FIPS"`
	CountyFIPS string `csv:"COUNTY_FIPS"` // Primary Key.
	CountyName string `csv:"COUNTY_NAME"`
}

// The return type map[string]any is used here to allow for
// consumption of different node types by the CreateNode
// function defined in nodes.go.
func (node *county) CountyData() map[string]any {
	countyData := map[string]any{
		"COUNTY_NAME": node.CountyName,
		"COUNTY_FIPS": node.CountyFIPS,
		"STATE_FIPS":  node.StateFIPS,
	}

	return countyData
}

// GetCounties parses county csv data stored in the data directory
// located one level above this directory (internal), it then
// populates a slice of county nodes with said data. The vanilla
// filepath parameter should be "./data/county_fips.csv"
func GetCounties(filepath string) ([]county, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%+v) err = %+v\n", filepath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("(*csv.Reader).ReadAll(file) err = %+v\n", err)
	}
	counties := []county{}
	for _, record := range records {
		counties = append(counties, county{
			StateFIPS:  record[0],
			CountyFIPS: record[1],
			CountyName: record[2],
		})
	}

	return counties, nil
}
