package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
)

func TestHomeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(homeHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "Hello! This is a demo app using Gorilla Mux v1.7.3"
    if !contains(rr.Body.String(), expected) {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestHealthHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "OK - Server is healthy"
    if !contains(rr.Body.String(), expected) {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestRouter(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/health", healthHandler).Methods("GET")

    testCases := []struct {
        method   string
        path     string
        expected int
    }{
        {"GET", "/", 200},
        {"GET", "/health", 200},
        {"GET", "/nonexistent", 404},
        {"POST", "/", 405},
    }

    for _, tc := range testCases {
        req, _ := http.NewRequest(tc.method, tc.path, nil)
        rr := httptest.NewRecorder()
        r.ServeHTTP(rr, req)

        if rr.Code != tc.expected {
            t.Errorf("%s %s: expected %d, got %d", tc.method, tc.path, tc.expected, rr.Code)
        }
    }
}

func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[:len(substr)] == substr || 
           (len(s) > len(substr) && contains(s[1:], substr))
}