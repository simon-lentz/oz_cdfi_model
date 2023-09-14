package internal

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateNode(nodeData map[string]interface{}, label string, session neo4j.SessionWithContext, ctx context.Context) error {
	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				fmt.Sprintf("CREATE (n:%s) SET n += $props", label),
				map[string]any{
					"props": nodeData,
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

func NewCountyNode(node County, session neo4j.SessionWithContext, ctx context.Context) error {
	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				"CREATE (n:County { COUNTY_NAME: $COUNTY_NAME, COUNTY_FIPS: $COUNTY_FIPS, STATE_FIPS: $STATE_FIPS })",
				map[string]any{
					"COUNTY_NAME": node.CountyName,
					"COUNTY_FIPS": node.CountyFIPS,
					"STATE_FIPS":  node.StateFIPS,
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
*/
