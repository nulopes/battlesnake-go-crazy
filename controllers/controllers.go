package controllers

import (
	"battlesnake-go-crazy/engine"
	"battlesnake-go-crazy/models"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	engine engine.Engine
}

func NewHandler(engine engine.Engine) *Handler {
	return &Handler{engine: engine}
}

func (h *Handler) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	response := h.engine.Info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func (h *Handler) HandleStart(_ http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	h.engine.Start(state)
}

func (h *Handler) HandleMove(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := h.engine.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func (h *Handler) HandleEnd(_ http.ResponseWriter, r *http.Request) {
	state := models.GameState{}

	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	h.engine.End(state)
}
