package service

import (
    "fmt"
    "around/backend"

    "around/constants"


    "bytes"
    "encoding/json"
    
    "net/http"
    
    "time"
	"io/ioutil"
)

// GenerateAIImage calls the OpenAI API to generate an image based on the given description
func GenerateAIImage(description string, userId string) (string, error) {
    // Call OpenAI API to generate image
    // (This is a simplified version - make sure to use your actual API calling logic)
    requestBody, _ := json.Marshal(map[string]string{
        "description": description,
    })

    req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generate", ioutil.NopCloser(bytes.NewReader(requestBody)))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.OpenAIAPIKey))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var responseData map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&responseData)

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to generate image: %v", responseData["error"])
    }

    // Assume the API returns a URL for the generated image
    imageUrl := responseData["data"].(string)

    // Store in GCS
    imageId := fmt.Sprintf("%d", time.Now().UnixNano())
    gcsUrl, err := backend.GCSBackend.SaveImageURLToGCS(imageUrl, imageId)
    if err != nil {
        return "", err
    }

    // Return the GCS URL
    return gcsUrl, nil
}

func DownloadImage(imageUrl string) ([]byte, error) {
    client := &http.Client{}
    resp, err := client.Get(imageUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}