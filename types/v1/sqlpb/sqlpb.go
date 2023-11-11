package sqlpb_v1

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func FromNullType[V interface{}, T interface{}](valid bool, value V, parse func(V) *T) *T {
	if valid {
		return parse(value)
	}
	return nil
}

func FromNullBool(v sql.NullBool) *wrapperspb.BoolValue {
	return FromNullType(v.Valid, v.Bool, wrapperspb.Bool)
}

func ToNullBool(v *wrapperspb.BoolValue) sql.NullBool {
	return sql.NullBool{
		Valid: v != nil,
		Bool:  v.GetValue(),
	}
}

func FromNullFloat64(v sql.NullFloat64) *wrapperspb.DoubleValue {
	return FromNullType(v.Valid, v.Float64, wrapperspb.Double)
}

func ToNullFloat64(v *wrapperspb.DoubleValue) sql.NullFloat64 {
	return sql.NullFloat64{
		Valid:   v != nil,
		Float64: v.GetValue(),
	}
}

func FromNullInt32(v sql.NullInt32) *wrapperspb.Int32Value {
	return FromNullType(v.Valid, v.Int32, wrapperspb.Int32)
}

func ToNullInt32(v *wrapperspb.Int32Value) sql.NullInt32 {
	return sql.NullInt32{
		Valid: v != nil,
		Int32: v.GetValue(),
	}
}

func FromNullInt64(v sql.NullInt64) *wrapperspb.Int64Value {
	return FromNullType(v.Valid, v.Int64, wrapperspb.Int64)
}

func ToNullInt64(v *wrapperspb.Int64Value) sql.NullInt64 {
	return sql.NullInt64{
		Valid: v != nil,
		Int64: v.GetValue(),
	}
}

func FromNullString(v sql.NullString) *wrapperspb.StringValue {
	return FromNullType(v.Valid, v.String, wrapperspb.String)
}

func ToNullString(v *wrapperspb.StringValue) sql.NullString {
	return sql.NullString{
		Valid:  v != nil,
		String: v.GetValue(),
	}
}

func FromNullTime(v sql.NullTime) *timestamppb.Timestamp {
	return FromNullType(v.Valid, v.Time, timestamppb.New)
}

func ToNullTime(v *timestamppb.Timestamp) sql.NullTime {
	return sql.NullTime{
		Valid: v != nil && v.IsValid(),
		Time:  v.AsTime(),
	}
}
