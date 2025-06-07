package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Example struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetExamples handles GET /examples
func GetExamples(w http.ResponseWriter, r *http.Request) {
	examples := []Example{}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "GET /examples - List all examples",
		"data":    examples,
	})
}

// GetExample handles GET /examples/{id}
func GetExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "GET /examples/" + id,
		"data":    Example{ID: id, Name: "Example " + id},
	})
}

// CreateExample handles POST /examples
func CreateExample(w http.ResponseWriter, r *http.Request) {
	var input Example
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// In a real application, you would save the example to a database here
	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Example created successfully",
		"data":    input,
	})
}

// UpdateExample handles PUT /examples/{id}
func UpdateExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input Example
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// In a real application, you would update the example in the database
	input.ID = id
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Example updated successfully",
		"data":    input,
	})
}

// PatchExample handles PATCH /examples/{id}
func PatchExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// In a real application, you would apply partial updates to the example in the database
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Example patched successfully",
		"id":      id,
		"updates": updates,
	})
}

// DeleteExample handles DELETE /examples/{id}
func DeleteExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// In a real application, you would delete the example from the database here
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Example deleted successfully",
		"id":      id,
	})
}

// HealthCheck handles GET /health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Helper function to respond with error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// SetupRouter configures the routes and returns a router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	// Example resource endpoints
	examples := r.PathPrefix("/examples").Subrouter()
	examples.HandleFunc("", GetExamples).Methods("GET")
	examples.HandleFunc("/{id}", GetExample).Methods("GET")
	examples.HandleFunc("", CreateExample).Methods("POST")
	examples.HandleFunc("/{id}", UpdateExample).Methods("PUT")
	examples.HandleFunc("/{id}", PatchExample).Methods("PATCH")
	examples.HandleFunc("/{id}", DeleteExample).Methods("DELETE")

	return r
}
