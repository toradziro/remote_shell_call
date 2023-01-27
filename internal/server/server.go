package server

import (
	"exec-processor/internal/handler"
	"fmt"
	"net/http"
)

type Server struct {
	Settings ServerSettings
}

func (server *Server) getURLConnString() string {
	return fmt.Sprintf("%s:%s", server.Settings.Settings.IpAddress, server.Settings.Settings.Port)
}

func (server *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(server.Settings.Settings.Endpoint, handler.HandleRequest)
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		handler.SendBadRequest(res)
	})
	err := http.ListenAndServeTLS(server.getURLConnString(),
		server.Settings.Settings.ServerCrt,
		server.Settings.Settings.ServerKey,
		mux)
	return err
}

func New(yamlFileName string) *Server {
	var serverSettings ServerSettings
	readSettingsData(yamlFileName, &serverSettings)
	printDebugSettings(&serverSettings)
	return &Server{Settings: serverSettings}
}
