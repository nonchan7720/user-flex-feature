package utils

import "time"

func ToTime(v string) *time.Time {
	t, err := time.Parse(v, time.RFC3339)
	if err != nil {
		return nil
	}
	return Time(t)
}

func ToTimeFormat(t time.Time) string {
	return t.Format(time.RFC3339)
}
