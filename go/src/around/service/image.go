package service

import (
    "fmt"
    "around/backend"

    "around/constants"


    "bytes"
    "encoding/json"
    
    "net/http"
    
   "errors"
	"io"
    "github.com/google/uuid"
)

func generateUniqueId() string {
    return uuid.New().String()
}

// GenerateAIImage calls the OpenAI API to generate an image based on the given description
func GenerateAIImage(description, userID string) (string, error) {
    imageId := generateUniqueId()

    requestBody, err := json.Marshal(map[string]string{
        "prompt": description,
    })
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", constants.OPENAI_API_URL, io.NopCloser(bytes.NewBuffer(requestBody)))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+constants.OpenAIAPIKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    fmt.Printf("OpenAI API response status: %s\n", resp.Status)

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    fmt.Printf("OpenAI API response body: %s\n", string(body))

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("OpenAI API request failed with status: " + resp.Status)
    }

    var jsonResponse map[string]interface{}
    if err := json.Unmarshal(body, &jsonResponse); err != nil {
        return "", err
    }

    // Extract the image URL from the "data" array
    data, ok := jsonResponse["data"].([]interface{})
    if !ok || len(data) == 0 {
        return "", errors.New("no image data found in OpenAI API response")
    }

    firstItem, ok := data[0].(map[string]interface{})
    if !ok {
        return "", errors.New("invalid image data format in OpenAI API response")
    }

    imageUrl, ok := firstItem["url"].(string)
    if !ok {
        return "", errors.New("image URL not found in OpenAI API response")
    }

    // Save image URL to GCS
    gcsUrl, err := backend.GCSBackend.SaveImageURLToGCS(imageUrl, imageId)
    if err != nil {
        return "", err
    }
    fmt.Printf("gcsUrl: %s\n", gcsUrl)
    fmt.Printf("imageUrl: %s\n", imageUrl)
    return imageUrl, nil

    
    
}




func DownloadImage(imageUrl string) ([]byte, error) {
    client := &http.Client{}
    resp, err := client.Get(imageUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}