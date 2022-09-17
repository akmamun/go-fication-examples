package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-fication-examples/models"
	"go-fication-examples/repository"
	"net/http"
	"strconv"
)

type ExampleHandler struct {
	repo repository.ExampleRepo
}

func NewExampleHandler(repo repository.ExampleRepo) *ExampleHandler {
	return &ExampleHandler{
		repo: repo,
	}
}
func (h *ExampleHandler) GetExamplesListData(w http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))

	data, err := h.repo.GetExamples(int64(limit), int64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}

func (h *ExampleHandler) CreateData(w http.ResponseWriter, request *http.Request) {
	example := new(models.Example)
	err := json.NewDecoder(request.Body).Decode(&example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.repo.CreateExample(example)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(example)
}

func (h *ExampleHandler) GetOne(w http.ResponseWriter, request *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(request, "id"))

	data, err := h.repo.GetExample(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
