package SongHandlers

import (
	"SongLibrary/internal/database"
	"SongLibrary/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetSong godoc
// @Summary      Get all songs with filtering and pagination
// @Description  Fetch songs from the database with optional filtering on all fields and pagination
// @Tags         songs
// @Produce      json
// @Param        group        query    string  false  "Filter by group"
// @Param        song         query    string  false  "Filter by song name"
// @Param        release_date query    string  false  "Filter by release date"
// @Param        page         query    int     false  "Page number (default is 1)"
// @Param        limit        query    int     false  "Number of items per page (default is 10)"
// @Success      200  {array}  models.Song
// @Failure      500  {string} string "Error fetching data"
// @Router       /songs [get]
func GetSong(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var songs []models.Song

	query := db.Table("songs")

	// Filtering
	group := r.URL.Query().Get("group")
	if group != "" {
		query = query.Where(`"group" LIKE ?`, "%"+group+"%")
	}

	song := r.URL.Query().Get("song")
	if song != "" {
		query = query.Where("song LIKE ?", "%"+song+"%")
	}

	releaseDate := r.URL.Query().Get("release_date")
	if releaseDate != "" {
		query = query.Where("release_date LIKE ?", "%"+releaseDate+"%")
	}

	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
		if page < 1 {
			page = 1
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
		if limit < 1 {
			limit = 10
		}
	}

	offset := (page - 1) * limit

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&songs).Error; err != nil {
		log.Printf("Error fetching data from database: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
