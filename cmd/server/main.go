package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brendisurfs/go-rest-api/internal/comment"
	"github.com/brendisurfs/go-rest-api/internal/database"
	transportHTTP "github.com/brendisurfs/go-rest-api/internal/transport/http"
	"github.com/joho/godotenv"
)

// App - struct that contains things such as ptrs to db connections
type App struct{}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not start server because loading env variables failed. : ", err)
	}
}

// Run - set up our application
func (app *App) Run() error {
	fmt.Println("setting up app")

	// starts our connection to the specified db.
	var err error
	db, err := database.NewDB()
	if err != nil {
		return err
	}

	// database migration exec.
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	// starts a new service with db passed in as the arg.
	commentService := comment.NewService(db)

	// creates a new handler with a struct of the router and service.
	handler := transportHTTP.NewHandler(commentService)
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
		fmt.Println("Error starting up REST API ", err)
	}
}
