package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow get current time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString get current formatted time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
