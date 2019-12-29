package virt

import (
	"github.com/just1689/silent-k8s-cluster/cli"
	"github.com/just1689/silent-k8s-cluster/model"
	"strconv"
)

func StartVM(machine model.Machine) (stderr string, err error) {
	s := `Start-VM -Name "` + machine.Name + `"`
	_, stderr, err = cli.ExecutePS(s)
	return
}

func CreateVM(machine model.Machine, dir string, spec model.MachineSpec, isoPath string) (err error) {

	s := `Remove-VM -Name "` + machine.Name + `" -Force`
	cli.ExecutePS(s)

	s = `New-VM -Name "` + machine.Name + `" -NewVHDPath "` + dir + "\\disk.VHDX" + `"` + ` -NewVHDSizeBytes ` + spec.Disk + ` -MemoryStartupBytes ` + spec.Memory + ` -Generation 1 ` + ` -SwitchName "` + machine.VirtualSwitch + `"`
	_, _, err = cli.ExecutePS(s)
	if err != nil {
		return
	}

	s = `SET-VMProcessor -VMname "` + machine.Name + `" -count ` + strconv.Itoa(spec.CPUs)
	_, _, err = cli.ExecutePS(s)
	if err != nil {
		return
	}

	s = `SET-VM -VMname "` + machine.Name + `" -DynamicMemory -MemoryMaximumBytes ` + spec.Memory + ` -Passthru`
	_, _, err = cli.ExecutePS(s)
	if err != nil {
		return
	}

	s = `Set-VMDvdDrive -VMName ` + machine.Name + ` -Path ` + isoPath
	_, _, err = cli.ExecutePS(s)
	if err != nil {
		return
	}

	return
}
