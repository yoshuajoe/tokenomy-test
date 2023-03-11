package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDataHandler_NoIdParam(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}],"message":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetDataHandler_WithId(t *testing.T) {
	// Test case 2: Request with single id
	req, err := http.NewRequest("GET", "/?id=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := `{"code":200,"data":[{"id":2,"name":"B"}],"message":"OK"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}

func TestGetDataHandler_MultipleIds(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=1,3,4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := `{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}],"message":"OK"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}

func TestGetDataHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=xxx", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status code %v but got %v", http.StatusBadRequest, rr.Code)
	}
	expected := `{"code":400,"message":"invalid or empty ID: \"xxx\""}`
	if rr.Body.String() != expected {
		t.Errorf("expected response body '%v' but got '%v'", expected, rr.Body.String())
	}
}

func TestGetDataHandler_IDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=4", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status code %v but got %v", http.StatusNotFound, rr.Code)
	}
	expected := `{"code":404,"message":"resource with ID 4 not exist"}`
	if rr.Body.String() != expected {
		t.Errorf("expected response body '%v' but got '%v'", expected, rr.Body.String())
	}
}
