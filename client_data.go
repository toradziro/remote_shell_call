package main

import "encoding/json"

type ClientData struct {
	Cmd    string `json:"cmd"`
	OsName string `json:"os"`
	Stdin  string `json:"stdin"`
}

type ClientRespondData struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func readClientData(readBytes []byte, clientData *ClientData) error {
	err := json.Unmarshal(readBytes, clientData)
	return err
}
