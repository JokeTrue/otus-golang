package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05 -0700 UTC")
}

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	localTime := time.Now()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("current time: %s\n", formatTime(localTime))
	fmt.Printf("exact time: %s\n", formatTime(exactTime))
}
