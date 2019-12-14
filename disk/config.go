package disk

import (
	"encoding/json"
	"github.com/just1689/silent-k8s-cluster/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func GenerateMachineSpecsConfigToFile(filename string) {
	item := []model.MachineSpec{
		{
			SpecName: "cp",
			Memory:   "2GB",
			Disk:     "20GB",
		},
		{
			SpecName: "wn",
			Memory:   "8GB",
			Disk:     "100GB",
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
	logrus.Println("Complete!")
	os.Exit(0)

}

func GenerateJobConfigToFile(filename string) {
	item := model.Job{
		Name: "majestic-job",
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
	logrus.Println("Complete!")
	os.Exit(0)

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
	logrus.Println("Complete!")
	os.Exit(0)

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
