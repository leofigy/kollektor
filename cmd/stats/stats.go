package main

import (
	"fmt"
	"log"

	"gihub.com/leofigy/kollector/win"
	"github.com/yusufpapurcu/wmi"
)

func main() {
	var (
		procs     []win.Win32_Process
		mem       []win.Win32_PhysicalMemory
		batteries []win.Win32_Battery
	)

	target_structs := []any{
		&procs,
		&mem,
		&batteries,
	}

	for _, target := range target_structs {
		query(target)
	}

	fmt.Println(mem)
	fmt.Println(batteries)
	fmt.Println(procs)
}

func query(data any) {
	q := wmi.CreateQuery(data, "")
	err := wmi.Query(q, data)
	if err != nil {
		log.Fatal(err)
	}
}
