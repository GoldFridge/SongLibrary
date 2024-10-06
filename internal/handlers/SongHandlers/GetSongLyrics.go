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

// GetSongLyrics godoc
// @Summary      Get song lyrics with pagination by verses
// @Description  Retrieve the lyrics of a song, paginated by verses
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id      path     int     true  "Song ID"
// @Param        page    query    int     false "Page number (default is 1)"
// @Param        limit   query    int     false "Number of verses per page (default is 2)"
// @Success      200     {object} map[string]interface{} "Paginated verses"
// @Failure      400     {string} string "Invalid input or song not found"
// @Failure      500     {string} string "Error fetching the song lyrics"
// @Router       /songs/lyrics/{id} [get]
func GetSongLyrics(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 4 || parts[1] != "songs" || parts[2] != "lyrics" {
		http.Error(w, "Invalid URL pattern", http.StatusBadRequest)
		return
	}

	songID, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	db := database.DB
	var song models.Song
	if err := db.First(&song, songID).Error; err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	verses := strings.Split(song.Text, "\n\n")

	page := 1
	limit := 2

	if p := r.URL.Query().Get("page"); p != "" {
		page, err = strconv.Atoi(p)
		if err != nil || page < 1 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil || limit < 1 {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(verses) {
		http.Error(w, "Page out of range", http.StatusBadRequest)
		return
	}
	if end > len(verses) {
		end = len(verses)
	}

	response := map[string]interface{}{
		"song":   song.Song,
		"group":  song.Group,
		"page":   page,
		"limit":  limit,
		"total":  len(verses),
		"verses": verses[start:end],
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
