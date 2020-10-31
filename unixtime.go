package unixtime

import (
	"time"
)

// UnixTime is alias of time.Time.
type UnixTime time.Time

// Time cast value into time.Time.
func (v UnixTime) Time() (t time.Time) {
	t = (time.Time)(v)
	return
}
