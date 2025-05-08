package scoreboard

import (
	"encoding/json"
	"net/http"
)

type ScoreboardHandler struct {
	Service ScoreboardService
}

func NewScoreboardHandler(service *ScoreboardService) *ScoreboardHandler {
	return &ScoreboardHandler{Service: *service}
}

func (ctrl *ScoreboardHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
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

//func (ctrl *ScoreboardHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
//	var request struct {
//		Name string `json:"name"`
//	}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	log.Println("Received request with name:", request.Name)
//
//	scoreboard, err := ctrl.Service.CreateScoreboard(request.Name)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
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
//func (ctrl *ScoreboardHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		http.Error(w, "Invalid ID format", http.StatusBadRequest)
//		return
//	}
//
//	scoreboard, err := ctrl.Service.GetScoreboardByID(int32(id))
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
