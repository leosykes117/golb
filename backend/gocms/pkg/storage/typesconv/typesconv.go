package typesconv

import (
	"database/sql"
	"time"
)

// StringToNull convierte a un tipo valido para nules en la DB
func StringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	null.Valid = null.String != ""
	return null
}

// TimeToNull convierte a un tipo valido para null en la DB
func TimeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	null.Valid = !null.Time.IsZero()
	return null
}
