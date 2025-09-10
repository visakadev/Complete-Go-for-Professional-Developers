package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/visakadev/go/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running out app")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port %d\n", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}

// Handler
// Response Handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
