package main

import (
	"flag"
	"github.com/just1689/silent-k8s-cluster/disk"
	"github.com/just1689/silent-k8s-cluster/model"
	"os"
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

	routerConfig := disk.LoadRouterConfig(*routerConfigFile)
	var machineSpecs model.MachineSpecs = disk.LoadMachineSpecsConfig(*machineSpecsConfigFile)
	job := disk.LoadJobConfig(*jobConfigFile)

	//devices := cli.GetDevices(routerConfig)
	//for _, d := range devices {
	//	if d.IsCandidate() {
	//		fmt.Println("Device: ", d.ToString())
	//	}
	//}

	//virt.CreateVM("zzz", "2GB")

	routerConfig.Println()
	machineSpecs.Println()
	job.Println()

}

func checkForGenerateFlags() {
	if *generateRouterConfig {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
		os.Exit(0)
	}

	if *generateJobConfig {
		disk.GenerateJobConfigToFile(*jobConfigFile)
		os.Exit(0)
	}

	if *generateMachineSpecsConfig {
		disk.GenerateMachineSpecsConfigToFile(*machineSpecsConfigFile)
		os.Exit(0)
	}

	if *generateAll {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
		disk.GenerateJobConfigToFile(*jobConfigFile)
		disk.GenerateMachineSpecsConfigToFile(*machineSpecsConfigFile)
		os.Exit(0)
	}
}
