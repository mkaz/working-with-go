/**
 * time-time-time.go
 *
 * How to format, parse and calculate date and times in Go
 * See: http://golang.org/pkg/time/
 *
 */

package main

import (
	"fmt"
	"time"
)

func main() {

	// create time
	now := time.Now()
	epoch := now.Unix()

	fmt.Println("Now: ", now)
	fmt.Println("Unix Time: ", epoch)

	// go uses a format by example method to format dates
	// you "simply" show the date format you want based off
	// the date "Monday, January 2nd, 2006 at 15:04:05"
	// the hard part is remembering the standard date, not sure
	// which is better, hard to rembmer the %b %c or a date
	// the date is based off Month=1, Day=2, Hour=3, Minute=4, Sec=5, Year=6

	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	// you can also grab any part of the date you may want to use
	// See Time package for more: http://golang.org/pkg/time/#Time
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())

	// a few built in time constants exist for date formatting
	// See Time constans for more: http://golang.org/pkg/time/#pkg-constants
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	// specify a specific date
	// a time zone is required when specifying a date
	// you can build a time zone using the LoadLocation
	// you could also use time.UTC constant
	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)

	// another example formatting a date
	fmt.Println("July 4, 1776 was a ", july4.Format("Monday"))

	// Time Conditionals
	// You can use  Before() After() Equal() to compare dates
	if july4.Before(now) {
		fmt.Println("July 4 is before Now ")
	}

	// Time Addition/Subtraction
	// returns a Duration object
	diff := now.Sub(july4)

	// Duration objections unfortunately don't have days due to timezones
	// and daylight savings time which I think is a cop out but whatever
	// we can estimate using
	days := int(diff.Hours() / 24)
	fmt.Printf("July 4 was about %d days ago \n", days)

	// You can add dates using Durations
	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	// Parsing Dates
	// you can parse user inputted dates using known formats
	// by a similar format by example
	input_form := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_form, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User Date: ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

}
