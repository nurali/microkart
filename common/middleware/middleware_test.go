package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurali/microkart/common/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddlewareChain(t *testing.T) {
	wrapper := middleware.Chain(middleware.RequestID, middleware.Logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/", wrapper(myHandler))
	server := httptest.NewServer(mux)

	url := "http://" + server.Listener.Addr().String()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	log.Infof("In myHandler")
	defer log.Infof("Out myHandler")
	rw.Write([]byte("OK"))
}
