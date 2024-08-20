package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestStatusInternalServerError(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "InternalServerErrorTest"},
		{name : "not internal"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ResponseRecorder, which satisfies the http.ResponseWriter interface
			rr := httptest.NewRecorder()

			// Call the function we're testing
			StatusInternalServerError(rr)

			// Check the status code is 500
			if status := rr.Code; status != http.StatusInternalServerError {
				t.Errorf("StatusInternalServerError() = %v, want %v", status, http.StatusInternalServerError)
			}

			// Optionally, you can also check the response body
		
		})
	}
}

