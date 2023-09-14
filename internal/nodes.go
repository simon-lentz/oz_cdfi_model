package internal

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateNode(nodeData map[string]any, label string, session neo4j.SessionWithContext, ctx context.Context) error {
	cypherQuery := fmt.Sprintf("CREATE (n:%s) SET n += $props", label)
	params := map[string]any{
		"props": nodeData,
	}

	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				cypherQuery,
				params,
			)
			if err != nil {
				return nil, fmt.Errorf("transaction.Run() err = %+v\n", err)
			}
			return nil, nil
		},
	)
	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
	}

	return nil
}

func CreateEdges(matchParam string, session neo4j.SessionWithContext, ctx context.Context) error {
	cypherQuery := `
	MATCH (c:County), (s:State)
	WHERE c.STATE_FIPS = $STATE_FIPS AND s.STATE_FIPS = $STATE_FIPS
	CREATE (c)-[:LOCATED_IN]->(s)
`
	params := map[string]any{
		"STATE_FIPS": matchParam,
	}

	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				cypherQuery,
				params)

			if err != nil {
				return nil, fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
			}

			return nil, nil
		})

	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
	}

	return nil
}

func CreateNextEdges(matchParam string, session neo4j.SessionWithContext, ctx context.Context) error {
	cypherQuery := `
	MATCH (o:Opportunity_Zone), (c:County)
	WHERE o.COUNTY_FIPS = $COUNTY_FIPS AND c.COUNTY_FIPS = $COUNTY_FIPS
	CREATE (o)-[:LOCATED_IN]->(c)
`
	params := map[string]interface{}{
		"COUNTY_FIPS": matchParam,
	}

	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				cypherQuery,
				params)

			if err != nil {
				return nil, fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
			}

			return nil, nil
		})

	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
	}

	return nil
}
