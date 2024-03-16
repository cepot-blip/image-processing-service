package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cepot-blip/image-processing-service/handlers"
)

func TestCompressImageHandler(t *testing.T) {
	reqBody := []byte(`{"input_path": "/path/to/input.jpg", "quality": 80}`)
	req, err := http.NewRequest("POST", "/compress", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CompressImageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
