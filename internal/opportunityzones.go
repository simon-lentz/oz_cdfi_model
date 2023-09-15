package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

// OppZone models US Opportunity Zones by their FIPS number (identical
// to census tract ID). Child nodes are linked to their Opportunity Zone
// with a LOCATED_IN edge that is created when the child's county FIPS
// field matches the parent's county FIPS ID.
type oppZone struct {
	CountyFIPS  string `csv:"COUNTY_FIPS"`
	OppZoneFIPS string `csv:"OPPORTUNITY_ZONE_FIPS"` // Primary Key.
}

// The return type map[string]any is used here to allow for
// consumption of diverse node types by the CreateNode
// function defined in the nodes.go file.
func (node *oppZone) OppZoneData() map[string]any {
	oppZoneData := map[string]any{
		"COUNTY_FIPS":           node.CountyFIPS,
		"OPPORTUNITY_ZONE_FIPS": node.OppZoneFIPS,
	}
	return oppZoneData
}

// GetOppZones parses opportunity zone csv data stored in the data
// directory located one level above this directory (internal),
// it then populates a slice of county nodes with said data.
// The vanilla filepath parameter should be "./data/opportunity_zone_fips.csv"
func GetOppZones(filepath string) ([]oppZone, error) {
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
	oppZones := []oppZone{}
	for _, record := range records {
		oppZones = append(oppZones, oppZone{
			CountyFIPS:  record[0],
			OppZoneFIPS: record[1],
		})
	}

	return oppZones, nil
}
