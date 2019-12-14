package model

import "fmt"

type Machine struct {
	Name          string `json:"name"`
	IPAddress     string `json:"ipAddress"`
	VirtualSwitch string `json:"virtualSwitch"`
	MachineSpec   string `json:"machineSpec"`
}

func (i *Machine) ToString() string {
	return "Name: " + i.Name + ", IPAddress: " + i.IPAddress + ", VirtualSwitch: " + i.VirtualSwitch + ", MachineSpec: " + i.MachineSpec
}

type MachineSpec struct {
	SpecName string `json:"specName"`
	Memory   string `json:"memory"`
	Disk     string `json:"disk"`
}

func (i *MachineSpec) ToString() string {
	return "SpecName: " + i.SpecName + ", Memory: " + i.Memory + ", Disk: " + i.Disk
}

type MachineSpecs []MachineSpec

func (m MachineSpecs) Println() {
	fmt.Println("---")
	fmt.Println("MachineSpecs:")
	for _, i := range m {
		fmt.Println("   SpecName: " + i.SpecName + ", Memory: " + i.Memory + ", Disk: " + i.Disk)
	}
}
