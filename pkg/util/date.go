package util

import "time"

const YMDLayout = "2006-01-02"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

// String to Date parses "yyyy-mm-dd"
func ParseYMD(s string) (time.Time, error) {
	return time.ParseInLocation(YMDLayout, s, jst)
}

// Date to String parses "yyyy-mm-dd"
func FormatYMD(t time.Time) string {
	return t.Format(YMDLayout)
}
