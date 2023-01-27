package server

import (
	"exec-processor/internal/utils"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ServerSettings struct {
	Settings struct {
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
	utils.CheckError(err, "can't open server configuration file")
	bytesRead, err := io.ReadAll(file)
	utils.CheckError(err, "can't read server configuration file")
	err = file.Close()
	utils.CheckError(err, "can't close configuration file")
	err = yaml.Unmarshal(bytesRead, serverSettings)
	utils.CheckError(err, "can't unmarshal configuration file")
}

func printDebugSettings(serverSettings *ServerSettings) {
	fmt.Println("IpAddress is: " + serverSettings.Settings.IpAddress)
	fmt.Println("Port is: " + serverSettings.Settings.Port)
	fmt.Println("ServerName is: " + serverSettings.Settings.ServerName)
	fmt.Println("Endpoint is: " + serverSettings.Settings.Endpoint)
}
