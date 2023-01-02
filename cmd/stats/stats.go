package main

import (
	"fmt"
	"log"

	"gihub.com/leofigy/kollector/win"
	"github.com/yusufpapurcu/wmi"
)

func main() {
	var (
		mem       []win.Win32_PhysicalMemory
		batteries []win.Win32_Battery
	)

	query(&mem)
	query(&batteries)

	fmt.Println(mem)
	fmt.Println(batteries)
}

func query(data any) {
	q := wmi.CreateQuery(data, "")
	err := wmi.Query(q, data)
	if err != nil {
		log.Fatal(err)
	}
}
