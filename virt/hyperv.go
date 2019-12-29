package virt

import (
	"errors"
	"github.com/just1689/silent-k8s-cluster/cli"
	"github.com/just1689/silent-k8s-cluster/model"
	"strconv"
)

func CreateVM(machine model.Machine, dir string, spec model.MachineSpec) (err error) {
	s := `New-VM -Name "` + machine.Name + `" -NewVHDPath "` + dir + "\\disk.VHDX" + `"` + ` -NewVHDSizeBytes ` + spec.Disk + ` -MemoryStartupBytes ` + spec.Memory + ` -Generation 1 `
	_, stdE := cli.ExecutePS(s)
	if stdE != "" {
		return errors.New(stdE)
	}

	s = `SET-VMProcessor -VMname "` + machine.Name + `" -count ` + strconv.Itoa(spec.CPUs)
	_, stdE = cli.ExecutePS(s)
	if stdE != "" {
		return errors.New(stdE)
	}

	return
}
