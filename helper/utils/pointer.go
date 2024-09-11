package utils

import "time"

func ToPointerBool(b bool) *bool {
	return &b
}

func ToPointerTime(t time.Time) *time.Time {
	return &t
}
