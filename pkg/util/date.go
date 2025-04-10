package util

import "time"

const YMDLayout = "2006-01-02"

// String to Date parses "yyyy-mm-dd"
func ParseYMD(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// Date to String parses "yyyy-mm-dd"
func FormatYMD(t time.Time) string {
	return t.Format(YMDLayout)
}
