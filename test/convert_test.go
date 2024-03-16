package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cepot-blip/image-processing-service/handlers"
)

func TestConvertImageHandler(t *testing.T) {
	reqBody := []byte(`{"png_path": "/path/to/input.png", "jpeg_path": "/path/to/output.jpg"}`)
	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ConvertImageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Image converted successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
