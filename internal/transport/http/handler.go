package http

import (
	"fmt"
	"net/http"

	"github.com/brendisurfs/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores the pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our applications
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// retrieve the id that we wish to get from the comment list.
}
