package ssh

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"log"
)

func RunWithPassword(host string, port int, username, password, cmd string) (stdout, stderr string, err error) {
	sshApi, err := sshwrapper.DefaultSshApiSetup(host, port, username, "")
	sshApi.Password = password
	err = sshApi.DefaultSshPasswordSetup()
	if err != nil {
		log.Fatal(err)
	}
	stdout, stderr, err = sshApi.Run(cmd)
	return
}
