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

// wrong method test
func TestWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/embed", nil)
	w := httptest.NewRecorder()

	handleEmbed(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}
}

// valid message test
func TestValidMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/embed", strings.NewReader(`{"text": "hello"}`))
	w := httptest.NewRecorder()

	handleEmbed(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Body.String() == "" {
		t.Error("expected a body response but got none")
	}
}
