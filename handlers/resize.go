package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cepot-blip/image-processing-service/utils"
)

type ResizeRequest struct {
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

func ResizeImageHandler(w http.ResponseWriter, r *http.Request) {
	var req ResizeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = utils.ResizeImage(req.InputPath, req.OutputPath, req.Width, req.Height)
	if err != nil {
		http.Error(w, "Failed to resize image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image resized successfully"))
}
