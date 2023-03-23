package filedifference
import (
        "fmt"
		"time"
        models "count/models"
        "encoding/json"
        "io/ioutil"
        "os"
        // "encoding/json"
    )


func File_difference(s1 []models.DeploymentStatus, s2 []models.DeploymentStatus) ([]models.DeploymentStatus, []models.DeploymentStatus) {
    onlyInS1 := []models.DeploymentStatus{}
    onlyInS2 := []models.DeploymentStatus{}
    m := make(map[string]time.Time)

    for _, item := range s2 {
        m[item.Name] = item.Timestamp
    }

    for _, item := range s1 {
        if _, ok := m[item.Name]; ok {
            delete(m, item.Name)
        } else {
            onlyInS1 = append(onlyInS1, item)
        }
    }

    for name, Timestamp := range m {
        onlyInS2 = append(onlyInS2, models.DeploymentStatus{Name: name, Timestamp: Timestamp})
    }
	fmt.Println("deployments1_slice:",onlyInS1)
	fmt.Println("deployments2_slice:",onlyInS2)
    return onlyInS1, onlyInS2
}

func Fileupdation(jsonData[]models.DeploymentStatus, result[]models.DeploymentStatus) []models.DeploymentStatus{
	// The JSON data as a string
	// Unmarshal the JSON data into a slice of Deployment structs
    //not in local case ......remove that deploymnt details that was the code 
	var deployments []models.DeploymentStatus = jsonData

	// Loop through the deployments and remove the one with name "python"
	var newDeployments []models.DeploymentStatus

	for _, d := range deployments {
        for _, r := range result{
		if d.Name == r.Name && d.Namespace == r.Namespace {
			newDeployments = append(newDeployments, d)
		}
	}}


	// Print the updated JSON data as a string
	fmt.Println("**************************",(newDeployments))
    return newDeployments
}

func FileModify() {
    // Read the existing JSON file into a data structure
    //new content update into file
    file, err := os.Open("downDeployment.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    var people []models.DeploymentStatus
    err = json.NewDecoder(file).Decode(&people)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Modify the data structure as needed
    newPerson := models.DeploymentStatus{Name:"nginx-deployment",Namespace:"default",Status:"",Timestamp:time.Now()}
    people = append(people, newPerson)

    // Write the modified data structure back to the JSON file
    file, err = os.Create("downDeployment.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    data, err := json.MarshalIndent(people, "", "  ")
    if err != nil {
        fmt.Println(err)
        return
    }

    err = ioutil.WriteFile("downDeployment.json", data, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
}
