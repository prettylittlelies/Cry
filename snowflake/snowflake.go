package snowflake

import (
	"strconv"
	"time"
)

const Epoch = 1420070400000

type ID uint64

func Parse(s string) (ID, error) {
	n, err := strconv.ParseUint(s, 10, 64)
	return ID(n), err
}

func FromTime(t time.Time) ID {
	return ID(uint64(t.UnixMilli()-Epoch) << 22)
}

func (id ID) Time() time.Time {
	return time.UnixMilli(int64(id>>22) + Epoch)
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id ID) IsZero() bool {
	return id == 0
}

func (id ID) MarshalText() ([]byte, error) {
	return strconv.AppendUint(nil, uint64(id), 10), nil
}

func (id *ID) UnmarshalText(b []byte) error {
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	*id = ID(n)
	return nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*id = 0
		return nil
	}
	if len(b) >= 2 && b[0] == '"' {
		b = b[1 : len(b)-1]
	}
	return id.UnmarshalText(b)
}
