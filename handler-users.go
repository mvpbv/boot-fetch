package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mvpbv/boot-fetch/internal/database"
)

func (apiCfg *apiConfig) updateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Nickname string `json:"nickname"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Go is braindead it's clearly there", http.StatusBadRequest)
		return
	}
	uui, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        uui,
		UpdatedAt: time.Now().UTC(),
		Nickname:  params.Nickname,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, "User updated")

}

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		BootName    string `json:"boot_name"`
		DiscordName string `json:"discord_name"`
		Nickname    string `json:"nickname"`
		Wizard      bool   `json:"wizard"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		BootName:    params.BootName,
		DiscordName: params.DiscordName,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Wizard:      params.Wizard,
		Nickname:    params.Nickname,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, user)

}
func (apiCfg *apiConfig) getUserStats(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Go is braindead it's clearly there", http.StatusBadRequest)
		return
	}

	uui, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stats, err := apiCfg.DB.GetUserProgress(r.Context(), uui)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, stats)
}
func (apiCfg *apiConfig) addMessage(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Insult string    `json:"insult"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insult, err := apiCfg.DB.CreateInsult(r.Context(), database.CreateInsultParams{
		ID:     uuid.New(),
		UserID: params.UserID,
		Insult: params.Insult,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, insult)

}

func (apiCfg *apiConfig) getArchWizard(w http.ResponseWriter, r *http.Request) {
	wizards, err := apiCfg.DB.GetWizards(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, wizards)

}
