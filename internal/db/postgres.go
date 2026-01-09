package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Conn *sqlx.DB
}

func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Initialize schema
	if _, err := db.Exec(Schema); err != nil {
		return nil, err
	}

	return &PostgresDB{Conn: db}, nil
}

func (db *PostgresDB) SaveResult(res *ScanResult) error {
	query := `INSERT INTO scan_results (platform, data_type, data, query) VALUES (:platform, :data_type, :data, :query)`
	_, err := db.Conn.NamedExec(query, res)
	return err
}

func (db *PostgresDB) ListResults(platform string) ([]ScanResult, error) {
	var results []ScanResult
	err := db.Conn.Select(&results, "SELECT * FROM scan_results WHERE platform = $1 ORDER BY created_at DESC", platform)
	return results, err
}

func (db *PostgresDB) Close() error {
	return db.Conn.Close()
}
