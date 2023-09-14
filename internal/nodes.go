package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func States() {
	states, err := GetStates("./data/state_fips.csv")
	if err != nil {
		log.Fatalf("GetStates() err = %+v", err)
	}

	fmt.Println(states)
}

func NewStateNode(node State, session neo4j.SessionWithContext, ctx context.Context) error {
	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				"CREATE (n:State { STATE_NAME: $STATE_NAME, STATE_FIPS: $STATE_FIPS })",
				map[string]any{
					"STATE_NAME": node.Name,
					"STATE_FIPS": node.StateFIPS,
				},
			)
			if err != nil {
				return nil, fmt.Errorf("transaction.Run() err = %+v", err)
			}
			return nil, nil
		},
	)
	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v", err)
	}

	return nil
}

/*
func TestNode(session neo4j.SessionWithContext, ctx context.Context) error {
	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(
				ctx,
				"CREATE (testnode:TestNode { Test: $test })",
				map[string]any{
					"test": "testing testing 1, 2, 3...",
				},
			)
			if err != nil {
				return nil, fmt.Errorf("transaction.Run() err = %+v", err)
			}
			record := result.Record()
			fmt.Println(record.AsMap())
			return nil, nil
		},
	)
	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v", err)
	}

	return nil
}
*/
