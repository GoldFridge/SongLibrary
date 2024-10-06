package SongHandlers

import (
	"SongLibrary/internal/database"
	"SongLibrary/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

// CreateSongRequest @Description
// @Accept json
// @Produce json
// @Success 200 {object} CreateSongRequest
type CreateSongRequest struct {
	Song  string `json:"song"`
	Group string `json:"group"`
}

// CreateSong godoc
// @Summary Create a new song
// @Description Create a new song in the database
// @Tags songs
// @Accept json
// @Produce json
// @Param song body CreateSongRequest true "Song to create"
// @Success 201 {object} models.Song
// @Failure 400 {string} string "Invalid JSON format or missing group/song"
// @Failure 500 {string} string "Failed to create song"
// @Router /songs/create [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("CreateSong: Invalid method")
		return
	}

	var newSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Println(err)
		return
	}

	if newSong.Group == "" || newSong.Song == "" {
		http.Error(w, "Missing group or song", http.StatusBadRequest)
		log.Println("Missing group or song")
		return
	}

	if err := db.Create(&newSong).Error; err != nil {
		http.Error(w, "Failed to create song", http.StatusInternalServerError)
		log.Println("Error creating song:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(newSong)
	if err != nil {
		log.Println("Error encoding song:", err)
		return
	}
}
