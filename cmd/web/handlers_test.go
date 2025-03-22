package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogVBreedsJSON(t *testing.T) {
	// create request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	// response recorder
	rr := httptest.NewRecorder()

	// handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	// serve hamdler
	handler.ServeHTTP(rr, req)

	// check response\
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code: got %d, wanted 200", rr.Code)
	}
}