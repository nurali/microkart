package ctrl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurali/microkart/common/test"
	"github.com/nurali/microkart/config"
	"github.com/nurali/microkart/ctrl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LoginCtrlSuite struct {
	test.DBSuite
	ctrl ctrl.LoginCtrl
}

func TestLoginCtrl(t *testing.T) {
	config := config.New()
	suite.Run(t, &LoginCtrlSuite{
		DBSuite: test.NewDBSuite(config),
	})
}

func (s *LoginCtrlSuite) SetupSuite() {
	s.DBSuite.SetupSuite()
	s.ctrl = ctrl.NewLoginCtrl(s.DB)
}

func (s *LoginCtrlSuite) Test1Signup() {
	handler := http.HandlerFunc(s.ctrl.Signup)

	s.T().Run("ok", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/login/signup?username=john", nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Basic abcd1234")
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)

		assert.Equalf(t, http.StatusOK, res.Code, "unexpected status code")
	})

}

func (s *LoginCtrlSuite) Test2Login() {
	handler := http.HandlerFunc(s.ctrl.Login)

	s.T().Run("ok", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/login/login?username=john", nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Basic abcd1234")
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)

		assert.Equalf(t, http.StatusOK, res.Code, "unexpected status code")
	})
}
