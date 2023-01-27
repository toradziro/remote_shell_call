package handler

import (
	"bufio"
	"encoding/json"
	"exec-processor/internal/entity"
	"exec-processor/internal/workers"
	"fmt"
	"io"
	"net/http"
	"time"
)

func sendInternalError(res http.ResponseWriter) {
	header := res.Header()
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, string("Internal error or wrong agrs, check it, please"))
}

func SendBadRequest(res http.ResponseWriter) {
	header := res.Header()
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(res, string("Bad request, not supported"))
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendInternalError(res)
	}

	reader := bufio.NewReader(req.Body)
	bytesRead, err := io.ReadAll(reader)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}
	var clientData entity.ClientData
	err = entity.ReadClientData(bytesRead, &clientData)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}

	header := res.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Date", time.Now().String())
	res.WriteHeader(http.StatusOK)
	err, outputData := workers.ProcessExecution(clientData)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}

	marshalledData, err := json.Marshal(outputData)
	if err != nil {
		sendInternalError(res)
		fmt.Println(err)
		return
	}
	fmt.Fprintf(res, string(marshalledData))
}
