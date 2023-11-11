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

func ToString(g Gender) string {
	if v, ok := Gender_name[int32(g)]; ok {
		return v[len(NAME_PREFIX):]
	}
	return UNSPECIFIED
}

func ToSqlNullString(g *Gender) sql.NullString {
	if g != nil {
		return sql.NullString{Valid: true, String: ToString(*g)}
	}
	return sql.NullString{Valid: false}
}
