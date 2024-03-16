package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cepot-blip/image-processing-service/utils"
)

type CompressRequest struct {
	InputPath string `json:"input_path"`
	Quality   int    `json:"quality"`
}

func CompressImageHandler(w http.ResponseWriter, r *http.Request) {
	var req CompressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	compressedData, err := utils.CompressImage(req.InputPath, req.Quality)
	if err != nil {
		http.Error(w, "Failed to compress image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(compressedData)
}
