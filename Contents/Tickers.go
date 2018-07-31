//Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it.

//Run code
package main

import "time"
import "fmt"

func main() {
	//Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

//When we run this program the ticker should tick 3 times before we stop it.

//$ go run tickers.go
//Tick at 2012-09-23 11:29:56.487625 -0700 PDT
//Tick at 2012-09-23 11:29:56.988063 -0700 PDT
//Tick at 2012-09-23 11:29:57.488076 -0700 PDT
//Ticker stopped
