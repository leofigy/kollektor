package win

import (
	"fmt"
	"runtime"

	"github.com/yusufpapurcu/wmi"
)

type Stats struct {
	// pending to populate
	CPU     int    `json:"cpu"`
	Memory  int64  `json:"memory"`
	Storage int64  `json:"storage"`
	OS      string `json:"arch"`
	// internals
	Procs     []Win32_Process
	Mem       []Win32_PhysicalMemory
	System    CIM_OperatingSystem
	batteries []Win32_Battery
}

func (s *Stats) Refresh() error {

	var (
		procs     []Win32_Process
		mem       []Win32_PhysicalMemory
		batteries []Win32_Battery
	)

	target_structs := []any{
		&procs,
		&mem,
		&batteries,
	}

	for _, target := range target_structs {
		if err := query(target, ""); err != nil {
			return err
		}
	}

	s.Procs = procs
	s.Mem = mem
	s.batteries = batteries

	system := []CIM_OperatingSystem{}

	s.CPU = runtime.NumCPU()

	err := query(&system, "root\\cimv2")
	if len(system) > 0 {
		s.System = system[0]
		s.Memory = int64(s.System.TotalVisibleMemorySize)
		s.OS = s.System.Name
	}
	return err
}

func query(data any, path string) error {
	q := wmi.CreateQuery(data, "")

	var err error
	if len(path) > 0 {
		fmt.Println(q)
		err = wmi.QueryNamespace(q, data, path)
	} else {
		err = wmi.Query(q, data)
	}
	return err
}
