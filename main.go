package main

import (
	"gcp-golang/resources/gke"
	"gcp-golang/resources/projects"
)

func main() {
	projects, err := projects.Get()
	if err != nil {
        panic(err)
    }
	for _, project := range projects {
		gke.Get(project)
	}
}
