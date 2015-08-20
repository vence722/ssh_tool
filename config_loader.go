package main

import(
	"encoding/xml"
	"os"
	"io/ioutil"
)

type SSHConfig struct{
	Hosts struct{
		Host []struct{
			Hostname string
			User 	string
			Password string
		}
	}
	Commands struct{
		Command []struct{
			Hostname string
			Line []string
		}
	}
} 

func LoadConfig(cfgPath string) (*SSHConfig, error) {
	var cfg SSHConfig
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	data, errR := ioutil.ReadAll(f)
	if errR != nil{
		return nil, errR
	}
	errM := xml.Unmarshal(data, &cfg)
	if errM != nil{
		return nil, errM
	}
	return &cfg, nil
}