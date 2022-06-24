package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello world at %s! V2\n", time.Now().Local().Format("Jan 2, 2006 at 3:04:05 PM"))
}
