package models

import (
	"errors"
	"time"
)

type Epoch uint

var (
	ErrorInvalidEpoch = errors.New("Invalid Epoch")
)

const (
	/** Times Epochs per Day (24H) [0:48]
	 * 	0 0:0 	-> 0:30
	 * 	1 0:30 	-> 1:0
	 *  2 1:00 	-> 1:30
	 */
	EpochDuration_Min = 30
)

func (e *Epoch) ToDuration() time.Duration {
	return time.Duration(*e) * time.Minute
}

func (e *Epoch) ToTime() time.Time {
	// today
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
