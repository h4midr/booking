package models_test

import (
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	type testCase struct {
		value  uint
		expect time.Time
		err    error
	}
	now := time.Now()
	Today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

}
