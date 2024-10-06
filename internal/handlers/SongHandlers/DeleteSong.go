package SongHandlers

import (
	"SongLibrary/internal/database"
	"SongLibrary/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// DeleteSong godoc
// @Summary      Delete a song by ID
// @Description  Delete a song from the database using its ID
// @Tags         songs
// @Param        id   path      int  true  "Song ID"
// @Success      200  {string}  string "Song deleted successfully"
// @Failure      400  {string}  string "Invalid song ID"
// @Failure      404  {string}  string "Song not found"
// @Failure      500  {string}  string "Error deleting the song"
// @Router       /songs/delete/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 4 || parts[1] != "songs" || parts[2] != "delete" {
		http.Error(w, "Invalid URL pattern", http.StatusBadRequest)
		return
	}

	songID, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&models.Song{}, songID).Error; err != nil {
		log.Printf("Error deleting song: %v", err)
		http.Error(w, "Error deleting the song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Song deleted successfully"))
	if err != nil {
		log.Println("Error writing response")
		return
	}
}
