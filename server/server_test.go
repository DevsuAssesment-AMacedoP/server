package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDevOpsHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      string
		headers    map[string]string
		want       string
		statusCode int
	}{
		{
			name:       "GET Method",
			method:     http.MethodGet,
			input:      ``,
			headers:    map[string]string{},
			want:       "ERROR",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "POST Method no Content-Type",
			method:     http.MethodPost,
			input:      ``,
			headers:    map[string]string{},
			want:       `{"message":"Content-Type header is not application/json"}`,
			statusCode: http.StatusUnsupportedMediaType,
		},
		{
			name:       "POST Method empty body",
			method:     http.MethodPost,
			input:      ``,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Request body must not be empty"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST Method JSON Syntax Error",
			method:     http.MethodPost,
			input:      `{ "message" : 45, "to": Juan Perez", "from": "Rita Asturia", "timeToLifeSec" : 45 }`,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Request body contains badly-formed JSON (at position 25)"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST Method bad JSON value type",
			method:     http.MethodPost,
			input:      `{ "message" : 45, "to": "Juan Perez", "from": "Rita Asturia", "timeToLifeSec" : 45 }`,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Request body contains an invalid value for the \"message\" field (at position 16)"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST Method JSON EOF",
			method:     http.MethodPost,
			input:      `{ "message" : "st`,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Request body contains badly-formed JSON"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST Method bad JSON unknown value",
			method:     http.MethodPost,
			input:      `{ "message" : "This is a test", "to": "Juan Perez", "from": "Rita Asturia", "timeToLifeSec" : 45, "unknown": "unknown" }`,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Request body contains unknown field \"unknown\""}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST Method ok",
			method:     http.MethodPost,
			input:      `{ "message" : "This is a test", "to": "Juan Perez", "from": "Rita Asturia", "timeToLifeSec" : 45 }`,
			headers:    map[string]string{"Content-Type": "application/json"},
			want:       `{"message":"Hello Juan Perez your message will be sent"}`,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bodyReader := strings.NewReader(tc.input)
			request := httptest.NewRequest(tc.method, "/DevOps", bodyReader)

			for key, val := range tc.headers {
				request.Header.Set(key, val)
			}

			responseRecorder := httptest.NewRecorder()

			HandleDevOps(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
