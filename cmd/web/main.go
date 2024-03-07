package main

import (
	"firstgo-project/config"
	"firstgo-project/internal/handlers"
	"fmt"
	"net/http"
)

// const -> variable qui ne peut pas changer


func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()
	
	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplates(&appConfig)

	// urls
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("(http://localhost:8080) - Server started on port", appConfig.Port)
	//le port ou ou ecoute
	http.ListenAndServe(appConfig.Port, nil)
	

}
