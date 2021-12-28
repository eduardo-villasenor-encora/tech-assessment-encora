package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserPostHandler(t *testing.T) {
	scenarios := []struct {
		methodCall     string
		target         string
		expectedStatus int
	}{
		{
			methodCall:     http.MethodGet,
			target:         "/v1/user-posts/1",
			expectedStatus: 200,
		},
		{
			methodCall:     http.MethodPost,
			target:         "/v1/user-posts/1",
			expectedStatus: 400,
		},
	}
	for _, scenario := range scenarios {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(scenario.methodCall, scenario.target, nil)
		GetUserPostsHandler(response, request)
		if response.Code != scenario.expectedStatus {
			t.Fatalf("expected %d , given %d for scenario: %v", scenario.expectedStatus, response.Code, scenario)
		}
	}
}
