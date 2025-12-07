package main

import (
	"fmt"
	"net/http"

	config "github.com/gitakash786/studentapi/internal"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()

	// setup Router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Student API is running"))
	})

	// setUp Server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("server started %s\n", cfg.HTTPServer.Addr)

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
