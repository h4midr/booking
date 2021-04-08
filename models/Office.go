package models

import "errors"

var (
	ErrorOfficeNotOpendYet = errors.New("Office Not Opened Yet")
	ErrorOfficeClosed      = errors.New("Office Closed")
)

type Office interface {
	Book(Sesion) (*Sesion, error)

	IsBookable(time, duration int) bool

	UnBook(BookID string) error
}

func NewOffice() Office {
	return &office{}
}

var _ Office = office{}

type office struct {
	Epochs []int
}
