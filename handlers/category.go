package handlers

import (
	categorydto "backend_project/dto/category"
	dto "backend_project/dto/result"
	"backend_project/models"
	"backend_project/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Declare handler struct here ...
type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

// Declare HandlerUser function here ...
func HandlerCategory(CategoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}

// Declare FindUsers method here ...
func (h *handlerCategory) FindCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category, err := h.CategoryRepository.FindCategory()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: category}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
func (h *handlerCategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(category)}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handlerCategory) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(categorydto.CreateCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db user
	category := models.Category{
		Name: request.Name,
		// Subscribe: request.Subscribe,
	}

	data, err := h.CategoryRepository.CreateCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(data)}
	json.NewEncoder(w).Encode(response)
}

// Declare convertResponse function here ...
func convertResponseCategory(u models.Category) categorydto.CategoryResponse {
	return categorydto.CategoryResponse{
		ID:   u.ID,
		Name: u.Name,
		// Subscribe: u.Subscribe,
	}
}

// Write this code
func (h *handlerCategory) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(categorydto.UpdateCategoryRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category, _ := h.CategoryRepository.GetCategory(int(id))

	if request.Name != "" {
		category.Name = request.Name
	}

	data, err := h.CategoryRepository.UpdateCategory(category, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(data)}
	json.NewEncoder(w).Encode(response)
}

// // Write this code
func (h *handlerCategory) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CategoryRepository.DeleteCategory(category, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(data)}
	json.NewEncoder(w).Encode(response)
}
