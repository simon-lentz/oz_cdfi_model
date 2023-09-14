package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/simon-lentz/oz_cdfi_model/internal"
)

type config struct {
	Neo4jURI         string
	Neo4jUserName    string
	Neo4jPassword    string
	AuraInstanceID   string
	AuraInstanceName string
}

func loadEnvConfig() *config {
	var cfg config
	cfg.Neo4jURI = os.Getenv("NEO4J_URI")
	cfg.Neo4jUserName = os.Getenv("NEO4J_USERNAME")
	cfg.Neo4jPassword = os.Getenv("NEO4J_PASSWORD")
	cfg.AuraInstanceID = os.Getenv("AURA_INSTANCEID")
	cfg.AuraInstanceName = os.Getenv("AURA_INSTANCENAME")
	return &cfg
}

func main() {
	// Load config from .env and use credentials to connect to DB.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() err = %+v\n", err)
	}
	cfg := loadEnvConfig()
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
	for _, state := range states {
		if err := internal.CreateNode(state.StateData(), "State", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", state, err)
		}
	}
	counties, err := internal.GetCounties("./data/county_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetCounties() err = %+v\n", err)
	}
	for _, county := range counties {
		if err := internal.CreateNode(county.CountyData(), "County", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", county, err)
		}
	}
	oppZones, err := internal.GetOppZones("./data/opportunity_zone_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetOppZones() err = %+v\n", err)
	}
	for _, oppZone := range oppZones {
		if err := internal.CreateNode(oppZone.OppZoneData(), "Opportunity_Zone", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", oppZone, err)
		}
	}

	// Iterate over each parent node, when a parent node and child node intersect with
	// respect to their shared identifier they are linked with an edge.
	countyToState := `
	MATCH (c:County), (s:State)
	WHERE c.STATE_FIPS = s.STATE_FIPS
	CREATE (c)-[:LOCATED_IN]->(s)
	`
	for _, state := range states {
		if err := internal.CreateEdges(countyToState, session, ctx); err != nil {
			log.Printf("Failed to create edge for %+v, err = %+v\n", state.StateFIPS, err)
		}
	}
	oppZoneToCounty := `
	MATCH (oz:Opportunity_Zone), (co:County)
	WHERE oz.COUNTY_FIPS = co.COUNTY_FIPS
	CREATE (oz)-[:LOCATED_IN]->(co)
	`
	for _, county := range counties {
		if err := internal.CreateEdges(oppZoneToCounty, session, ctx); err != nil {
			log.Printf("Failed to create edge for CountyFIPS %+v, err = %+v\n", county.CountyFIPS, err)
		}
	}
}
