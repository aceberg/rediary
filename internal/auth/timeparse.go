package auth

import (
	"errors"
	"time"
)

func timeParse(timeout string) (time.Duration, error) {
	var err error
	var t time.Duration

	suffix := timeout[len(timeout)-1]

	switch string(suffix) {
	case "h":
		t, err = time.ParseDuration(timeout)
	case "m":
		t, err = time.ParseDuration(timeout)
	case "d": //day
		t, err = time.ParseDuration(timeout[:len(timeout)-1] + "h")
		t = 24 * t
	case "M": // month
		t, err = time.ParseDuration(timeout[:len(timeout)-1] + "h")
		t = 730 * t
	default:
		err = errors.New("TimeParse: wrong time format")
	}

	return t, err
}

// ToTime - converts string (example: 3d) to time.Duration
func ToTime(s string) time.Duration {

	t, err := timeParse(s)
	if err != nil {
		t, _ = timeParse("7d")
	}

	return t
}
