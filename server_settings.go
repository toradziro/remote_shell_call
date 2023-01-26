package main

import (
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ServerSettings struct {
	Settings []struct {
		IpAddress  string `yaml:"ip_address"`
		Port       string `yaml:"port"`
		ServerName string `yaml:"server_name"`
		Endpoint   string `yaml:"endpoint"`
		ServerCrt  string `yaml:"server_crt"`
		ServerKey  string `yaml:"server_key"`
	} `yaml:"server_properties"`
}

func readSettingsData(pathToSettings string, serverSettings *ServerSettings) {
	file, err := os.Open(pathToSettings)
	checkError(err)
	bytesRead, err := io.ReadAll(file)
	checkError(err)
	err = file.Close()
	checkError(err)
	err = yaml.Unmarshal(bytesRead, serverSettings)
	checkError(err)
}

func printDebugSettings(serverSettings *ServerSettings) {
	fmt.Println("IpAddress is: " + serverSettings.Settings[0].IpAddress)
	fmt.Println("Port is: " + serverSettings.Settings[0].Port)
	fmt.Println("ServerName is: " + serverSettings.Settings[0].ServerName)
	fmt.Println("Endpoint is: " + serverSettings.Settings[0].Endpoint)
}
