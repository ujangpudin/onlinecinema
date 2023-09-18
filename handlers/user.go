package handlers

import (
	dto "backend_project/dto/result"
	usersdto "backend_project/dto/users"
	"backend_project/models"
	"backend_project/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Declare handler struct here ...
type handler struct {
	UserRepository repositories.UserRepository
}

// Declare HandlerUser function here ...
func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

// Declare FindUsers method here ...
func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	for i, p := range users {
		users[i].Image = os.Getenv("PATH_FILE") + p.Image
	}
	// for i, p := range users {
	// 	users[i].Transaction.Film.Thumbnailfilm = os.Getenv("PATH_FILE") + p.Film.Thumbnailfilm
	// }
	// fmt.Println(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	Image := user.Image
	user.Image = os.Getenv("PATH_FILE") + Image

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

// Declare convertResponse function here ...
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		// Password: u.Password,
		Phone: u.Phone,
		// Transaction: u.Transaction,
	}
}

// Write this code
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(r.Body)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)

	request := usersdto.UpdateUserRequest{
		Fullname: r.FormValue("fullname"),
		Image:    filename,
		// Password: r.FormValue("password"),
		Phone: r.FormValue("phone"),
	}
	// fmt.Println(request)

	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		fmt.Println("GetUser")

		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}

	// if request.Email != "" {
	// 	user.Email = request.Email
	// }

	// if request.Password != "" {
	// 	user.Password = request.Password
	// }

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Image != "" {
		user.Image = request.Image
	}

	data, err := h.UserRepository.UpdateUser(user, id)
	fmt.Println(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}
