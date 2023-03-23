package readfromlocal

import (
    models "count/models"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func Read_json_from_local() (present_deploymet_name string, present_deployment_time string, local_fileContent []byte, result string) {
    // var result = ""
    local_file, err := os.Open("downDeployment.json")
    if err != nil {
        fmt.Println(err,)
        return
    } // Read the contents of the file.
    local_fileContent, err = ioutil.ReadAll(local_file)
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }
    defer local_file.Close()
    // Unmarshal the file into the struct.
    var present_deployments []models.DeploymentStatus
    if err := json.Unmarshal(local_fileContent, &present_deployments); err != nil {
        log.Fatalf("Failed to unmarshal JSON: %v", err)
    }

    // Print the contents of the struct.
    for _, Newdeployment := range present_deployments {
        result = result + ""+"Namesapace "+":"+" "+Newdeployment.Namespace+ " " + " \n"+"Deployment Name " +":"+ Newdeployment.Name + " " + " \n" +"Time :"+ Newdeployment.Timestamp.Format("15:04:05 PM") + "\n"
        present_deploymet_name = Newdeployment.Name
        present_deployment_time = Newdeployment.Timestamp.Format("15:04:05 PM")
        // fmt.Println("------------------------>", present_deployment_time)
        // fmt.Printf("Name: %s\nNamespace: %s\nStatus: %s\nTimestamp: %s\n\n", Newdeployment.Name, Newdeployment.Namespace, Newdeployment.Status, Newdeployment.Timestamp)

    }

    fmt.Println("local depl details", present_deploymet_name, "time:", present_deployment_time)
    return present_deployment_time, present_deploymet_name, local_fileContent, result

}