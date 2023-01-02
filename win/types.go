package win

import "time"

type Win32_Process struct {
	Name string
}

type Win32_PhysicalMemory struct {
	Capacity uint64
	Caption  string
	Name     string
}

type Win32_Battery struct {
	Availability           uint16
	BatteryRechargeTime    uint32
	BatteryStatus          uint16
	Caption                string
	Chemistry              uint16
	ConfigManagerErrorCode uint32

	ConfigManagerUserConfig     bool
	CreationClassName           string
	Description                 string
	DesignCapacity              uint32
	DesignVoltage               uint64
	DeviceID                    string
	ErrorCleared                bool
	ErrorDescription            string
	EstimatedChargeRemaining    uint16
	EstimatedRunTime            uint32
	ExpectedBatteryLife         uint32
	ExpectedLife                uint32
	FullChargeCapacity          uint32
	InstallDate                 time.Time
	LastErrorCode               uint32
	MaxRechargeTime             uint32
	Name                        string
	PNPDeviceID                 string
	PowerManagementCapabilities []int
	PowerManagementSupported    bool
	SmartBatteryVersion         string
	Status                      string
	StatusInfo                  uint16
	SystemCreationClassName     string
	SystemName                  string
	TimeOnBattery               uint32
	TimeToFullCharge            uint32
}
