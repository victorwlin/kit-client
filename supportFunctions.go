package main

import (
	"fmt"
	"time"
)

func (f *friends) calcNextContacts() {
	for i := range f.Friends {
		// convert string to date
		lastContactDate, err := time.Parse("2006-01-02", f.Friends[i].LastContact)
		if err != nil {
			fmt.Println(err)
		}

		// calculate recommended next contact date
		nextDate := lastContactDate.AddDate(0, 0, f.Friends[i].DesiredFreq)

		// convert date to string
		f.Friends[i].NextContact = nextDate.Format(("2006-01-02"))
	}
}
