package api_test

import (
	"net/http"
	"reward-system-api/internal/api"
	"testing"
)

func TestApplication_RegisterHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var app api.Application
			app.RegisterHandler(tt.w, tt.r)
		})
	}
}
