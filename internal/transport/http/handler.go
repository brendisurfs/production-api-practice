package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/brendisurfs/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores the pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our api.
type Response struct {
	Message string
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
	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET") // can specify which methods can access which route.
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// this will format so it looks like actual json, not mush.
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		// take note of this, this is cool.
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive and well :0"}); err != nil {
			log.Fatal(err)
		}
	})
}

// GetAllComments - retrieves all comments from the comment service.
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "failed to retrieve all comments.")
	}

	fmt.Fprintf(w, "%+v", comments)
}

// PostComment - adds a new comment.
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "failed to post new comment.")
	}
	// if successful, return the comment.
	fmt.Fprintf(w, "%+v", comment)
}

// GetComment - retrieve comment by id.
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// retrieve the id that we wish to get from the comment list.
	vars := mux.Vars(r)
	// takes out from the request parms.BUT returns a string.
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "unable to parse uint from id.")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "error retrieving comment by id.")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates comment by id
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "failed to update comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// DeleteComment - deletes comment by id.
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "failed to parse uint from id.")
	}

	err = h.Service.DeleteComment(uint(commentID)) //always remember to cast!
	if err != nil {
		fmt.Fprintf(w, "failed to delete comment by comment id.")
	}

	fmt.Fprintf(w, "successfully deleted comment.")
}
