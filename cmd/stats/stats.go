package main

import (
	"fmt"
	"log"

	"gihub.com/leofigy/kollector/win"
)

func main() {

	stats := &win.Stats{}

	err := stats.Refresh()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(stats.System)
	fmt.Println(stats.OS)
	fmt.Println(stats.Memory)
	fmt.Println(stats.Disks)

}
