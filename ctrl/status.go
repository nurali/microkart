package ctrl

import "net/http"

type StatusCtrl interface {
	BaseCtrl
	Show(rw http.ResponseWriter, r *http.Request)
}

type statusCtrl struct {
	baseCtrl
}

func NewStatusCtrl() StatusCtrl {
	c := new(statusCtrl)
	return c
}

func (c *statusCtrl) Name() string {
	return "status controller"
}

func (c *statusCtrl) Show(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcome to MicroKart, Status OK"))
}
