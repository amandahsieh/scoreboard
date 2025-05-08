package scoreboard

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Service Service
}

func NewScoreboardHandler(service *Service) *Handler {
	return &Handler{Service: *service}
}

func (ctrl *Handler) ListHandler(w http.ResponseWriter) {
	scoreboards, err := ctrl.Service.ListAllScoreboards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(scoreboards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}
	// Decode incoming request and save to defined struct
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Name == "" {
		log.Println("Received request with empty name")
		return
	}
	log.Println("Received request with name:", request.Name)
	scoreboard, err := ctrl.Service.CreateScoreboard(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(scoreboard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Handler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	scoreboard, err := ctrl.Service.GetScoreboardByID(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(scoreboard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func (ctrl *ScoreboardHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		http.Error(w, "Invalid ID format", http.StatusBadRequest)
//		return
//	}
//
//	var request struct {
//		Name string `json:"name"`
//	}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	scoreboard, err := ctrl.Service.UpdateScoreboard(int32(id), request.Name)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	if err := json.NewEncoder(w).Encode(scoreboard); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ctrl *ScoreboardHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		http.Error(w, "Invalid ID format", http.StatusBadRequest)
//		return
//	}
//
//	err = ctrl.Service.DeleteScoreboard(int32(id))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
