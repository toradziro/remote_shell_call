package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func sendInternalError(res http.ResponseWriter) {
	header := res.Header()
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, string("Internal error or wrong agrs, check it, please"))
}

func sendBadRequest(res http.ResponseWriter) {
	header := res.Header()
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(res, string("Bad request, not supported"))
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendInternalError(res)
	}

	reader := bufio.NewReader(req.Body)
	bytesRead, err := io.ReadAll(reader)
	checkError(err)
	var clientData ClientData
	err = readClientData(bytesRead, &clientData)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}

	var splitedCommand = strings.Split(clientData.Cmd, " ")
	var commandName = splitedCommand[0]

	cmd := exec.Command(commandName)
	cmd.Args = splitedCommand
	stdin, err := cmd.StdinPipe()
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if runtime.GOOS != clientData.OsName {
		sendInternalError(res)
		fmt.Print("Error: Wrong OS in client JSON")
		return
	}
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Start(); err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}
	_, err = io.WriteString(stdin, clientData.Stdin)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}

	header := res.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusOK)
	marshuledData, err := json.Marshal(ClientRespondData{string(stdout.Bytes()), string(stderr.Bytes())})
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}
	fmt.Fprintf(res, string(marshuledData))
}

func main() {
	var pathToSettings = "server_config.yaml"
	var serverSettings ServerSettings
	readSettingsData(pathToSettings, &serverSettings)
	printDebugSettings(&serverSettings)
	mux := http.NewServeMux()
	mux.HandleFunc(serverSettings.Settings[0].Endpoint, handleRequest)
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		sendBadRequest(res)
	})
	err := http.ListenAndServeTLS(serverSettings.Settings[0].IpAddress+":"+serverSettings.Settings[0].Port,
		serverSettings.Settings[0].ServerCrt,
		serverSettings.Settings[0].ServerKey,
		mux)
	if err != nil {
		log.Fatal(err)
	}
}
