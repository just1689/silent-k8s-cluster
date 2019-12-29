package cli

import (
	"github.com/just1689/silent-k8s-cluster/model"
	"gopkg.in/routeros.v2"
	"log"
)

func FindNew(before, after []model.DeviceLease) []model.DeviceLease {
	result := make([]model.DeviceLease, 0)
	var found bool
	for _, a := range after {
		found = false
		for _, b := range before {
			if a.Address == b.Address {
				found = true
				break
			}
		}
		if !found {
			result = append(result, a)
		}
	}
	return result
}

func GetDevices(config model.RouterConfig) []model.DeviceLease {

	c, err := routeros.Dial(config.Address, config.Username, config.Password)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := c.Run("/ip/dhcp-server/lease/print")
	if err != nil {
		log.Fatal(err)
	}

	devices := make([]model.DeviceLease, 0)
	for _, re := range reply.Re {
		if re.Map["active-address"] == "" {
			continue
		}
		device := model.NewDeviceLeaseFromMap(re.Map)
		devices = append(devices, device)
	}
	return devices
}
