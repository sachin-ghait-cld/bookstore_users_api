package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

// GetNow get current time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowDBFormat get current formatted time in db format
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}

// GetNowString get current formatted time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
