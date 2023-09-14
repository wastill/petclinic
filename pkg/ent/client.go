package ent

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3" // use sqlite3 for ent
)

// 初始化数据化驱动
func InitDatabaseDriver(dbType string, connectURI string) (*entsql.Driver, error) {
	var dbDriver *entsql.Driver
	switch dbType {
	case dialect.Postgres:
		rawdb, err := sql.Open("pgx", connectURI)
		if err != nil {
			return nil, fmt.Errorf("ent open error: %w", err)
		}
		// Create an ent.Driver from `db`.
		dbDriver = entsql.OpenDB(dbType, rawdb)
	case dialect.SQLite:
		rawdb, err := sql.Open("sqlite3", connectURI)
		if err != nil {
			return nil, fmt.Errorf("ent open error: %w", err)
		}
		// Create an ent.Driver from `db`.
		dbDriver = entsql.OpenDB(dbType, rawdb)
	case dialect.MySQL:
		rawdb, err := sql.Open("mysql", connectURI)
		if err != nil {
			return nil, fmt.Errorf("ent open error: %w", err)
		}
		dbDriver = entsql.OpenDB(dbType, rawdb)
	default:
		return nil, fmt.Errorf("Database type: %v not support", dbType)
	}

	return dbDriver, nil
}
