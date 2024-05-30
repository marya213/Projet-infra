package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"forum"
)

func TestRegisterHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/register", strings.NewReader("email=test@example.com&username=testuser&password=testpass"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(forum.RegisterHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}
