package model

import "fmt"

type Job struct {
	Name     string    `json:"name"`
	VMPath   string    `json:"vm-path"`
	ISOPath  string    `json:"iso-path"`
	SSH      SSH       `json:"ssh"`
	Machines []Machine `json:"machines"`
}

func (i *Job) Println() {
	fmt.Println("---")
	fmt.Println("Job Name: ", i.Name)
	for _, m := range i.Machines {
		fmt.Println("   " + m.ToString())
	}
}
