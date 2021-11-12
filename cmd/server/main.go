package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/brendisurfs/go-rest-api/internal/transport/http"
)

// App - struct that contains things such as ptrs to db connections
type App struct{}

// Run - set up our application
func (app *App) Run() error {
	fmt.Println("setting up app")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go RES API")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
	}
}
