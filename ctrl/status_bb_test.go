package ctrl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/nurali/microkart/ctrl"
)

type StatusCtrlSuite struct {
	suite.Suite
	ctrl ctrl.StatusCtrl
}

func TestStatusCtrl(t *testing.T) {
	suite.Run(t, &StatusCtrlSuite{ctrl: ctrl.NewStatusCtrl()})
}

func (s *StatusCtrlSuite) TestStatusShow() {
	handler := http.HandlerFunc(s.ctrl.Show)

	s.T().Run("ok", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/status", nil)
		require.NoError(t, err)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)

		assert.Equalf(t, http.StatusOK, res.Code, "unexpected status code")
		expected := "Welcome to MicroKart, Status OK"
		assert.Equalf(t, expected, res.Body.String(), "unexpected response body")
	})
}
