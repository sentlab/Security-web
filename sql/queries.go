// Package sql performs SQL operations
package sql

import (
	"database/sql"
	"log"

	"github.com/sentlab/Security-Audit/db"
)

type VulnBySeverity struct {
	Severity string
	Count    int
}

type TopTenVulnHosts struct {
	HostName string
	Count    int
}

type MostDangerousVulns struct {
	CVE      string
	Count    int
	Severity string
}

type VulnByType struct {
	VulnType string
	Count    int
}

type CountCVSSYear struct {
	Year  string
	Count int
}

func RunQueries(dbLocation string, tableName string) (VulnBySeverity, []TopTenVulnHosts, []MostDangerousVulns, VulnByType, []CountCVSSYear) {
	conn, err := db.InitDB(dbLocation)
	if err != nil {
		log.Fatal(err)
	}

	vulnBySeverity := vulnBySeverity(conn, tableName)
	topTenVulnHosts := topTenVulnHosts(conn, tableName)
	mostDangerousVulns := mostDangerousVulns(conn, tableName)
	vulnByType := vulnByType(conn, tableName)
	countCVSSYear := countCVSSYear(conn, tableName)

	return vulnBySeverity, topTenVulnHosts, mostDangerousVulns, vulnByType, countCVSSYear
}

func vulnBySeverity(conn *sql.DB, tableName string) VulnBySeverity {
	// Perform query and return result
	return VulnBySeverity{} // placeholder, replace with actual implementation
}

func topTenVulnHosts(conn *sql.DB, tableName string) []TopTenVulnHosts {
	// Perform query and return result
	return []TopTenVulnHosts{} // placeholder, replace with actual implementation
}

func mostDangerousVulns(conn *sql.DB, tableName string) []MostDangerousVulns {
	// Perform query and return result
	return []MostDangerousVulns{} // placeholder, replace with actual implementation
}

func vulnByType(conn *sql.DB, tableName string) VulnByType {
	// Perform query and return result
	return VulnByType{} // placeholder, replace with actual implementation
}

func countCVSSYear(conn *sql.DB, tableName string) []CountCVSSYear {
	// Perform query and return result
	return []CountCVSSYear{} // placeholder, replace with actual implementation
}
