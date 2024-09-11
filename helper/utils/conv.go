package utils

import "github.com/gookit/goutil/mathutil"

// ToInt64 convert value to int
// nolint: wrapcheck
func ToInt64(v any) (int64, error) {
	var data any
	switch d := v.(type) {
	case []byte:
		data = string(d)
	default:
		data = d
	}

	return mathutil.ToInt64(data)
}
