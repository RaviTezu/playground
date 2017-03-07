package main

import (
	"fmt"
	"time"

	"github.com/ravitezu/playground/plays"
)

func main() {
	// The main func to use/run the functions defined in the playground package.
	// Yea, I know, this is an ugly hack. But, can't help!
	// Add a case statement here, so that the user can choose, which play to run.
	fmt.Println()
	stopTime := time.Date(2017, 03, 10, 15, 45, 0, 0, time.UTC)
	stopTime = stopTime.Add(8 * time.Hour)
	plays.StopWatch(stopTime)
}
