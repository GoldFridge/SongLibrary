package SongHandlers

import (
	"SongLibrary/internal/database"
	"SongLibrary/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// UpdateSong godoc
// @Summary      Update a song by ID
// @Description  Update an existing song's data in the database
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Song ID"
// @Param        song body      models.Song true  "Updated song data"
// @Success      200  {object}  models.Song
// @Failure      400  {string}  string "Invalid song ID or bad request"
// @Failure      404  {string}  string "Song not found"
// @Failure      500  {string}  string "Error updating the song"
// @Router       /songs/update/{id} [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 4 || parts[1] != "songs" || parts[2] != "update" {
		log.Println("Invalid URL pattern")
		http.Error(w, "Invalid URL pattern", http.StatusBadRequest)
		return
	}

	songID, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Println("Invalid song ID")
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	var updatedSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db := database.DB
	var song models.Song
	if err := db.First(&song, songID).Error; err != nil {
		log.Println("Song not found")
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	song.Group = updatedSong.Group
	song.Song = updatedSong.Song
	song.ReleaseDate = updatedSong.ReleaseDate
	song.Link = updatedSong.Link
	song.Text = updatedSong.Text

	if err := db.Save(&song).Error; err != nil {
		log.Printf("Error updating song: %v", err)
		http.Error(w, "Error updating the song", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(song); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
