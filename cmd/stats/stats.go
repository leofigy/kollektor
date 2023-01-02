package main

import (
	"fmt"
	"log"

	"gihub.com/leofigy/kollector/win"
	"github.com/yusufpapurcu/wmi"
)

func main() {
	var dst []win.Win32_Battery
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		fmt.Printf("outcome %+v", v)
	}
}
