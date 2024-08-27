package backend

import (
	"context"
	"fmt" 
	"io"
	"time"
	"cloud.google.com/go/storage"
	"net/http"
	"log"
	"around/constants"
)

/*
var (
	GCSBackend *GoogleCloudStorageBackend
)


type GoogleCloudStorageBackend struct {
	client *storage.Client
	bucket string
}
*/

type GCSBackendStruct struct {
    Client *storage.Client
    Bucket string
}
var GCSBackend *GCSBackendStruct

func InitGCSBackend() {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create GCS client: %v", err)
    }

    GCSBackend = &GCSBackendStruct{
        Client: client,
        Bucket: constants.GCS_BUCKET, // Replace with your actual bucket name
    }
}

func (backend *GCSBackendStruct) SaveToGCS(r io.Reader, objectName string)(string, error) {
	ctx := context.Background()
	object := backend.Client.Bucket(backend.Bucket).Object(objectName)
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}

	fmt.Printf("File is saved to GCS: %s\n", attrs.MediaLink)
	return attrs.MediaLink, nil
}

func (gcs *GCSBackendStruct) SaveImageURLToGCS(imageUrl string, imageId string) (string, error) {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Second*50)
    defer cancel()

    // Get the image data from the URL
    resp, err := http.Get(imageUrl)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bucket := gcs.Client.Bucket(gcs.Bucket)
    obj := bucket.Object(imageId)

    wc := obj.NewWriter(ctx)
    if _, err := io.Copy(wc, resp.Body); err != nil {
        return "", err
    }
    if err := wc.Close(); err != nil {
        return "", err
    }

    return fmt.Sprintf("https://storage.googleapis.com/%s/%s", gcs.Bucket, imageId), nil
}
