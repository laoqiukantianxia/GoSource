package main

import (
	"fmt"
	"time"
)
const timeFormat = "2006-01-02 15:04:05"

func timer() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t := time.NewTimer(10 * time.Second)
	c := <- t.C
	fmt.Println(c.Format("2006-01-02 15:04:05"))
}

func afterFunc() {
	f := func() {
		fmt.Println(time.Now().Format(timeFormat))
	}
	t := time.AfterFunc(10 * time.Second, f)
	<-t.C
}

func main() {
	fmt.Println(time.Now().Format(timeFormat))
	//timer()
	afterFunc()
}
