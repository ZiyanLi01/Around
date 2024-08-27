package handler

import (
    "encoding/json"
    "net/http"
    "around/service"
    "around/model"
    "github.com/gorilla/mux"
    "log"
    "github.com/google/uuid"
)

func generateUniqueId() string {
    return uuid.New().String()
}

// GenerateImageHandler handles the request to generate an AI image
func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Description string `json:"description"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Printf("Failed to parse request body: %v", err)
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

    userID, ok := r.Context().Value(userIDKey).(string)
    if !ok || userID == "" {
        log.Println("UserID not found in context")
        http.Error(w, "UserID not found", http.StatusUnauthorized)
        return
    }

    imageUrl, err := service.GenerateAIImage(req.Description, userID)
    if err != nil {
        log.Printf("Failed to generate image: %v", err)
        http.Error(w, "Failed to generate image", http.StatusInternalServerError)
        return
    }

    image := model.Image{
        Id:          generateUniqueId(),
        Url:         imageUrl,
        Description: req.Description,
        UserId:      userID,  // Consistent naming here
    }

    if err := json.NewEncoder(w).Encode(image); err != nil {
        log.Printf("Failed to encode response: %v", err)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
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