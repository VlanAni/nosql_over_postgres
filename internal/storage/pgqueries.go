package storage

import (
	"fmt"
)

type PGSQLGenerator struct {
}

func (PGSQLGenerator) generateCreateRequest(collectionName string) string {
	return fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (id TEXT PRIMARY KEY, payload JSONB)",
		collectionName,
	)
}

func (PGSQLGenerator) generateInsertRequest(collectionName string) string {
	return fmt.Sprintf(
		"INSERT INTO %s (id, payload) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET payload=EXCLUDED.payload",
		collectionName,
	)
}

func (PGSQLGenerator) generateGetRequest(collectionName string) string {
	return fmt.Sprintf(
		"SELECT payload FROM %s WHERE id=$1",
		collectionName,
	)
}

func (PGSQLGenerator) generateDeleteRequest(collectionName string) string {
	return fmt.Sprintf(
		"DELETE FROM %s WHERE id=$1",
		collectionName,
	)
}
