package ulids

import (
	"crypto/rand"
	"database/sql/driver"
	"time"

	"github.com/oklog/ulid"
)

type ULID struct {
	ulid.ULID
}

func ExampleULID() {
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3
}

func New() ULID {
	entropy := ulid.Monotonic(rand.Reader, 0)
	uid := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return ULID{uid}
}

func (n ULID) Value() (driver.Value, error) {
	return n.ULID.String(), nil
}
