package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cepot-blip/image-processing-service/utils"
)

type ConvertRequest struct {
	PNGPath  string `json:"png_path"`
	JPEGPath string `json:"jpeg_path"`
}

func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	var req ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = utils.ConvertPNGToJPEG(req.PNGPath, req.JPEGPath)
	if err != nil {
		http.Error(w, "Failed to convert image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image converted successfully"))
}
