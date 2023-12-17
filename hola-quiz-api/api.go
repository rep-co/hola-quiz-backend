package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	storage       Storage
}

func NewAPIServer(
	listenAddress string,
	storage Storage,
) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage:       storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/pack", makeHTTPHandlerFunc(s.handlePack))
	router.HandleFunc("/pack/{id}", makeHTTPHandlerFunc(s.handleGetPackByID))

	log.Println("JSON API server is listening on port: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handlePack(
	w http.ResponseWriter,
	r *http.Request,
) error {
	if r.Method == "GET" {
		return s.handleGetPackPreview(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreatePack(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeletePack(w, r)
	}

	return fmt.Errorf("method is not allowed: %s", r.Method)
}

func (s *APIServer) handleGetPackPreview(
	w http.ResponseWriter,
	r *http.Request,
) error {
	packs, err := s.storage.GetPacks()
	if err != nil {
		return err
	}
	packPreviews := ConvertPacksToPackPreviews(packs)
	return WriteJSON(w, http.StatusOK, packPreviews)
}

func (s *APIServer) handleGetPackByID(
	w http.ResponseWriter,
	r *http.Request,
) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id: %s", idStr)
	}

	pack, err := s.storage.GetPackByID(id)
	if err != nil {
		return err
	}

	questions := []*QuizQuestion{}
	results, err := GetQuizQuestions(pack.Category, "20", "medium")
	if err != nil {
		return err
	}

	for _, res := range results {
		quizQuestion := NewQuizQuestion(
			res.ID,
			res.Question,
			res.Description,
			res.Answers.ConvertToSlice(),
			res.CorrectAnswers.ConvertToSlice(),
			res.MultipleCorrectAnswers,
			res.Explanation)
		questions = append(questions, quizQuestion)
	}
	packQuiz := NewPackQuiz(pack, questions)

	return WriteJSON(w, http.StatusOK, packQuiz)
}

func (s *APIServer) handleCreatePack(
	w http.ResponseWriter,
	r *http.Request,
) error {
	createPackReq := new(CreatePackRequest)
	if err := json.NewDecoder(r.Body).Decode(createPackReq); err != nil {
		return err
	}

	pack := NewPack(createPackReq.Name,
		createPackReq.Description,
		createPackReq.Category,
		createPackReq.Emojis)
	if err := s.storage.CreatePack(pack); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, pack)
}

func (s *APIServer) handleDeletePack(
	w http.ResponseWriter,
	r *http.Request,
) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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
