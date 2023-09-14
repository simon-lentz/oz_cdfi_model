package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type County struct {
	StateFIPS  string `csv:"STATE_FIPS"`
	CountyFIPS string `csv:"COUNTY_FIPS"`
	CountyName string `csv:"COUNTY_NAME"`
}

func (node *County) CountyData() map[string]interface{} {
	countyData := map[string]interface{}{
		"COUNTY_NAME": node.CountyName,
		"COUNTY_FIPS": node.CountyFIPS,
		"STATE_FIPS":  node.StateFIPS,
	}
	return countyData
}

func GetCounties(filepath string) ([]County, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%q) err = %v", filepath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("(*csv.Reader).ReadAll(file) err = %v", err)
	}
	counties := []County{}
	for _, record := range records {
		counties = append(counties, County{
			StateFIPS:  record[0],
			CountyFIPS: record[1],
			CountyName: record[2],
		})
	}

	return counties, nil
}
