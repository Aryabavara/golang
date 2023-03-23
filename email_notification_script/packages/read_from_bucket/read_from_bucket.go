package readfrombucket

import (
	"context"
	models "count/models"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"cloud.google.com/go/storage"
)

func Read_from_bucket() (deploy_Name_in_bucket string, deploy_Time_in_bucket string, bucket_file_content []byte,  result string, err error) {

	var previous_down_deployments []models.DeploymentStatus

	bucketName := "prd-asp-deployment-fail-test"
	// The name of the file you want to create
	fileName := "downDeployment.json"

	// if file exist inside gcp bucket, read the content into a struct and compare it with the local file which is created.
	ctx := context.Background()

	// Create a client for interacting with GCS.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	rc, err := client.Bucket(bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		log.Printf("Error opening reader: %v\n", err)
		return
		// TODO: handle error.
	}
	defer rc.Close()
	bucket_file_content, err = ioutil.ReadAll(rc)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", reflect.TypeOf(bucket_file_content))
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(reflect.TypeOf(bucket_file_content))
		return
	}
	// Unmarshal the JSON contents into  struct DeploymentStatus .
	if err := json.Unmarshal(bucket_file_content, &previous_down_deployments); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// var deploy_Name_in_bucket string
	// var deploy_Time_in_bucket string
	// Print the contents of the struct.

	for _, deployment := range previous_down_deployments {
		fmt.Printf("Name: %s\nName: %s\nStatus: %s\n\n", deployment.Name,deployment.Namespace, deployment.Status)
		result = result + " " + " " + deployment.Name + " " +deployment.Namespace+ " " +deployment.Status + " "+ deployment.Status + " " + deployment.Timestamp.Format("15:04:05 PM") + "\n"
		deploy_Name_in_bucket = deployment.Name
		deploy_Time_in_bucket = deployment.Timestamp.Format("15:04:05 PM")
	}

	rc.Close()
	fmt.Printf("File %s exists in bucket %s\n", fileName, bucketName)
	return deploy_Name_in_bucket, deploy_Time_in_bucket, bucket_file_content, result, nil
}
