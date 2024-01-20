package modelos

import (
	"database/sql"
)

type NullTime struct {
	sql.NullTime
}

// MarshalJSON implementa a interface json.Marshaler.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return nt.Time.MarshalJSON()
}
