package cli

import (
	"fmt"
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func ExecutePS(s string) (out, error string) {
	// choose a backend
	back := &backend.Local{}

	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	stdout, stderr, err := shell.Execute(s)
	if err != nil {
		fmt.Println(err.Error())
		if stderr == "" {
			stderr = err.Error()
		}
	}
	return stdout, stderr
}
