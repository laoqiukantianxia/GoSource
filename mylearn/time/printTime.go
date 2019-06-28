package main

import (
	"fmt"
	"time"
)

const MyFormat = "2006/01/02 15:04:05"

func PriTime() {
	t := time.Now()
	fmt.Println("time.Now(): ", t)
	fmt.Println("time.Now().Unix(): ", t.Unix())
	year, month, date := t.Date()
	fmt.Println("time.Now().Date(): ", year, month, date)
	hour, min, sec := t.Clock()
	fmt.Println("time.Now().Clock(): ", hour, min, sec)
	fmt.Println("time.Now().Year(): ", t.Year())
	fmt.Println("time.Now().Month(): ", t.Month())

	fmt.Println("time.Now().Format(time.ANSIC): ", t.Format(time.ANSIC))
	fmt.Println("time.Now().Format(MyFormat): ", t.Format(MyFormat))

}
