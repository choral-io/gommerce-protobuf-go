package sqlpb_v1

import (
	"database/sql"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
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

func FromNullTimestamp(v sql.NullTime) *timestamppb.Timestamp {
	return FromNullTime(v)
}

func ToNullTimestamp(v *timestamppb.Timestamp) sql.NullTime {
	return ToNullTime(v)
}

type EnumType interface {
	~int32
	Type() protoreflect.EnumType
	String() string
	Number() protoreflect.EnumNumber
}

func EnumNamePrefix[E EnumType](e E) string {
	vname := e.String()
	tname := string(e.Type().Descriptor().Name())
	prefix := []byte{}
	// https://github.com/protocolbuffers/protobuf-go/blob/v1.36.4/internal/strs/strings.go#L110
	for i := 0; i < len(tname); i++ {
		c := tname[i]
		if 'A' <= c && c <= 'Z' && len(prefix) > 0 {
			prefix = append(prefix, '_')
		}
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		prefix = append(prefix, c)
	}
	prefix = append(prefix, '_')
	if strings.HasPrefix(vname, string(prefix)) {
		return string(prefix)
	}
	return ""
}

func EnumFromName[E EnumType](n string) E {
	var zero E = 0
	prefix := EnumNamePrefix(zero)
	values := (E(0)).Type().Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		if strings.EqualFold(string(v.Name()), n) || strings.EqualFold(string(v.Name()), prefix+n) {
			return E(v.Number())
		}
	}
	return zero
}

func EnumFromValue[E EnumType](v int32) E {
	return E(v)
}

func EnumToName[E EnumType](e E) string {
	vname := e.String()
	prefix := EnumNamePrefix(e)
	if strings.HasPrefix(vname, prefix) {
		return vname[len(prefix):]
	}
	return vname
}

func EnumToValue[E EnumType](e E) int32 {
	return int32(e.Number())
}

func EnumFromNullName[E EnumType](v sql.NullString) E {
	if v.Valid {
		return EnumFromName[E](v.String)
	}
	return 0
}

func EnumFromNullValue[E EnumType](v sql.NullInt32) E {
	if v.Valid {
		return EnumFromValue[E](v.Int32)
	}
	return 0
}

func EnumToNullName[E EnumType](e E) sql.NullString {
	if e == 0 {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: EnumToName(e)}
}

func EnumToNullValue[E EnumType](e E) sql.NullInt32 {
	if e == 0 {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Valid: true, Int32: EnumToValue(e)}
}
