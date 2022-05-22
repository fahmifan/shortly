package sqlite

import (
	"database/sql"
	"database/sql/driver"

	"github.com/oklog/ulid"
	_ "modernc.org/sqlite"
)

// Open open sql connection using sqlite driver
func Open(fileSrc string) (*sql.DB, error) {
	return sql.Open("sqlite", fileSrc+"?cache=shared&mode=rwc&_journal_mode=WAL")
}

type ULIDStringValuer ulid.ULID

// Value implements driver valuer
func (u ULIDStringValuer) Value() (driver.Value, error) {
	return ulid.ULID(u).String(), nil
}
