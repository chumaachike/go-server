package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)

	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Homepage",
			path:         "/",
			expectedCode: http.StatusOK,
			expectedBody: "Welcome to he Homepage!\n",
		},
		{
			name:         "Goodbye",
			path:         "/goodbye",
			expectedCode: http.StatusOK,
			expectedBody: "Goodbye!\n",
		},
		{
			name:         "Unknown route",
			path:         "/unknown",
			expectedCode: http.StatusNotFound,
			expectedBody: "404 page not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf(
					"expected status %d, got %d",
					tt.expectedCode,
					w.Code,
				)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf(
					"expected body %q, got %q",
					tt.expectedBody,
					w.Body.String(),
				)
			}
		})
	}
}

func TestHandleHelloParametized(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid user",
			url:            "/hello?user=TestMan",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, TestMan!\n",
		},
		{
			name:           "Another user",
			url:            "/hello?user=Chuma",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Chuma!\n",
		},
		{
			name:           "Missing user",
			url:            "/hello",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing user parameter\n",
		},
		{
			name:           "Empty user",
			url:            "/hello?user=",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing user parameter\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			handleHelloParameterized(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf(
					"expected status %d got %d",
					tt.expectedStatus,
					w.Code,
				)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf(
					"expected body %q, got %q",
					tt.expectedBody,
					w.Body.String(),
				)
			}
		})
	}

}

func TestHandleHelloHeader(t *testing.T) {
	tests := []struct {
		name           string
		headValue      string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid username header",
			headValue:      "chuma",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, chuma!\n",
		},
		{
			name:           "Missing username header",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "No header value given\n",
		},
		{
			name:           "Username with spaces",
			headValue:      "Chuma Achike",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Chuma Achike!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/", nil)
			if tt.headValue != "" {
				req.Header.Set("User", tt.headValue)
			}

			w := httptest.NewRecorder()

			handleHelloHeader(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf(
					"expected status %d got %d",
					tt.expectedStatus,
					w.Code,
				)
			}
			if w.Body.String() != tt.expectedBody {
				t.Errorf(
					"expected body %q, got %q",
					tt.expectedBody,
					w.Body.String(),
				)
			}
		})
	}

}
