package main

import (
	"flag"
	"github.com/just1689/silent-k8s-cluster/disk"
	"github.com/just1689/silent-k8s-cluster/model"
)

var (
	routerConfigFile     = flag.String("routerConfigFile", "router-config.json", "The name of the router-config.json file")
	generateRouterConfig = flag.Bool("generateRouter", false, "Generate a router-config.json file")

	jobConfigFile     = flag.String("jobConfigFile", "job-config.json", "The name of the job-config.json file")
	generateJobConfig = flag.Bool("generateJob", false, "Generate a job-config.json file")

	machineSpecsConfigFile     = flag.String("machineSpecsConfigFile", "machine-specs-config.json", "The name of the machine-specs-config.json file")
	generateMachineSpecsConfig = flag.Bool("generateMachineSpecs", false, "Generate a machine-specs-config.json file")
)

func main() {
	flag.Parse()

	if *generateRouterConfig {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
	}

	if *generateJobConfig {
		disk.GenerateJobConfigToFile(*jobConfigFile)
	}

	if *generateMachineSpecsConfig {
		disk.GenerateMachineSpecsConfigToFile(*machineSpecsConfigFile)
	}

	var machineSpecs model.MachineSpecs
	routerConfig := disk.LoadRouterConfig(*routerConfigFile)
	machineSpecs = disk.LoadMachineSpecsConfig(*machineSpecsConfigFile)
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
