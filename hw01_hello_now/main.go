package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timer, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("current time:", time.Now())
	fmt.Println("exact time:", timer)
}
