package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mvpbv/boot-fetch/internal/database"
)

func (apiCfg *apiConfig) addDuel(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name   string `json:"name"`
		Racer1 string `json:"racer_1"`
		Racer2 string `json:"racer_2"`
		Level  int    `json:"level"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if params.Name == "" || params.Racer1 == "" || params.Racer2 == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	uui1, err := uuid.Parse(params.Racer1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uui2, err := uuid.Parse(params.Racer2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if params.Level < 1 {
		http.Error(w, "Level value must be greater than zero", http.StatusBadRequest)
	}
	DuelXp := int32(xp_at_level(params.Level))

	vals, err := apiCfg.DB.CreateDuel(r.Context(), database.CreateDuelParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Racer1ID:  uui1,
		Racer2ID:  uui2,
		RaceXp:    DuelXp,
		StartTime: time.Now().UTC(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, vals)
}
