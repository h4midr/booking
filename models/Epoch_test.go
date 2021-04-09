package models_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/h4midr/booking/models"
)

func TestEpoch(t *testing.T) {
	type testCase struct {
		value  uint
		expect time.Time
		err    error
	}
	now := time.Now()
	Today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	testCases := []testCase{
		testCase{
			0,
			Today,
			nil,
		},
		testCase{
			1,
			Today.Add(1 * models.EpochDuration_Min * time.Minute),
			nil,
		},
		testCase{
			12,
			Today.Add(12 * models.EpochDuration_Min * time.Minute),
			nil,
		},
		testCase{
			24,
			Today.Add(24 * models.EpochDuration_Min * time.Minute),
			nil,
		},
		testCase{
			36,
			Today.Add(36 * models.EpochDuration_Min * time.Minute),
			nil,
		},
		testCase{
			47,
			Today.Add(47 * models.EpochDuration_Min * time.Minute),
			nil,
		},
	}
	for _, v := range testCases {
		t.Run(fmt.Sprintf("Testing Epoch Converting for value of \t %d", v.value), func(tt *testing.T) {
			var E models.Epoch = models.Epoch(v.value)
			if !v.expect.Equal(E.ToTime()) {
				tt.Errorf("Expected %s got %s", v.expect.Format(time.RFC3339), v.expect.Format(time.RFC3339))
			} else {
				tt.Logf("Parsed %s", E.ToTime().Format(time.RFC3339))
			}

		})
	}
}
func TestEpochFromTime(t *testing.T) {
	type testCase struct {
		value  string
		expect models.Epoch
		err    error
	}
	testCases := []testCase{
		testCase{
			"00:00",
			0,
			nil,
		},
		testCase{
			"00:30",
			1,
			nil,
		},
		testCase{
			"00:40",
			1,
			nil,
		},
		testCase{
			"12:00",
			24,
			nil,
		},
		testCase{
			"24:00",
			48,
			nil,
		},
	}
	for _, v := range testCases {
		t.Run(fmt.Sprintf("Epoch From Time %s", v.value), func(tt *testing.T) {
			e, err, _ := models.EpochFromTime(v.value)
			if e != v.expect || err != v.err {
				tt.Errorf("Expected %d,%v got %d,%v", v.expect, v.err, e, err)
			} else {
				tt.Logf("Equals to %d", e)
			}

		})
	}
}
