package models

import "time"

type Sesion struct {
	ID string

	StartDate Epoch
	EndDate   Epoch
}

// return the Session Duration in time.Duration format
func (s *Sesion) Duration() time.Duration {
	d := s.EndDate - s.StartDate
	return d.ToDuration()
}

// return the included epochs
func (s *Sesion) Epochs() []Epoch {
	es := []Epoch{}
	du := int(s.EndDate - s.StartDate)
	for i := 0; i < du; i++ {
		es = append(es, s.StartDate+Epoch(i))
	}
	return es
}

func (s *Sesion) UnBook(o *office) (err error) {
	return
}
