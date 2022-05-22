package ulids

import (
	"bytes"
	"database/sql/driver"

	"github.com/oklog/ulid"
)

type Null struct {
	ULID  ulid.ULID
	Valid bool
}

func NullULIDFrom(uid ulid.ULID) Null {
	return Null{
		ULID:  uid,
		Valid: true,
	}
}

func (n *Null) Scan(src interface{}) error {
	if src == nil {
		n.ULID, n.Valid = ulid.ULID{}, false
		return nil
	}
	return n.ULID.Scan(src)
}

func (n *Null) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.ULID.String(), nil
}

func (t *Null) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.ULID.MarshalText()
}

func (t *Null) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.Valid = false
		return nil
	}

	return t.ULID.UnmarshalText(data)
}
