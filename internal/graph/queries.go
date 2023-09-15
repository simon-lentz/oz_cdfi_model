package internal

const (
	CountyToStateEdge = `
	MATCH (c:County), (s:State)
	WHERE c.STATE_FIPS = s.STATE_FIPS
	CREATE (c)-[:LOCATED_IN]->(s)
	`
	OppZoneToCountyEdge = `
	MATCH (oz:Opportunity_Zone), (co:County)
	WHERE oz.COUNTY_FIPS = co.COUNTY_FIPS
	CREATE (oz)-[:LOCATED_IN]->(co)
	`
)
