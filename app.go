package main

import (
	"flag"
	"fmt"
	"github.com/just1689/silent-k8s-cluster/cli"
	"github.com/just1689/silent-k8s-cluster/disk"
	"github.com/just1689/silent-k8s-cluster/model"
	"github.com/just1689/silent-k8s-cluster/ssh"
	"github.com/just1689/silent-k8s-cluster/virt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	routerConfigFile     = flag.String("routerConfigFile", "router-config.json", "The name of the router-config.json file")
	generateRouterConfig = flag.Bool("generateRouter", false, "Generate a router-config.json file")

	jobConfigFile     = flag.String("jobConfigFile", "job-config.json", "The name of the job-config.json file")
	generateJobConfig = flag.Bool("generateJob", false, "Generate a job-config.json file")

	machineSpecsConfigFile     = flag.String("machineSpecsConfigFile", "machine-specs-config.json", "The name of the machine-specs-config.json file")
	generateMachineSpecsConfig = flag.Bool("generateMachineSpecs", false, "Generate a machine-specs-config.json file")

	generateAll = flag.Bool("generateAll", false, "Generate router, machine spec and job files")
)

func main() {
	flag.Parse()

	checkForGenerateFlags()

	fmt.Println("Loading config")

	routerConfig := disk.LoadRouterConfig(*routerConfigFile)
	machineSpecs := disk.LoadMachineSpecsConfig(*machineSpecsConfigFile)
	job := disk.LoadJobConfig(*jobConfigFile)

	fmt.Println("   > config loaded")

	routerConfig.Println()
	machineSpecs.Println()
	job.Println()

	model.RunSpecTests(job, machineSpecs)

	fmt.Println("---")
	fmt.Println("Starting job:", job.Name)
	for _, machine := range job.Machines {
		fmt.Println("  ", machine.ToString())

		fmt.Println("  Creating Machine as spec:")
		_, spec := machineSpecs.FindByName(machine.MachineSpec)
		fmt.Println("    ", spec.ToString())
		dir := job.VMPath + machine.Name
		disk.DeleteDir(dir)
		disk.CreateDir(dir)
		err := virt.CreateVM(machine, dir, spec, job.ISOPath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("    > success")
		}

		before := cli.GetDevices(routerConfig)
		if se, err := virt.StartVM(machine); err != nil {
			logrus.Errorln(err)
			logrus.Errorln(se)
		}
		time.Sleep(5 * time.Minute)
		after := cli.GetDevices(routerConfig)
		diff := cli.FindNew(before, after)
		if len(diff) == 0 {
			logrus.Panic("Unexpectedly found 0 new devices! Check that the network adapter is correct.")
		}

		for _, r := range diff {
			logrus.Println("Device:", r.ToString())
			so, se, err := ssh.RunWithPassword(r.Address, 22, job.SSH.Username, job.SSH.Password, "ifconfig")
			logrus.Println(so, se, err)
		}
		logrus.Println("------------------")

		//TODO: set IP address of VM, reboot
		//TODO: ssh clear, ssh-copy-id copy

	}

	//TODO: initial K8s setup
	//TODO: modify ANSIBLE file
	//TODO: run ANSIBLE

}

func checkForGenerateFlags() {
	generated := false

	if *generateRouterConfig {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
		generated = true
	}

	if *generateJobConfig {
		disk.GenerateJobConfigToFile(*jobConfigFile)
		generated = true
	}

	if *generateMachineSpecsConfig {
		disk.GenerateMachineSpecsConfigToFile(*machineSpecsConfigFile)
		generated = true
	}

	if *generateAll {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
		disk.GenerateJobConfigToFile(*jobConfigFile)
		disk.GenerateMachineSpecsConfigToFile(*machineSpecsConfigFile)
	}

	if generated {
		os.Exit(0)
	}

}
