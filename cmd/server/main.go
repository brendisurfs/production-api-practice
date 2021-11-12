package main

import "fmt"

// App - struct that contains things such as ptrs to db connections
type App struct{}

// Run - set up our application
func (app *App) Run() error {
	fmt.Println("setting up app")
	return nil
}

func main() {
	fmt.Println("Go RES API")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")

	}
}
