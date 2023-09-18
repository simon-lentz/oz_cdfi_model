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
	// This edge is troublesome...
	TestDataToOppZone = `
	MATCH (td:TestData), (oz:Opportunity_Zone)
	WHERE td.CENSUS_TRACT_FIPS_NUMER = oz.OPPORTUNITY_ZONE_FIPS
	CREATE (td)-[:DESCRIBES]->(oz)
	`
)
