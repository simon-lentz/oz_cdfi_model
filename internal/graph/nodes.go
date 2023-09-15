package internal

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Create node uses nodeData from any of the node types defined
// in this directory along with nodeType to create a new labeled
// node with properties in the neo4j graph instance. The neo4j
// session and the context are used to initiate a db transaction and
// execute a write.
func CreateNode(nodeData map[string]any, nodeType string, session neo4j.SessionWithContext, ctx context.Context) error {
	cypherQuery := fmt.Sprintf("CREATE (n:%s) SET n += $props", nodeType) // nodeType is used as the node label.
	params := map[string]any{
		"props": nodeData,
	}

	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				cypherQuery,
				params, // params is a map of the key $props to the labeled node's properties.
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
