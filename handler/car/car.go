package car

import (
	"encoding/json"
	"golangSecond/models"
	"golangSecond/service"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := h.service.GetCarByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	// Write response body here
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Errore Writting response:", err)
		return
	}

}
func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"
	resp, err := h.service.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	// Write response body here
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Errore Writting response:", err)
		return
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	createdCar, err := h.service.CreateCar(ctx, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}

	responseBody, err := json.Marshal(createdCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error while writing response:", err)
	}
}
func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error reading request body:", err)
		return
	}
	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}
	updatedCar, err := h.service.UpdateCar(ctx, id, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error updating car:", err)
		return
	}
	responseBody, err := json.Marshal(updatedCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error while writing response:", err)
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedCar, err := h.service.DeleteCar(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error deleting car:", err)
		return
	}

	responseBody, err := json.Marshal(deletedCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error while writing response:", err)
	}
}
