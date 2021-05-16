// SSH operations
package background

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang/src/configuration"
	"log"
	"net"
	"os"
	"strconv"
)

// use SSHagent to provide a key
func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

// use SSHconfig to provide user and define host key behaviour
func SSHConfig(server configuration.ServerConfiguration) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: server.User,
		Auth: []ssh.AuthMethod{
			SSHAgent(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

// connect to the server
func SSHConnect(server *configuration.ServerConfiguration) {
	sshclient, err := ssh.Dial("tcp", server.Server + ":" + strconv.Itoa(server.Port), SSHConfig(*server))
	if err != nil {
		server.Client = nil
	} else {
		server.Client = sshclient
	}

}

// create session
func SSHCreateSession(server *configuration.ServerConfiguration) {
	if server.Client != nil {
		session, err := server.Client.NewSession()
		if err != nil {
			log.Fatal(err)
		}
		server.Session = session
	}
}

// close session
func SSHCloseSession(server *configuration.ServerConfiguration) {
	if server.Client != nil {
		server.Session.Close()
	}
}

// close connection
func SSHClose(server *configuration.ServerConfiguration) {
	if server.Client != nil {
		server.Client.Close()
	}
}

// execute command over SSH
func SSHExec(server *configuration.ServerConfiguration, command string) string {
	if server.Client != nil {
		stdout, _ := server.Session.Output(command)
		return string(stdout)
	} else {
		return ""
	}
}
