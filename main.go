package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/simon-lentz/oz_cdfi_model/internal"
	graph "github.com/simon-lentz/oz_cdfi_model/internal/graph"
)

func main() {
	// Load config from .env and use credentials to connect to DB.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() err = %+v\n", err)
	}
	cfg := internal.LoadEnvConfig()
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		cfg.Neo4jURI,
		neo4j.BasicAuth(cfg.Neo4jUserName,
			cfg.Neo4jPassword,
			""))
	if err != nil {
		log.Fatalf("Failed to connect to Neo4j DB: %+v\n", err)
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// For each csv row (that models a new node) create a node and  write it to the graph.
	states, err := internal.GetStates("./data/state_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetStates() err = %+v\n", err)
	}
	counties, err := internal.GetCounties("./data/county_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetCounties() err = %+v\n", err)
	}
	oppZones, err := internal.GetOppZones("./data/opportunity_zone_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetOppZones() err = %+v\n", err)
	}

	for _, state := range states {
		if err := graph.CreateNode(state.StateData(), "State", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", state, err)
		}
	}
	for _, county := range counties {
		if err := graph.CreateNode(county.CountyData(), "County", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", county, err)
		}
	}
	for _, oppZone := range oppZones {
		if err := graph.CreateNode(oppZone.OppZoneData(), "Opportunity_Zone", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", oppZone, err)
		}
	}

	// Iterate over each parent node, when a parent node and child node intersect with
	// respect to their shared identifier they are linked with an edge.
	for _, state := range states {
		if err := graph.CreateEdges(graph.CountyToStateEdge, session, ctx); err != nil {
			log.Printf("Failed to create edge for %+v, err = %+v\n", state.StateFIPS, err)
		}
	}
	for _, county := range counties {
		if err := graph.CreateEdges(graph.OppZoneToCountyEdge, session, ctx); err != nil {
			log.Printf("Failed to create edge for CountyFIPS %+v, err = %+v\n", county.CountyFIPS, err)
		}
	}
}
