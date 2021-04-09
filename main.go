package main

import (
	"log"

	"github.com/h4midr/booking/models"
)

func main() {
	o, err := models.NewOffice("7:30", "22:00")
	if err != nil {
		log.Panic(err)
	}

	ses, err := models.NewSession("9:30", "10:30")
	if err != nil {
		log.Panic(err)
	}
	{
		bk, err := o.IsBookable(ses)
		log.Println(err)
		log.Printf("Ses %#v is Bookable : %v\n", ses, bk)

	}
	s, err := o.Book(ses)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Session 1 Booked")
	// log.Printf("Now the Office is %#v\n", o)

	ses2, err := models.NewSession("8:30", "9:30")
	if err != nil {
		log.Panic(err)
	}
	if bk, err := o.IsBookable(ses2); bk {
		log.Println("Session 2 is Bookable")
	} else {
		log.Println(err)
		log.Println("Session 2 is not Bookable")
	}

	_, err = o.Book(ses2)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Session 2 Booked")
	}

	//
	err = o.UnBook(s.ID)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Session1 Unbooked")

}
