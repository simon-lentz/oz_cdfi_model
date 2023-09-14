package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type State struct {
	Name      string `csv:"STATE_NAME"`
	StateFIPS string `csv:"STATE_FIPS"` // Counties match by this field.
}

func (node *State) StateData() map[string]any {
	stateData := map[string]any{
		"STATE_NAME": node.Name,
		"STATE_FIPS": node.StateFIPS,
	}
	return stateData
}

func GetStates(filepath string) ([]State, error) {
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
	states := []State{}
	for _, record := range records {
		states = append(states, State{
			Name:      record[1],
			StateFIPS: record[0],
		})
	}

	return states, nil
}
