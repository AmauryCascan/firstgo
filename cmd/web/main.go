package main

import (
	"fmt"
	"net/http"

	"github.com/AmauryCascan/firstgo/internal/handlers"
)

// const -> variable qui ne peut pas changer
const port = ":8000"

func main() {
	// urls
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("(http://localhost: 8080) - Server started on port", port)
	//le port ou ou ecoute
	http.ListenAndServe(port, nil)
	
}
