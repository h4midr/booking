package models

import (
	"errors"
	"strings"
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
	d := uint(*e)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Add(time.Duration(d) * EpochDuration_Min * time.Minute)
}

// EpochFromTime get a time and retur the Epoch wich it's in [min-max]
func EpochFromTime(strTime string) (e Epoch, err error, fixed bool) {
	// for special case strTime == 24:00
	if strings.HasPrefix(strTime, "24:") {
		return Epoch((24 * 60) / int(EpochDuration_Min)), nil, true
	}
	// elseif the strTime was a good boy
	t, err := time.Parse("15:04", strTime)
	if err != nil {
		return
	}
	totalMins := (t.Hour()*60 + t.Minute()) % (24 * 60)
	e = Epoch(int(totalMins / int(EpochDuration_Min)))
	if (totalMins % EpochDuration_Min) == 0 {
		fixed = true
	}
	return
}
