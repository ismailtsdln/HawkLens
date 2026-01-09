package db

import (
	"time"
)

// ScanResult matches the database schema for OSINT data
type ScanResult struct {
	ID        int64                  `db:"id" json:"id"`
	Platform  string                 `db:"platform" json:"platform"`
	DataType  string                 `db:"data_type" json:"data_type"`
	Data      map[string]interface{} `db:"data" json:"data"`
	Query     string                 `db:"query" json:"query"`
	CreatedAt time.Time              `db:"created_at" json:"created_at"`
}

// Schema is the SQL to initialize the database
const Schema = `
CREATE TABLE IF NOT EXISTS scan_results (
    id SERIAL PRIMARY KEY,
    platform TEXT NOT NULL,
    data_type TEXT NOT NULL,
    data JSONB NOT NULL,
    query TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_scan_results_platform ON scan_results(platform);
CREATE INDEX IF NOT EXISTS idx_scan_results_query ON scan_results(query);
`
