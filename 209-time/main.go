package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// fmt.Println(time.Now().Local().UTC())
	layout := "2006-01-02 15:04:05 -0700 MST"
	// fmt.Println(time.Parse(layout, "2021-08-10 12:21:20.0542436 +0000 UTC"))
	d := "2020-02"
	year := strings.Split(d, "-")[0]
	month := strings.Split(d, "-")[1]
	newDate := year + "-" + month + "-02 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, newDate)

	tf := FirstDayOfMonth(t)
	tl := LastDayOfMonth(t)
	fmt.Println(tf, tl)
	fmt.Printf("%v", tl.Day()-tf.Day()+1)
	n, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 12.3456), 64)
	fmt.Printf("\n%T, %v\n", n, n)

	tt := time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(tt)

}

func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
  Takes a time.Time object and returns a time.Time object
  which is the next day.

  NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
  => time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

  NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
  => time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
//   func NextDay(t time.Time) time.Time {
// 	return t.AddDate(0, 0, 1)
//   }
