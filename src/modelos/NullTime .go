package modelos

import (
	"database/sql"
	"fmt"
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

func (n *NullTime) UnmarshalJSON(b []byte) error {

	n.Valid = string(b) != "null"
	e := n.Time.UnmarshalJSON(b)
	fmt.Println(e)
	return e
}
