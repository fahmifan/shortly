package ulids

import (
	"crypto/rand"
	"database/sql/driver"
	"time"

	"github.com/oklog/ulid"
)

// ULID save the github.com/oklog/ulid.ULID as string in database.
type ULID struct {
	ulid.ULID
}

// New generate new ULID.
// Use the rand.Reader from crypto/rand, so it should be safe for concurrent usage.
func New() ULID {
	entropy := ulid.Monotonic(rand.Reader, 0)
	uid := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return ULID{uid}
}

// Value implements driver.Valuer.
func (n ULID) Value() (driver.Value, error) {
	return n.ULID.String(), nil
}
