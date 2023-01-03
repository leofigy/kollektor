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
		query(target, "")
	}

	// special query namespace
	system := []win.CIM_OperatingSystem{}

	query(&system, "root\\cimv2")
	fmt.Println(mem)
	fmt.Println(batteries)
	fmt.Println(procs)
	fmt.Println(system)
}

func query(data any, path string) {
	q := wmi.CreateQuery(data, "")

	var err error
	if len(path) > 0 {
		fmt.Println(q)
		err = wmi.QueryNamespace(q, data, path)
	} else {
		err = wmi.Query(q, data)
	}
	if err != nil {
		log.Fatal(err)
	}
}
