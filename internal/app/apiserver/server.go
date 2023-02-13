package apiserver

import (
	"log"
	"net/http"
	"ordersBook/pkg"
)

type APIServer struct {
	config *Config
	logger *log.Logger
	mux    *http.ServeMux
}

// New | Default config
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		mux:    http.NewServeMux(),
		logger: log.Default(),
	}
}

func (s *APIServer) Start() error {
	s.logInfo(textAboutServerInfo)
	s.configureRoutes()
	s.logInfo(textStartServer, s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.mux)
}

func (s *APIServer) logInfo(v ...any) {
	s.logger.Printf("[Info]: %s\n", pkg.StringBuild(v...))
}

func (s *APIServer) logWarn(v ...any) {
	s.logger.Printf("[Warn]: %v\n", pkg.StringBuild(v...))
}

func (s *APIServer) logError(err error, v ...any) {
	s.logger.Printf("[Error]: %v %v\n", err, v)
}
