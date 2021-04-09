package models_test

import (
	"fmt"
	"testing"

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
