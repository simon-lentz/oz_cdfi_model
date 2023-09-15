package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

// State models US states by their FIPS number and name.
// Child nodes are linked to their state with a LOCATED_IN edge
// that is created when the child's state FIPS field matches the
// parent's state FIPS ID.
type state struct {
	Name      string `csv:"STATE_NAME"`
	StateFIPS string `csv:"STATE_FIPS"` // Primary Key.
}

// The return type map[string]any is used here to allow for
// consumption of different node types by the CreateNode
// function defined in nodes.go.
func (node *state) StateData() map[string]any {
	stateData := map[string]any{
		"STATE_NAME": node.Name,
		"STATE_FIPS": node.StateFIPS,
	}
	return stateData
}

// The filepath parameter should equal "./data/state_fips.csv"
func GetStates(filepath string) ([]state, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%+v) err = %+v\n", filepath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("(*csv.Reader).ReadAll(file) err = %v\n", err)
	}
	states := []state{}
	for _, record := range records {
		states = append(states, state{
			Name:      record[1],
			StateFIPS: record[0],
		})
	}

	return states, nil
}
