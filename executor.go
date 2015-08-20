package main

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"strconv"
)

const PORT_SSH_DEFAULT = 22

type Result struct {
	Message string
	Error  error
}

type SSHExecutor struct {
	Hostname string
	Port     int
	User     string
	Password string
	Commands []string
	Session  *ssh.Session
}

func (this *SSHExecutor) SetCommands(commands []string) {
	this.Commands = commands
}

func (this *SSHExecutor) Exec() *Result {
	result := &Result{}

	err := this.connect()
	if err != nil {
		result.Error = err
		return result
	}

	commands := ""
	for _, command := range this.Commands {
		commands += (command + "\n")
	}
	
	output, err1 := this.Session.Output(commands)
	if err1 != nil {
		result.Error = err1
	} else if len(output) > 0 {
		result.Message = string(output)
	}
	
	this.disconnect()
	return result
}

func (this *SSHExecutor) connect() error {
	// auths holds the detected ssh auth methods
	auths := []ssh.AuthMethod{}

	// figure out what auths are requested, what is supported
	if this.Password != "" {
		auths = append(auths, ssh.Password(this.Password))
	}

	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers))
		defer sshAgent.Close()
	}

	config := &ssh.ClientConfig{
		User: this.User,
		Auth: auths,
	}

	client, err := ssh.Dial("tcp", this.Hostname+":"+strconv.Itoa(this.Port), config)
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		return err
	}

	this.Session = session

	return nil
}

func (this *SSHExecutor) disconnect() error {
	if this.Session != nil {
		return this.Session.Close()
	}
	return nil
}

func NewExecutor(hostname, user, password string, port int) (*SSHExecutor, error) {
	if hostname == "" || user == "" {
		return nil, errors.New("error:hostname or user can not be empty")
	}
	exec := &SSHExecutor{Hostname: hostname, Port: port, User: user, Password: password}
	return exec, nil
}
