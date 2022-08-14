package date

import "time"

const (
	apiDateLayaout = "2006-01-02T15:04:05Z"
	apiDBLayaut    = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayaout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayaut)
}
