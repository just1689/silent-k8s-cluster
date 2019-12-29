package cli

import (
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func ExecutePS(s string) (stdout, stderr string, err error) {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()
	return shell.Execute(s)
}
