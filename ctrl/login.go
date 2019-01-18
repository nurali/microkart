package ctrl

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/nurali/microkart/model"
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
	repo model.Respository
}

func NewLoginCtrl(db *sql.DB) LoginCtrl {
	c := &loginCtrl{
		repo: model.NewRepository(db),
	}
	return c
}

func (c *loginCtrl) Name() string {
	return "login controller"
}

func (c *loginCtrl) Login(rw http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		log.Errorf("username is missing")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.repo.Load(username)
	if err != nil {
		log.Errorf("load user failed, err:%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	passwd := user.Passwd
	if passwd == nil {
		log.Errorf("user not found")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	userPasswd, err := extractPasswd(r.Header.Get("Authorization"))
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userPasswd != *passwd {
		log.Errorf("wrong password")
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
		log.Errorf("username is missing")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.repo.Load(username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("load user failed, err:%v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if user != nil {
		log.Errorf("username '%s' already exists", username)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("username already exists"))
		return
	}

	passwd, err := extractPasswd(r.Header.Get("Authorization"))
	if err != nil {
		log.Errorf("authentication is missing, err:%v", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	newUser := &model.User{
		Name:   &username,
		Passwd: &passwd,
	}
	_, err = c.repo.Create(newUser)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("db insert operation failed"))
		return
	}

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
