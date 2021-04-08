package models

import (
	"errors"

	"time"
)

var (
	InvalidSessionDate = errors.New("Invalid Session Date")
)

type Session struct {
	ID string

	StartDate Epoch // ex booking from 0:00 to 0:30 should set 0,1 - start is and end is not
	EndDate   Epoch // ex booking from 0:00 to 3:00 should set 0, 6 or ,48
}

func newSession(ses Session) *Session {
	// Uidv4 := "uuid.NewUUID" // TODO add uuid.v4

	return &Session{
		ID:        "UUID",
		StartDate: ses.StartDate,
		EndDate:   ses.EndDate,
	}
}

// newSession create a Session for booking and checking porpose
func NewSession(start, end string) (ses Session, err error) {
	var fixed bool
	ses.StartDate, err, _ = EpochFromTime(start)
	if err != nil {
		return
	}
	ses.EndDate, err, fixed = EpochFromTime(end)
	if err != nil {
		return
	}
	if fixed {
		ses.EndDate--
	}
	if ses.EndDate-ses.StartDate < 0 {
		return ses, InvalidSessionDate
	}
	return
}

// return the Session Duration in time.Duration format
func (s *Session) Duration() time.Duration {
	d := (s.EndDate - s.StartDate) + 1
	return d.ToDuration()
}

// return the included epochs
func (s *Session) Epochs() []Epoch {
	es := []Epoch{s.StartDate}
	du := int(s.EndDate - s.StartDate)

	// Note : i must start from 1;
	for i := 1; i <= du; i++ {
		es = append(es, s.StartDate+Epoch(i))
	}
	return es
}
