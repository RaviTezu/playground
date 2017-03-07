package plays

import "time"
import "fmt"

// StopWatch - Sorta Stop Watch
func StopWatch(ftime time.Time) {
	for {
		now := time.Now()
		remaining := ftime.Sub(now)
		hours := remaining.Hours()
		mins := (remaining - time.Duration(time.Duration(hours)*time.Hour)).Minutes()
		secs := (remaining - time.Duration(time.Duration(remaining.Minutes())*time.Minute)).Seconds()
		fmt.Printf("\rTime Remaining: %2d Hours %2d Minutes %2d Seconds", int(hours), int(mins), int(secs))
		time.Sleep(1 * time.Second)
	}
}
