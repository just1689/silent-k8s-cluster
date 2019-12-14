package main

import (
	"flag"
	"fmt"
	"github.com/just1689/silent-k8s-cluster/disk"
	"github.com/just1689/silent-k8s-cluster/virt"
	"gopkg.in/routeros.v2"
	"log"
)

var (
	generateRouterConfig = flag.Bool("generate", false, "Generate a router-config.json file")
)

func main() {
	flag.Parse()

	if *generateRouterConfig {
		disk.GenerateRouterConfigToFile()
	}

	config := disk.LoadRouterConfig()

	c, err := routeros.Dial(config.Address, config.Username, config.Password)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := c.Run("/ip/dhcp-server/lease/print")
	if err != nil {
		log.Fatal(err)
	}

	devices := make([]DeviceLease, 0)

	for _, re := range reply.Re {
		if re.Map["active-address"] == "" {
			continue
		}
		device := NewDeviceLeaseFromMap(re.Map)
		devices = append(devices, device)
	}

	for _, d := range devices {
		if d.IsCandidate() {
			fmt.Println("Device: ", d.ToString())
		}
	}

	virt.CreateVM("zzz", "2GB")

}

type DeviceLease struct {
	ActiveMacAddress string
	Hostname         string
	Address          string
	MacAddress       string
}

func (d *DeviceLease) ToString() string {
	return "AMA: " + d.ActiveMacAddress + ", HN: " + d.Hostname + " Addr: " + d.Address + ", MA: " + d.MacAddress
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
