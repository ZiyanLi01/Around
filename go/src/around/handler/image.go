package handler

import (
    "encoding/json"
    "net/http"
    "around/service"
    "around/model"
    "github.com/gorilla/mux"
)

// GenerateImageHandler handles the request to generate an AI image
func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Description string `json:"description"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

    userId := r.Context().Value("userID").(string)

    imageUrl, err := service.GenerateAIImage(req.Description, userId)
    if err != nil {
        http.Error(w, "Failed to generate image", http.StatusInternalServerError)
        return
    }

    image := model.Image{
        Id:          imageUrl, // You can generate an ID if needed
        Url:         imageUrl,
        Description: req.Description,
        UserId:      userId,
    }

    json.NewEncoder(w).Encode(image)
}


// DownloadImageHandler serves the saved image for download
func DownloadImageHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    imageUrl := vars["imageUrl"]

    imageData, err := service.DownloadImage(imageUrl)
    if err != nil {
        http.Error(w, "Failed to download image", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "image/png")
    w.Write(imageData)
}