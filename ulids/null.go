package ulids

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

// Null is the nullable ulid.ULID. It support SQL & JSON serialization.
type Null struct {
	ULID
	Valid bool
}

// NewNull create new Null
func NewNull(uid ULID, valid bool) Null {
	return Null{
		ULID:  uid,
		Valid: valid,
	}
}

// NullFrom create a new Null that will always be valid.
func NullFrom(uid ULID) Null {
	return NewNull(uid, true)
}

// Value implements driver.Valuer
func (n *Null) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.ULID.String(), nil
}

// MarshalJSON implements json.Marshaler
func (t Null) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.ULID)
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Null) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.Valid = false
		return nil
	}

	err := json.Unmarshal(data, &t.ULID)
	if err != nil {
		return err
	}

	t.Valid = true
	return nil
}
