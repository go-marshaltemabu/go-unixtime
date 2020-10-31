package unixtime

import (
	"strconv"
	"time"
)

// MarshalJSON implements the json.Marshaler interface.
// Time is converted into UNIX epoch timestamp in seconds.
func (v UnixTime) MarshalJSON() ([]byte, error) {
	t := (time.Time)(v)
	if t.IsZero() {
		return ([]byte)("0"), nil
	}
	return ([]byte)(strconv.FormatInt(t.Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a integer number in UNIX epoch seconds.
func (v *UnixTime) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "0" {
		*v = UnixTime{}
		return nil
	} else if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	i64, err := strconv.ParseInt(s, 10, 64)
	if nil != err {
		return err
	}
	*v = (UnixTime)(time.Unix(i64, 0))
	return nil
}
