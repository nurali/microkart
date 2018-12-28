package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/nurali/microkart/common/middleware"
	"github.com/nurali/microkart/config"
	"github.com/nurali/microkart/ctrl"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.New()

	initLogger(config.GetLogLevel())

	db := setupDb(config.GetPostgresConfigString())
	defer db.Close()

	mux := mountEndpoints(db)
	startServer(config.GetHttpPort(), mux)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func setupDb(dbConfig string) *sql.DB {
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Panicf("db open failed, err:%v", err)
	}
	return db
}

func mountEndpoints(db *sql.DB) *http.ServeMux {
	wrapper := middleware.Chain(middleware.RequestID, middleware.Recover, middleware.Logger)
	mux := http.NewServeMux()

	statusCtrl := ctrl.NewStatusCtrl(db)
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
