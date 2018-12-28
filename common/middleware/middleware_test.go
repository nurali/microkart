package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurali/microkart/common/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MiddlewareSuite struct {
	suite.Suite
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, &MiddlewareSuite{})
}

func (s *MiddlewareSuite) TestMiddlewareChain() {
	wrapper := middleware.Chain(middleware.RequestID, middleware.Recover, middleware.Logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/simple", wrapper(simpleHandler))
	mux.HandleFunc("/panic", wrapper(panicHandler))
	server := httptest.NewServer(mux)

	s.T().Run("ok", func(t *testing.T) {
		url := "http://" + server.Listener.Addr().String() + "/simple"
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
	})

	s.T().Run("panic", func(t *testing.T) {
		url := "http://" + server.Listener.Addr().String() + "/panic"
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		assert.Equal(t, 500, res.StatusCode)
	})

}

func simpleHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Header.Get(middleware.RequestIDHeader) != "" {
		rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func panicHandler(rw http.ResponseWriter, r *http.Request) {
	panic("testing")
}
