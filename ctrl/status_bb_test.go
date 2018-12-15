package ctrl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurali/microkart/ctrl"
)

var statusCtrl = ctrl.NewStatusCtrl()

func TestStatusShow(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(statusCtrl.Show)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("call to show status failed with unexpected status code, got:%d, want:%d", status, http.StatusOK)
	}

	expected := "Welcome to MicroKart, Status OK"
	if res.Body.String() != expected {
		t.Errorf("call to show status failed with unexpected response, got:%s, want:%s", res.Body.String(), expected)
	}
}
