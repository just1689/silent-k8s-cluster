package main

import (
	"flag"
	"fmt"
	"github.com/just1689/silent-k8s-cluster/cli"
	"github.com/just1689/silent-k8s-cluster/disk"
)

var (
	routerConfigFile     = flag.String("routerConfigFile", "router-config.json", "The name of the router-config.json file")
	generateRouterConfig = flag.Bool("generateRouter", false, "Generate a router-config.json file")

	jobConfigFile     = flag.String("jobConfigFile", "job-config.json", "The name of the job-config.json file")
	generateJobConfig = flag.Bool("generateJob", false, "Generate a job-config.json file")
)

func main() {
	flag.Parse()

	if *generateRouterConfig {
		disk.GenerateRouterConfigToFile(*routerConfigFile)
	}

	if *generateJobConfig {
		disk.GenerateJobConfigToFile(*jobConfigFile)
	}

	config := disk.LoadRouterConfig(*routerConfigFile)
	devices := cli.GetDevices(config)
	for _, d := range devices {
		if d.IsCandidate() {
			fmt.Println("Device: ", d.ToString())
		}
	}

	//virt.CreateVM("zzz", "2GB")

}
