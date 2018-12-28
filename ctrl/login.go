package ctrl

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type LoginCtrl interface {
	BaseCtrl
	Login(rw http.ResponseWriter, r *http.Request)
	Logout(rw http.ResponseWriter, r *http.Request)
	Signup(rw http.ResponseWriter, r *http.Request)
}

type loginCtrl struct {
	baseCtrl
	users map[string]string
}

func NewLoginCtrl() LoginCtrl {
	c := new(loginCtrl)
	c.users = make(map[string]string)
	return c
}

func (c *loginCtrl) Name() string {
	return "login controller"
}

func (c *loginCtrl) Login(rw http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	passwd := c.users[username]
	if passwd == "" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	userPasswd, err := extractPasswd(r.Header.Get("Authorization"))
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userPasswd != passwd {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	return
}

func (c *loginCtrl) Logout(rw http.ResponseWriter, r *http.Request) {
	return
}

func (c *loginCtrl) Signup(rw http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if c.users[username] != "" {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("username already exists"))
		return
	}

	passwd, err := extractPasswd(r.Header.Get("Authorization"))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	c.users[username] = passwd

	return
}

func extractPasswd(authInfo string) (string, error) {
	authParts := strings.Split(authInfo, " ")
	log.Debugf("authParts:%v", len(authParts))
	if len(authParts) != 2 {
		return "", errors.New("Invalid authorization")
	}
	return authParts[len(authParts)-1], nil
}
