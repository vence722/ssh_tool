package main

import(
	"strings"
	"errors"
)

var(
	ErrEmptyExec error = errors.New("error:Executors map can't be empty.")
)

type ErrHostNotExist struct{
	hostname string
}

func (this *ErrHostNotExist) Error() string {
	return "error:Host definition " + this.hostname + " not found, please make sure the hostname is defined in the hosts definition section."
}

func filterCommandHostname(hostname string, executors map[string]*SSHExecutor) (execs []*SSHExecutor, er error) {
	if executors == nil {
		return nil, ErrEmptyExec
	}
	if hostname == "*" {
		for _, exec := range executors{
			execs = append(execs, exec)
		}
		return execs,nil
	}
	if strings.Contains(hostname, ",") {
		hostnames := strings.Split(hostname, ",")
		for _, hostname := range hostnames {
			exec := executors[hostname]
			if exec == nil{
				return nil, &ErrHostNotExist{hostname:hostname}
			}
			execs = append(execs, exec)
		}
		return execs, nil
	}
	exec := executors[hostname]
	if exec == nil{
		return nil, &ErrHostNotExist{hostname:hostname}
	}
	execs = append(execs, exec)
	return execs, nil
}