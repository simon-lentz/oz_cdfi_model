package internal

import "os"

type config struct {
	Neo4jURI         string
	Neo4jUserName    string
	Neo4jPassword    string
	AuraInstanceID   string
	AuraInstanceName string
}

func LoadEnvConfig() *config {
	var cfg config
	cfg.Neo4jURI = os.Getenv("NEO4J_URI")
	cfg.Neo4jUserName = os.Getenv("NEO4J_USERNAME")
	cfg.Neo4jPassword = os.Getenv("NEO4J_PASSWORD")
	cfg.AuraInstanceID = os.Getenv("AURA_INSTANCEID")
	cfg.AuraInstanceName = os.Getenv("AURA_INSTANCENAME")
	return &cfg
}
