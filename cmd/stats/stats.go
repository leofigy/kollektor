package main

import (
	"log"

	"gihub.com/leofigy/kollector/win"
	"github.com/yusufpapurcu/wmi"
)

func main() {
	var dst []win.Win32_PhysicalMemory
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dst {
		println(i, v.Capacity)
	}
}
