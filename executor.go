package main

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/appleboy/easyssh-proxy"
)

const PORT_SSH_DEFAULT = 22

var (
	ErrHostUsername error = errors.New("error: Hostname or user can not be empty.")
)

type ErrConnect struct {
	hostname string
	reason   error
}

func (this *ErrConnect) Error() string {
	return "error: Can't connect to host " + this.hostname + ", reason: " + this.reason.Error()
}

type Result struct {
	Message string
	Error   error
}

type SSHExecutor struct {
	Hostname string
	Port     int
	User     string
	Password string
	Commands []string
	Session  *easyssh.MakeConfig
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

	commands := strings.Join(this.Commands, ";")

	stdout, stderr, _, err1 := this.Session.Run(commands)
	if err1 != nil || stderr != "" {
		result.Error = errors.New(stderr)
	} else if stdout != "" {
		result.Message = stdout
	}

	this.disconnect()
	return result
}

func (this *SSHExecutor) connect() error {
	homeDir, _ := os.UserHomeDir()
	this.Session = &easyssh.MakeConfig{
		User:     this.User,
		Server:   this.Hostname,
		Password: this.Password,
		KeyPath:  homeDir + "/.ssh/id_rsa",
		Port:     "22",
		Timeout:  60 * time.Second,
	}
	return nil
}

func (this *SSHExecutor) disconnect() error {
	if this.Session != nil {
		this.Session = nil
	}
	return nil
}

func NewExecutor(hostname, user, password string, port int) (*SSHExecutor, error) {
	if hostname == "" || user == "" {
		return nil, ErrHostUsername
	}
	exec := &SSHExecutor{Hostname: hostname, Port: port, User: user, Password: password}
	testErr := exec.connect()
	defer exec.disconnect()
	if testErr != nil {
		return nil, &ErrConnect{hostname: hostname, reason: testErr}
	}
	return exec, nil
}
