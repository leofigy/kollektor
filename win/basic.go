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
	Disks     []Win32_LogicalDisk
}

func (s *Stats) Refresh() error {

	var (
		procs     []Win32_Process
		mem       []Win32_PhysicalMemory
		batteries []Win32_Battery
		storage   []Win32_LogicalDisk
	)

	target_structs := []any{
		&procs,
		&mem,
		&batteries,
		&storage,
	}

	for _, target := range target_structs {
		if err := query(target, ""); err != nil {
			return err
		}
	}

	s.Procs = procs
	s.Mem = mem
	s.batteries = batteries
	s.Disks = storage

	system := []CIM_OperatingSystem{}

	s.CPU = runtime.NumCPU()

	err := query(&system, "root\\cimv2")
	if len(system) > 0 {
		s.System = system[0]
		s.Memory = int64(s.System.TotalVisibleMemorySize)
		s.OS = s.System.Name
	}

	if len(storage) > 0 {
		for _, disk := range storage {
			s.Storage += int64(disk.Size)
		}
	}

	if len(mem) > 0 {
		for _, ram := range mem {
			s.Memory += int64(ram.Capacity)
		}
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
