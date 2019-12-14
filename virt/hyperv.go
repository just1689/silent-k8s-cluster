package virt

import (
	"errors"
	"github.com/just1689/silent-k8s-cluster/cli"
	"github.com/just1689/silent-k8s-cluster/model"
)

func CreateVM(machine model.Machine, spec model.MachineSpec) (err error) {
	s := `New-VM -Name "` + machine.Name + `" -MemoryStartupBytes ` + spec.Memory
	_, stdE := cli.ExecutePS(s)
	if stdE != "" {
		return errors.New(stdE)
	}
	return
}
