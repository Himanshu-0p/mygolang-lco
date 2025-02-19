package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time is")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:05:02"))
}
