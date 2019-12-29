package disk

import (
	"encoding/json"
	"github.com/just1689/silent-k8s-cluster/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func GenerateMachineSpecsConfigToFile(filename string) {
	item := []model.MachineSpec{
		{
			Name:   "cp",
			CPUs:   4,
			Memory: "2GB",
			Disk:   "20GB",
		},
		{
			Name:   "wn",
			CPUs:   4,
			Memory: "8GB",
			Disk:   "100GB",
		},
	}
	b, err := json.Marshal(item)
	if err != nil {
		logrus.Panicln(err)
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.Println("Completed Machine Specs Generation")

}

func GenerateJobConfigToFile(filename string) {
	item := model.Job{
		Name:    "majestic-job",
		VMPath:  "C:\\vms\\",
		ISOPath: "C:\\ubuntu-silent-install.iso",
		Machines: []model.Machine{
			{
				Name:          "z-cp1",
				IPAddress:     "192.168.0.151",
				VirtualSwitch: "My Virtual Switch (ext)",
				MachineSpec:   "cp",
			},
			{
				Name:          "z-wn1",
				IPAddress:     "192.168.0.152",
				VirtualSwitch: "My Virtual Switch (ext)",
				MachineSpec:   "wn",
			},
			{
				Name:          "z-wn2",
				IPAddress:     "192.168.0.153",
				VirtualSwitch: "My Virtual Switch (ext)",
				MachineSpec:   "wn",
			},
		},
	}
	b, err := json.Marshal(item)
	if err != nil {
		logrus.Panicln(err)
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.Println("Completed Job Generation")

}

func GenerateRouterConfigToFile(filename string) {
	c := model.RouterConfig{
		Address:  "192.168.0.1:8728",
		Username: "admin",
		Password: "password",
	}
	b, err := json.Marshal(c)
	if err != nil {
		logrus.Panicln(err)
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.Println("Completed Router Config Generation")

}

func LoadRouterConfig(filename string) model.RouterConfig {
	result := model.RouterConfig{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Panicln(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Panicln(err)
	}
	return result
}

func LoadMachineSpecsConfig(filename string) model.MachineSpecs {
	result := []model.MachineSpec{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Panicln(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Panicln(err)
	}
	return result
}

func LoadJobConfig(filename string) model.Job {
	result := model.Job{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Panicln(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Panicln(err)
	}
	return result
}
