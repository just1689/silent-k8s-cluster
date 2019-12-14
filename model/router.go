package model

import (
	"fmt"
	"strings"
)

type RouterConfig struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *RouterConfig) Println() {
	fmt.Println("---")
	fmt.Println("Router Address: " + i.Address + ", Username: " + i.Username + ", Password: " + strings.Repeat("*", len(i.Password)))
}

type DeviceLease struct {
	ActiveMacAddress string
	Hostname         string
	Address          string
	MacAddress       string
}

func (d *DeviceLease) ToString() string {
	return "ActiveMacAddress: " + d.ActiveMacAddress + ", Hostname: " + d.Hostname + ", Address: " + d.Address + ", MacAddress: " + d.MacAddress
}
func (d *DeviceLease) IsCandidate() bool {
	return d.Hostname == ""
}

func NewDeviceLeaseFromMap(in map[string]string) DeviceLease {
	result := DeviceLease{
		ActiveMacAddress: in["active-mac-address"],
		Hostname:         in["host-name"],
		Address:          in["address"],
		MacAddress:       in["mac-address"],
	}
	return result
}
