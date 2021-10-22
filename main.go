package main

import (
	"gcp-golang/resources/gke"
	"gcp-golang/resources/projects"
)

func main() {
	projects.Get()
	gke.Get()
}
