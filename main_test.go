package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// empty message test
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

// broken message test
func TestBrokenMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/embed", strings.NewReader(`{"text":""`))
	w := httptest.NewRecorder()

	handleEmbed(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

// invalid message test
func TestValidMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/embed", nil)
	w := httptest.NewRecorder()

	handleEmbed(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
