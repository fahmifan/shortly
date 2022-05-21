package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Open open sql connection using sqlite driver
func Open(fileSrc string) (*sql.DB, error) {
	return sql.Open("sqlite", fileSrc)
}
