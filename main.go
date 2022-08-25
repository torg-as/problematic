package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Printf("Hello world at %s! V2\n", time.Now().Local().Format("Jan 2, 2006 at 3:04:05 PM"))

	now := time.Now()

	fmt.Println(now)

	tok, err := time.LoadLocation("Asia/Tokyo")what

	if err != nil {
		panic(err)
	}

	fmt.Println(now.In(tok).Weekday())
}
