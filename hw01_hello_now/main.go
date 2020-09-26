package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timer, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	t := time.Now()
	if err != nil {
		log.Fatalf("error %e", err)
		return
	}
	fmt.Println("current time:", t.Round(0))
	fmt.Println("exact time:", timer)
}
