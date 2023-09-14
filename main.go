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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() err = %+v", err)
	}
	cfg := loadEnvConfig()
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		cfg.Neo4jURI,
		neo4j.BasicAuth(cfg.Neo4jUserName,
			cfg.Neo4jPassword,
			""))
	if err != nil {
		log.Fatalf("Failed to connect to Neo4j DB: %+v", err)
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	states, err := internal.GetStates("./data/state_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetStates() err = %+v", err)
	}
	for _, state := range states {
		if err := internal.CreateNode(state.StateData(), "State", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v", state, err)
		}
	}

	counties, err := internal.GetCounties("./data/county_fips.csv")
	if err != nil {
		log.Fatalf("internal.GetCounties() err = %+v", err)
	}
	for _, county := range counties {
		if err := internal.CreateNode(county.CountyData(), "County", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v", county, err)
		}
	}
}
