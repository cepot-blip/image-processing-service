package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cepot-blip/image-processing-service/handlers"
)

func TestResizeImageHandler(t *testing.T) {
	reqBody := []byte(`{"input_path": "/path/to/input.jpg", "output_path": "/path/to/output.jpg", "width": 100, "height": 100}`)
	req, err := http.NewRequest("POST", "/resize", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ResizeImageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Image resized successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
