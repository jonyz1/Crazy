package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	// Create a new HTTP request for signup
	signupData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	signupPayload, _ := json.Marshal(signupData)
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(signupPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SignUpHandler) // Your handler function here
	handler.ServeHTTP(rr, req) // Serve the request using the handler

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Account created successfully! Please log in.", response["message"])
}

func TestLogin(t *testing.T) {
	// Create a new HTTP request for login
	loginData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	loginPayload, _ := json.Marshal(loginData)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler) // Your handler function here
	handler.ServeHTTP(rr, req) // Serve the request using the handler

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the token is returned
	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NotNil(t, response["token"])
}

func TestRoomSelection(t *testing.T) {
	// Create a new HTTP request for room selection with valid token
	req, err := http.NewRequest("GET", "/roomselection?suit=hearts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Authorization header with a valid token
	req.Header.Set("Authorization", "Bearer valid-token")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RoomSelectionHandler) // Your handler function here
	handler.ServeHTTP(rr, req) // Serve the request using the handler

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response message
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "good", response["message"])
}

func TestRoomSelectionUnauthorized(t *testing.T) {
	// Create a new HTTP request for room selection with invalid token
	req, err := http.NewRequest("GET", "/roomselection?suit=hearts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Authorization header with an invalid token
	req.Header.Set("Authorization", "Bearer invalid-token")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RoomSelectionHandler) // Your handler function here
	handler.ServeHTTP(rr, req) // Serve the request using the handler

	// Assert the response status code
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// Assert the error message
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Authentication failed", response["error"])
}
