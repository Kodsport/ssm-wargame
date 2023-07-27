package utils

import "github.com/volatiletech/null/v8"

func NullTimeToUnix(x null.Time) *int64 {
	if x.Valid {
		a := x.Time.Unix()
		return &a
	}
	return nil
}
