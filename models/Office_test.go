package models_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/h4midr/booking/models"
)

func TestOffice(t *testing.T) {
	type testCase struct {
		openstr string
		clsstr  string

		// expected
		err error
	}
	testCases := []testCase{
		testCase{
			"7:30",
			"8:30",
			nil,
		},
		testCase{ // you mean that you want 7:30,8:00
			"7:30",
			"7:30",
			nil,
		},
		testCase{
			"7:30",
			"7:31",
			nil,
		},
		testCase{
			"7:30",
			"24:00",
			models.ErrorOfficeClosed,
		},
		testCase{
			"0:00",
			"14:00",
			models.ErrorOfficeNotOpendYet,
		},
		testCase{
			"7:30",
			"12:00",
			nil,
		},
		testCase{
			"7:30",
			"8:00",
			nil,
		},
	}

	o, err := models.NewOffice("7:30", "22:00")
	if err != nil {
		t.Error(err)
	}
	for _, v := range testCases {
		t.Run(fmt.Sprintf("Office Booking test %s to %s", v.openstr, v.clsstr), func(tt *testing.T) {
			ses, err := models.NewSession(v.openstr, v.clsstr)
			if err != nil {
				tt.Errorf("expected nil but get %v", err)
			}
			if s, err := o.Book(ses); err != v.err {
				tt.Errorf("expected %v but get %v", v.err, err)
			} else if err == nil {
				// s is booked
				tt.Logf("%v booked", s)
				err = o.UnBook(s.ID)
				if err != nil {
					tt.Errorf("expected nil but get %v", err)
				}
			}

		})
	}

}

func BenchmarkOfficeBooking(b *testing.B) {
	o, err := models.NewOffice("7:30", "22:00")
	if err != nil {
		b.Error(err)
	}
	rand.Seed(time.Now().UnixNano())
	sucsessBook := 0
	sucsessUnBook := 0
	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			ses, err := models.NewSession(fmt.Sprintf("%d:%d", rand.Intn(24), rand.Intn(60)), fmt.Sprintf("%d:%d", rand.Intn(24), rand.Intn(60)))
			if !isAllowedError(err) {
				b.Error(err)
				continue
			}
			s, err := o.Book(ses)
			if !isAllowedError(err) {
				b.Error(err)
				continue
			}
			if err == nil {
				sucsessBook++
				err = o.UnBook(s.ID)
				if err != nil {
					b.Error(err)
					continue
				} else {
					sucsessUnBook++
				}
			}

		}
	})
	b.Logf("Made %d sucsessfull Book and %d sucsessfull UnBook", sucsessBook, sucsessUnBook)
}

func isAllowedError(err error) bool {
	switch err {
	case
		nil,
		models.ErrorBlockAllocation,
		models.ErrorEpochReservedBefore,
		models.ErrorInvalidEpoch,
		models.ErrorInvalidOfficeHours,
		models.ErrorOfficeClosed,
		models.ErrorOfficeNotOpendYet:
		return true
	default:
		return false
	}
}
