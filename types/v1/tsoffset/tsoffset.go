package tsoffset_v1

import (
	"time"
)

func New(t time.Time) *TimestampOffset {
	_, offset := t.Zone()
	return &TimestampOffset{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
		Offset:  int32(offset),
	}
}

func Now() *TimestampOffset {
	return New(time.Now())
}

func (t *TimestampOffset) UTC() *TimestampOffset {
	return New(t.AsTime().UTC())
}

func (t *TimestampOffset) Local() *TimestampOffset {
	return New(t.AsTime().Local())
}

func (t *TimestampOffset) AsTime() time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos)).Local()
}
