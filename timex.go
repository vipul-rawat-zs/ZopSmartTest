package main

import (
	"fmt"
	"time"
)

var t time.Time = time.Now()

const layout = "2006-01-02 15:04:05"

func main() {

	// Time object / type
	// time.Add(d Duration)
	// f, _ := time.ParseDuration("1000000000s")
	fmt.Println(t.Format(layout), " ", t.Add(time.Duration(time.Second*1000000000)).Format(layout))

	// time.AddDate(years int, months int, days int)
	fmt.Println(t.Format(layout), " ", t.AddDate(10, 2, 2).Format(layout))

	// time.After(t time) and time.Before(t time)
	xTime := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t.After(xTime))
	fmt.Println(xTime.Before(t))

	// t.Clock(), t.Date(), t.Day()
	fmt.Println(t.Clock())
	fmt.Println(t.Date())
	fmt.Println(t.Day())

	// t.Equal(t2 Time)
	t2 := t
	fmt.Println(t.Equal(t2))
	fmt.Println(t.Equal(xTime))

	// t.In(Location)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(t.In(loc).Format(time.Kitchen))

}
