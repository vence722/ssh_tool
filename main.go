package main

import (
	"fmt"
	"os"
)

var hostname,user,password string

func main() {
	cfg, errCfg := LoadConfig("./config")
	if errCfg != nil{
		fmt.Println(errCfg.Error())
		os.Exit(1)
	}

	var executors map[string]*SSHExecutor = make(map[string]*SSHExecutor)
	for _, host := range cfg.Hosts.Host{
		exec, errHost := NewExecutor(host.Hostname, host.User, host.Password, PORT_SSH_DEFAULT)
		if errHost != nil{
			fmt.Println(errHost.Error())
			os.Exit(1)
		}
		executors[host.Hostname] = exec
	}
	
	for _, command := range cfg.Commands.Command{
		execs, errExec := filterCommandHostname(command.Hostname, executors)
		if errExec != nil {
			fmt.Println(errExec)
			os.Exit(1)
		}
		for _, exec := range execs {
			runCommand(exec, command.Line)
		}	
	}	
}

func runCommand(exec *SSHExecutor, commandLines []string){
	exec.SetCommands(commandLines)
	result := exec.Exec()
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.Message != "" {
		fmt.Println(result.Message)
	}
}


