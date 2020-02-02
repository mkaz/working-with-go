---
title: Date/Time
layout: page
order: 14
---

# Working with Date and Time

The [`time`](http://golang.org/pkg/time/) standard library provides the methods needed for working with dates and time in Go.


## Now

Create a date using `time.Now()`, returns a `Time` type

```go
now := time.Now()
```

Create a Unix timestamp, returns a `int64` type

```go
unix := time.Unix()
```


## Format Date

Go uses a format-by-example method for formating dates. This is odd and you may never get used to it, I still haven't. There is a base date `Monday, January 2nd, 2006 at 15:04:05` and all formats specified are based off examples using that date.

If you want a date in `YYYY-MM-DD` format, you would use `2006-01-02` as the format string.

It is debatable what is easier to remember? The sample date or the `%b %c` formats other languages use. A way to try to remember is `Month=1, Day=2, Hour=3, Minute=4, Sec=5, Year=6`

```go
fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))
```

Use a method to get any part of the date, see [`Time`](http://golang.org/pkg/time/#Time) documentation for list of available time methods.

```go
fmt.Println("Year: ", now.Year())
fmt.Println("Month: ", now.Month())
```

## Built-in Formats

The time library contains a set of [built-in constants](http://golang.org/pkg/time/#pkg-constants) to assist with standard date formatting.

```go
const (
  ANSIC       = "Mon Jan _2 15:04:05 2006"
  UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
  RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
  RFC822      = "02 Jan 06 15:04 MST"
  RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
  RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
  RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
  RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
  RFC3339     = "2006-01-02T15:04:05Z07:00"
  RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
  Kitchen     = "3:04PM"
  // Handy time stamps.
  Stamp      = "Jan _2 15:04:05"
  StampMilli = "Jan _2 15:04:05.000"
  StampMicro = "Jan _2 15:04:05.000000"
  StampNano  = "Jan _2 15:04:05.000000000"
)
```

An example using the constant formats

```go
fmt.Println(now.Format(time.RFC850))
fmt.Println(now.Format(time.RFC1123))
```

## Setting a Specific Date

A time zone is required when specifying a date, you can build a time zone using the LoadLocation, or you could also use the `time.UTC` constant for UTC.

```go
est, _ := time.LoadLocation("EST")
july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)
```

## Parse Dates

You can parse user inputted dates using known formats by specifying the format by example

```go
input_format := "1/2/2006 3:04pm"
user_str := "4/16/2014 11:38am"
user_date, err := time.Parse(input_format, user_str)
if err != nil {
	fmt.Println(">>> Error parsing date string")
}
fmt.Println("User Date: ", user_date.Format("Jan 2, 2006 @ 3:04pm"))
```

## Date/Time Comparisons

You can use `Before()`, `After()`, or `Equal()` to compare dates.

```go
if july4.Before(now) {
	fmt.Println("July 4 is before Now ")
}
```

## Date Arithmetic

The Time library also includes a [`Duration type`](https://golang.org/pkg/time/#Duration) which represents the different between two dates. The Duration type does not include days due to timezones and daylight savings.

Calculate a `Duration:`

```go
diff := now.Sub(july4)
```

Use a `Duration`:

```go
fmt.Printf("July 4 was about %d hours ago \n", diff.Hours())
```

You can add dates using Durations

```go
twodaysDiff := time.Hour * 24 * 2
twodays := now.Add(twodaysDiff)
fmt.Println("Two Days: ", twodays.Format(time.ANSIC))
```
