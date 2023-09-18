package handlers

import (
	filmsdto "backend_project/dto/films"
	dto "backend_project/dto/result"
	"backend_project/models"
	"backend_project/repositories"
	"context"
	"log"
	"os"

	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Declare handler struct here ...
type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

// Declare HandlerUser function here ...
func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

// Declare FindUsers method here ...
func (h *handlerFilm) FindFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.FilmRepository.FindFilm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// for i, p := range films {
	// 	films[i].Thumbnailfilm = os.Getenv("PATH_FILE") + p.Thumbnailfilm
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(id)
	// film.Thumbnailfilm = os.Getenv("PATH_FILE") + film.Thumbnailfilm

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// UserID := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") // add this code
	filepath := dataContex.(string)    
	log.Println(filepath)         // add this code
	// filename := dataContex.(string) // add this code

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// var category_id []int
	// for _, r := range r.FormValue("categoryId") {
	// 	if int(r-'0') >= 0 {
	// 		category_id = append(category_id, int(r-'0'))
	// 	}
	// }

	request := filmsdto.CreateFilmRequest{
		Title:         r.FormValue("title"),
		Thumbnailfilm: filepath,
		// Thumbnailfilm: filename,
		Price:         r.FormValue("price"),
		Linkfilmbuyed: r.FormValue("linkfilmbuyed"),
		Linkfilm:      r.FormValue("linkfilm"),
		CategoryID:    r.FormValue("category_id"),
		Description:   r.FormValue("description"),
	}
	// fmt.Println(request)
	// request := new(filmsdto.CreateFilmRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "onlinecinema"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// data form pattern submit to pattern entity db user
	film := models.Film{
		Title: request.Title,
		// Thumbnailfilm: filename,
		Thumbnailfilm: resp.SecureURL,
		Price:         request.Price,
		Linkfilmbuyed: request.Linkfilmbuyed,
		Linkfilm:      request.Linkfilm,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
	}

	data, err := h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	// fmt.Println(data)
	film, _ = h.FilmRepository.CreateFilm(data)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	// response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseFilm(u models.Film) models.Film {
	return models.Film{
		ID:            u.ID,
		Title:         u.Title,
		Thumbnailfilm: u.Thumbnailfilm,
		Price:         u.Price,
		Linkfilmbuyed: u.Linkfilmbuyed,
		Linkfilm:      u.Linkfilm,
		CategoryID:    u.CategoryID,
		Category:      u.Category,
		Description:   u.Description,
	}
}

// Write this code
func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(filmsdto.UpdateFilmRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, _ := h.FilmRepository.GetFilm(int(id))

	// if id != 0 {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.Thumbnailfilm != "" {
		film.Thumbnailfilm = request.Thumbnailfilm
	}
	if request.CategoryID != " " {
		film.CategoryID = request.CategoryID
	}

	if request.Description != "" {
		film.Description = request.Description
	}
	if request.Linkfilmbuyed != "" {
		film.Linkfilmbuyed = request.Linkfilmbuyed
	}
	if request.Linkfilm != "" {
		film.Linkfilm = request.Linkfilm
	}

	data, err := h.FilmRepository.UpdateFilm(film, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
	json.NewEncoder(w).Encode(response)
}
