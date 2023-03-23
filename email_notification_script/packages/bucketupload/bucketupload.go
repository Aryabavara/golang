package bucketupload

import (
	"context"
	"fmt"
	"io"
	"os"
	"encoding/json"
	"cloud.google.com/go/storage"
	models "count/models"
)

func UploadToGCS(bucketName, objectName string) error {
	ctx := context.Background()

	// Create a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open the local JSON file.
	file, err := os.Open("./downDeployment.json")
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()

	// Create a writer to the bucket.
	w := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	if err != nil {
		// Handle the error
	}
	if _, err := io.Copy(w, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

// func UpdateGCS(bucketName string, objectName string, newDeployments[]models.DeploymentStatus) error {

//     // // Loop through the slice and add each element to the map
//     // for i, s := range newDeployments {
//     //     m[s] = i
//     // }

//     // // Print the map
//     // fmt.Println(m)
// 	jsonStr := `[{"name":"nginx-deployment","namespace":"default","status":"","timestamp":"2023-03-07T11:03:14.60213017+05:30"},{"name":"python","namespace":"default","status":"","timestamp":"2023-03-07T11:03:14.602130649+05:30"}]`
// 	var deployments []map[string]string
// 	if err := json.Unmarshal([]byte(jsonStr), &deployments); err != nil {
// 		panic(err)
// 	}

// 	deploymentsMap := make(map[string]map[string]string)
// 	for _, deployment := range deployments {
// 		key := deployment["name"]
// 		deploymentsMap[key] = deployment
// 	}

// 	fmt.Println(deploymentsMap)

// 	ctx := context.Background()

// 	// Create a client.
// 	client, err := storage.NewClient(ctx)
// 	w := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

// 	// Write the new data to the object
// 	_, err = io.Copy(w, deploymentsMap)
// 	if err != nil {
// 		// Handle the error
// 	}

// 	// Close the Writer to finalize the object update
// 	err = w.Close()
// 	if err != nil {
// 		// Handle the error
// 	}
// }
func UpdateGCS(bucketName string, objectName string) error {
	ctx := context.Background()

	// Create a client.
	client, err := storage.NewClient(ctx)
		// Create a new Writer to update the existing object in the bucket
	w := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open the local JSON file.
	file, err := os.Open("./downDeployment.json")
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()
	// Write the new data to the object
	_, err = io.Copy(w, file)
	if err != nil {
		// Handle the error
	}

	// Close the Writer to finalize the object update
	err = w.Close()
	if err != nil {
		// Handle the error
 	}
	fmt.Println("++++++++++++++++++++++++++++++++++")
	return nil
}

//upload new slice into bucket

func UploadNewSlice(bucketName string, fileName string, data []models.DeploymentStatus) error {
	ctx := context.Background()

	// Create a Google Cloud Storage client
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}

	// Open a writeable file in the bucket
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)

	// Encode the data as JSON and write it to the file
	enc := json.NewEncoder(wc)
	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	// Close the file
	err = wc.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %v", err)
	}

	return nil
}