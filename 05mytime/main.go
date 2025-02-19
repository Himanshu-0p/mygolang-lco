package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time is")

	presentTime := time.Now()
    fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:05:02"))

	createdDate := time.Date(2020,time.August,10,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
