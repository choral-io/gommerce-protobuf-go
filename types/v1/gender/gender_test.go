package gender_v1_test

import (
	"database/sql"
	"testing"

	gender "github.com/choral-io/gommerce-protobuf-go/types/v1/gender"
	sqlpb "github.com/choral-io/gommerce-protobuf-go/types/v1/sqlpb"
)

func TestEnumNamePrefix(t *testing.T) {
	if sqlpb.EnumNamePrefix(gender.Gender_GENDER_OTHER) != "GENDER_" {
		t.Fatalf("Expected prefix")
	}
}

func TestEnumFromName(t *testing.T) {
	if sqlpb.EnumFromName[gender.Gender]("OTHER") != gender.Gender_GENDER_OTHER {
		t.Fatalf("Expected value")
	}
	if sqlpb.EnumFromName[gender.Gender]("GENDER_OTHER") != gender.Gender_GENDER_OTHER {
		t.Fatalf("Expected value")
	}
}

func TestEnumFromValue(t *testing.T) {
	if sqlpb.EnumFromValue[gender.Gender](99) != gender.Gender_GENDER_OTHER {
		t.Fatalf("Expected value")
	}
}

func TestEnumToName(t *testing.T) {
	if sqlpb.EnumToName(gender.Gender_GENDER_OTHER) != "OTHER" {
		t.Fatalf("Expected name")
	}
}

func TestEnumToValue(t *testing.T) {
	if sqlpb.EnumToValue(gender.Gender_GENDER_OTHER) != 99 {
		t.Fatalf("Expected value")
	}
}

func TestEnumFromNullName(t *testing.T) {
	if sqlpb.EnumFromNullName[gender.Gender](sql.NullString{String: "OTHER", Valid: true}) != gender.Gender_GENDER_OTHER {
		t.Fatalf("Expected value")
	}
	if sqlpb.EnumFromNullName[gender.Gender](sql.NullString{Valid: false}) != gender.Gender_GENDER_UNSPECIFIED {
		t.Fatalf("Expected value")
	}
}

func TestEnumFromNullValue(t *testing.T) {
	if sqlpb.EnumFromNullValue[gender.Gender](sql.NullInt32{Int32: 99, Valid: true}) != gender.Gender_GENDER_OTHER {
		t.Fatalf("Expected value")
	}
	if sqlpb.EnumFromNullValue[gender.Gender](sql.NullInt32{Valid: false}) != gender.Gender_GENDER_UNSPECIFIED {
		t.Fatalf("Expected value")
	}
}

func TestEnumToNullName(t *testing.T) {
	if name := sqlpb.EnumToNullName(gender.Gender_GENDER_OTHER); !name.Valid || name.String != "OTHER" {
		t.Fatalf("Expected name")
	}
	if name := sqlpb.EnumToNullName(gender.Gender_GENDER_UNSPECIFIED); name.Valid {
		t.Fatalf("Expected name")
	}
}

func TestEnumToNullValue(t *testing.T) {
	if value := sqlpb.EnumToNullValue(gender.Gender_GENDER_OTHER); value.Valid && value.Int32 != 99 {
		t.Fatalf("Expected value")
	}
	if value := sqlpb.EnumToNullValue(gender.Gender_GENDER_UNSPECIFIED); value.Valid {
		t.Fatalf("Expected value")
	}
}
