package ctrl

import (
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type StatusCtrl interface {
	BaseCtrl
	Show(rw http.ResponseWriter, r *http.Request)
}

type statusCtrl struct {
	baseCtrl
	db *sql.DB
}

func NewStatusCtrl(db *sql.DB) StatusCtrl {
	c := &statusCtrl{db: db}
	return c
}

func (c *statusCtrl) Name() string {
	return "status controller"
}

func (c *statusCtrl) Show(rw http.ResponseWriter, r *http.Request) {
	err := c.db.Ping()
	if err != nil {
		log.Errorf("db not up, error:%v", err)
		rw.Write([]byte("Welcome to MicroKart !!!\nStatus: Error, Message: DB is not up"))
		return
	}
	rw.Write([]byte("Welcome to MicroKart !!!\nStatus: OK"))
}
