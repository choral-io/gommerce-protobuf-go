package gender_v1

import (
	"database/sql"
)

const (
	NAME_PREFIX = "GENDER_"
	UNSPECIFIED = "UNSPECIFIED"
)

func FromString(v string) Gender {
	if g, ok := Gender_value[NAME_PREFIX+v]; ok {
		return Gender(g)
	}
	return Gender_GENDER_UNSPECIFIED
}

func FromSqlNullString(v sql.NullString) Gender {
	if v.Valid {
		return FromString(v.String)
	}
	return Gender_GENDER_UNSPECIFIED
}

func ToSqlNullString(g Gender) sql.NullString {
	if g.Number() > 0 {
		return sql.NullString{Valid: true, String: g.String()[len(NAME_PREFIX):]}
	}
	return sql.NullString{Valid: false}
}
