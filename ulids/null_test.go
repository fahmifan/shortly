package ulids

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/require"
)

func TestNull(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		null := NullFrom(New())
		bt, err := json.Marshal(null)
		require.NoError(t, err)
		require.Contains(t, string(bt), null.String())
	})

	t.Run("valid struct to json", func(t *testing.T) {
		zeroULID := ulid.ULID{}
		null := NewNull(ULID{zeroULID}, true)
		mock := struct {
			ID   Null   `json:"id"`
			Name string `json:"name"`
		}{
			ID:   null,
			Name: "john doe",
		}

		rawJSON := fmt.Sprintf(`{"id":"%s","name":"john doe"}`, zeroULID.String())

		bt, err := json.Marshal(mock)
		require.NoError(t, err)
		require.Equal(t, rawJSON, string(bt))

	})

	t.Run("valid null json", func(t *testing.T) {
		null := NewNull(ULID{}, false)
		bt, err := json.Marshal(null)
		require.NoError(t, err)
		require.Equal(t, string(bt), "null")
	})

	t.Run("unmarshal null json", func(t *testing.T) {
		null := Null{}
		err := json.Unmarshal([]byte("null"), &null)
		require.NoError(t, err)
		require.False(t, null.Valid)
		require.Equal(t, ulid.ULID{}.String(), null.String())
	})

	t.Run("unmarshal json", func(t *testing.T) {
		null := Null{}
		uid := New()
		bt, err := json.Marshal(uid)
		require.NoError(t, err)

		err = json.Unmarshal(bt, &null)
		require.NoError(t, err)
		require.True(t, null.Valid)
		require.Equal(t, uid.String(), null.String())
	})

	t.Run("unmarshal valid json to struct", func(t *testing.T) {
		uid := New()
		null := NewNull(uid, true)
		mock := struct {
			ID   Null   `json:"id"`
			Name string `json:"name"`
		}{
			ID:   null,
			Name: "john doe",
		}

		bt, err := json.Marshal(mock)
		require.NoError(t, err)

		mock2 := mock
		err = json.Unmarshal(bt, &mock2)
		require.NoError(t, err)
		require.Equal(t, uid.String(), mock2.ID.String())
		require.Equal(t, "john doe", mock2.Name)
	})
}
