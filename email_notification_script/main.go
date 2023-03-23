package main

import (
	"context"
	models "count/models"
	gcpUpload "count/packages/bucketupload"
	count "count/packages/deployment"

	// email "count/packages/email"
	fileDiffrence "count/packages/filediffrence"
	readFromBucket "count/packages/read_from_bucket"
	readFromLocal "count/packages/read_from_local"

	// email "count/packages/email"
	"fmt"

	// "time"
	"encoding/json"
	"os"

	"cloud.google.com/go/storage"

	"log"
	"reflect"
)

func main() {
	var deploy_Name_in_bucket string
	bucketName := "prd-asp-deployment-fail-test"
	// The name of the file you want to create
	fileName := "downDeployment.json"

	deploymentStatuses := count.Deployments()

	file, err := os.Create("downDeployment.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(deploymentStatuses)
	if err != nil {
		panic(err)
	}

	///////////// Checking whether a json file is existing inside the gcp bucket ////////////////
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	_, err = client.Bucket(bucketName).Object(fileName).Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			//////////////// If the file is not existing uploading the json file from the loacal which is containing failed deployment details.
			gcpUpload.UploadToGCS(bucketName, fileName)

		} else {
			log.Fatalf("Failed to check if file exists: %v", err)
		}
	} else {
		_, _, bucket_file_content, _, err := readFromBucket.Read_from_bucket()
		_, _, local_fileContent, _ := readFromLocal.Read_json_from_local()
		fmt.Println("_-__-_-__-__-__-__-", deploy_Name_in_bucket)
		if err != nil {
			log.Fatalf("Error reading from bucket: %v", err)
		}

		var deployments1 []models.DeploymentStatus
		var deployments2 []models.DeploymentStatus

		fmt.Println(reflect.TypeOf(deployments2))
		

		json.Unmarshal(bucket_file_content, &deployments1)
		json.Unmarshal(local_fileContent, &deployments2)

		result1, result2 := fileDiffrence.File_difference(deployments1, deployments2)

		var result string

		if len(result2) != 0 && len(result1) == 0 {
			newDeployments := fileDiffrence.Fileupdation(deployments1, deployments2)
			newDeployments = append(newDeployments, result2...)
			gcpUpload.UploadNewSlice(bucketName,fileName,newDeployments)

		} else if len(result1) != 0 && len(result2) == 0 { 
			newDeployments := fileDiffrence.Fileupdation(deployments1, deployments2)
			for _, deployment := range result1 {
				fmt.Printf("Name: %s\nName: %s\nStatus: %s\n\n", deployment.Name,deployment.Namespace, deployment.Status)
				result = result + ""+"Namesapace "+":"+" "+deployment.Namespace+ " " + " \n"+"Deployment Name " +":"+ deployment.Name + " " + " \n" +"Time :"+ deployment.Timestamp.Format("15:04:05 PM") + "\n"
			}
			gcpUpload.UploadNewSlice(bucketName,fileName,newDeployments)
		} else if len(result1) == 0 && len(result1) == 0 {

			fmt.Println("33333333333333333")
			fmt.Sprintln("no deployment failed")
	}else{
		newDeployments := fileDiffrence.Fileupdation(deployments1, deployments2)
		newDeployments = append(newDeployments, result2...)
	}}
	_, _, _, result := readFromLocal.Read_json_from_local()
	fmt.Println("___________________________________________________", result)
}