package virt

import (
	"fmt"
	"github.com/just1689/silent-k8s-cluster/cli"
)

func CreateVM(name string, memory string) {

	s := `New-VM -Name "` + name + `" -MemoryStartupBytes ` + memory
	stdO, stdE := cli.ExecutePS(s)
	fmt.Println(stdO, stdE)

}
