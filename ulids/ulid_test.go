package ulids

import (
	"database/sql/driver"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/require"
)

func TestULIDNew(t *testing.T) {
	t.Run("should ok", func(t *testing.T) {
		uid := New()
		require.NotEmpty(t, uid.String())
		require.NotEqual(t, ulid.ULID{}.String(), uid.String())
	})
}

func TestULIDValuer(t *testing.T) {
	var _ driver.Valuer = &ULID{}

	uid := New()

	dr, err := uid.Value()
	require.NoError(t, err)

	st, ok := dr.(interface{}).(string)
	require.True(t, ok)
	require.NotEmpty(t, st)
	require.Equal(t, uid.String(), st)
}
