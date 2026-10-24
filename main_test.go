package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmptyMessage(t *testing.T) {
	//pretend incoming request
	req := httptest.NewRequest(http.MethodPost, "/embed", strings.NewReader(`{"text":""}`))

	//saves the outgoing reply

	w := httptest.NewRecorder()

	//run handler with both
	handleEmbed(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
