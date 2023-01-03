package win

import "time"

type Win32_Process struct {
	Name                       string
	CreationClassName          string
	Caption                    string
	CommandLine                string
	CreationDate               time.Time
	CSCreationClassName        string
	CSName                     string
	Description                string
	ExecutablePath             string
	ExecutionState             uint16
	Handle                     string
	HandleCount                uint32
	InstallDate                time.Time
	KernelModeTime             uint64
	MaximumWorkingSetSize      uint32
	MinimumWorkingSetSize      uint32
	OSCreationClassName        string
	OSName                     string
	OtherOperationCount        uint64
	OtherTransferCount         uint64
	PageFaults                 uint32
	PageFileUsage              uint32
	ParentProcessId            uint32
	PeakPageFileUsage          uint32
	PeakVirtualSize            uint64
	PeakWorkingSetSize         uint32
	Priority                   uint32
	PrivatePageCount           uint64
	ProcessId                  uint32
	QuotaNonPagedPoolUsage     uint32
	QuotaPagedPoolUsage        uint32
	QuotaPeakNonPagedPoolUsage uint32
	QuotaPeakPagedPoolUsage    uint32
	ReadOperationCount         uint64
	ReadTransferCount          uint64
	SessionId                  uint32
	Status                     string
	TerminationDate            time.Time
	ThreadCount                uint32
	UserModeTime               uint64
	VirtualSize                uint64
	WindowsVersion             string
	WorkingSetSize             uint64
	WriteOperationCount        uint64
	WriteTransferCount         uint64
}

type Win32_PhysicalMemory struct {
	Attributes           uint32
	BankLabel            string
	Capacity             uint64
	Caption              string
	ConfiguredClockSpeed uint32
	ConfiguredVoltage    uint32
	CreationClassName    string
	DataWidth            uint16
	Description          string
	DeviceLocator        string
	FormFactor           uint16
	HotSwappable         bool
	InstallDate          time.Time
	InterleaveDataDepth  uint16
	InterleavePosition   uint32
	Manufacturer         string
	MaxVoltage           uint32
	MemoryType           uint16
	MinVoltage           uint32
	Model                string
	Name                 string
	OtherIdentifyingInfo string
	PartNumber           string
	PositionInRow        uint32
	PoweredOn            bool
	Removable            bool
	Replaceable          bool
	SerialNumber         string
	SKU                  string
	SMBIOSMemoryType     uint32
	Speed                uint32
	Status               string
	Tag                  string
	TotalWidth           uint16
	TypeDetail           uint16
	Version              string
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

type CIM_OperatingSystem struct {
	Version            string
	FreePhysicalMemory uint64
	Name               string
	OSType             uint16
	Status             string
}
