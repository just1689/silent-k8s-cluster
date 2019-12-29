package model

import "fmt"

func RunSpecTests(job Job, specs MachineSpecs) {
	fmt.Println("---")
	fmt.Println("Checking that Job->[]Machines.SpecName exist")
	for _, m := range job.Machines {
		found, spec := specs.FindByName(m.MachineSpec)
		if !found {
			panic(fmt.Errorf("  Could not find MachineSpec named: ", m.MachineSpec))
		}
		fmt.Println("  > success: ", spec.Name)
	}

}
