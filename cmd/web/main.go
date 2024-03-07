package main

import (
	"firstgo-project/internal/handlers"
	"fmt"
	"net/http"
)

// const -> variable qui ne peut pas changer
const port = ":8080"

func main() {
	// urls
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	//le port ou ou ecoute
	http.ListenAndServe(port, nil)

}
