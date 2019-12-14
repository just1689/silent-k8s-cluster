package model

type Machine struct {
	Name          string `json:"name"`
	IPAddress     string `json:"ipAddress"`
	VirtualSwitch string `json:"virtualSwitch"`
	MachineSpec   string `json:"machineSpec"`
}

type MachineSpec struct {
	SpecName string `json:"specName"`
	Memory   string `json:"memory"`
	Disk     string `json:"disk"`
}
