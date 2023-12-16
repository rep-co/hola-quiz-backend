package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
    Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
        }
    }
}

type APIServer struct {
    listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
    return &APIServer{
        listenAddress: listenAddress,
    }
}

func (s *APIServer) Run() {
    router := mux.NewRouter()

    router.HandleFunc("/packs", makeHTTPHandlerFunc(s.handlePack))

    log.Println("JSON API server is listening on port: ", s.listenAddress)

    http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handlePack(
    w http.ResponseWriter, 
    r *http.Request,
) error {
    if r.Method == "GET" {
        return s.handleGetPack(w, r)
    }
    if r.Method == "POST" {
        return s.handleCreatePack(w, r)
    }
    if r.Method == "DELETE" {
        return s.handleDeletePack(w, r)
    }

    return fmt.Errorf("method is not allowed: %s", r.Method)
}

func (s *APIServer) handleGetPack(
    w http.ResponseWriter,
    r *http.Request,
) error {
    return nil
}

func (s *APIServer) handleCreatePack(
    w http.ResponseWriter, 
    r *http.Request,
) error {
    return nil
}

func (s *APIServer) handleDeletePack(
    w http.ResponseWriter, 
    r *http.Request,
) error {
    return nil
}


