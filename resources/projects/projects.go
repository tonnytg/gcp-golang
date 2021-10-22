package projects

import (
	"encoding/json"
	"fmt"
	"gcp-golang/pkg/web"
	"time"
)

type AllProjects struct {
	Projects []struct {
		ProjectNumber  string    `json:"projectNumber"`
		ProjectID      string    `json:"projectId"`
		LifecycleState string    `json:"lifecycleState"`
		Name           string    `json:"name"`
		CreateTime     time.Time `json:"createTime"`
	} `json:"projects"`
}

func Get() ([]byte, error) {

	// GET ServiceAccount of one project
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"

	// get []byte of request
	data, err := web.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))

	//var pjs Projects
	var allProjects AllProjects
	json.Unmarshal(data, &allProjects)
	for i, v := range allProjects.Projects {
		fmt.Printf("Project[%d]: %s\n", i, v.Name)
	}
	return data, nil
}
