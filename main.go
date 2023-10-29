package main

import (
	"fmt"
	"net/http"

	"projectx/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// ROTAS
	r.StaticFS("/portal", http.Dir("internal/static"))
	routes.SetupSearch(r)
	routes.SetupExport(r)

	// Power on GIN server
	serverAddr := "127.0.0.1:8080"
	fmt.Printf("Starting server at %s\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		fmt.Println("Error starting Gin server:", err)
	}

	fmt.Println("Server stopped.")
}
