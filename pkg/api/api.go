package api

import (
	"gcp-golang/pkg/api/routes"
	"net/http"
)

func StartServer() {
	routes.CallRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
