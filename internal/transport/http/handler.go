package http

import (
	"encoding/json"
	"fmt"
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
	Error   string
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		// take note of this, this is cool.
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive and well :0"}); err != nil {
			panic(err)
		}
	})
}

// GetAllComments - retrieves all comments from the comment service.
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	SendErrorResponse(w, "failed to retrieve all comments.", err)

	if err := sendOKResponse(w, comments); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment.
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	SendErrorResponse(w, "failed to decode json body.", err)

	comment, err = h.Service.PostComment(comment)
	SendErrorResponse(w, "failed to post new comment.", err)

	if err := sendOKResponse(w, comment); err != nil {
		panic(err)
	}
}

// GetComment - retrieve comment by id.
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// retrieve the id that we wish to get from the comment list.
	vars := mux.Vars(r)

	// takes out from the request parms.BUT returns a string.
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	SendErrorResponse(w, "unable to parse UINT from ID.", err)

	comment, err := h.Service.GetComment(uint(i))
	SendErrorResponse(w, "error retrieving comment by ID.", err)

	if err := sendOKResponse(w, comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates comment by id
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	SendErrorResponse(w, "Failed to decode JSON Body", err)

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	SendErrorResponse(w, "Failed to parse uint from ID", err)

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	SendErrorResponse(w, "failed to update comment", err)

	if err := sendOKResponse(w, comment); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes comment by id.
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	SendErrorResponse(w, "failed to parse uint from id.", err)

	err = h.Service.DeleteComment(uint(commentID)) //always remember to cast!
	SendErrorResponse(w, "Failed to delete comment by ID", err)

	err = sendOKResponse(w, Response{Message: "successfully deleted"})
	if err != nil {
		panic(err)
	}
}

// this is the hardest func known to man. This is too smart.
// |
// v
func sendOKResponse(w http.ResponseWriter, resp interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

// SendErrrorMessage - a custom func to format and handle errors.
func SendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
			panic(err)
		}
	}
}
