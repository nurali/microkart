package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nurali/microkart/common/middleware"
	"github.com/nurali/microkart/config"
	"github.com/nurali/microkart/ctrl"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.New()

	initLogger(config.GetLogLevel())

	mux := mountEndpoints()
	startServer(config.GetHttpPort(), mux)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func mountEndpoints() *http.ServeMux {
	wrapper := middleware.Chain(middleware.RequestID, middleware.Recover, middleware.Logger)
	mux := http.NewServeMux()

	statusCtrl := ctrl.NewStatusCtrl()
	log.Infof("mounted:%s", statusCtrl.Name())
	mux.HandleFunc("/api/status", wrapper(statusCtrl.Show))

	loginCtrl := ctrl.NewLoginCtrl()
	log.Infof("mounted:%s", loginCtrl.Name())
	mux.HandleFunc("/api/login/login", wrapper(loginCtrl.Login))
	mux.HandleFunc("/api/login/logout", wrapper(loginCtrl.Logout))
	mux.HandleFunc("/api/login/signup", wrapper(loginCtrl.Signup))

	return mux
}

func startServer(port int, mux *http.ServeMux) {
	addr := fmt.Sprintf("localhost:%d", port)
	log.Infof("microkart running at:%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start microkart, error:%v", err)
	}
}
