package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type OppZone struct {
	CountyFIPS  string `csv:"COUNTY_FIPS"`
	OppZoneFIPS string `csv:"OPPORTUNITY_ZONE_FIPS"` // Counties match by this field.
}

func (node *OppZone) OppZoneData() map[string]any {
	oppZoneData := map[string]any{
		"COUNTY_FIPS":           node.CountyFIPS,
		"OPPORTUNITY_ZONE_FIPS": node.OppZoneFIPS,
	}
	return oppZoneData
}

func GetOppZones(filepath string) ([]OppZone, error) {
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
	oppZones := []OppZone{}
	for _, record := range records {
		oppZones = append(oppZones, OppZone{
			CountyFIPS:  record[0],
			OppZoneFIPS: record[1],
		})
	}

	return oppZones, nil
}
