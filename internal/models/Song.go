package models

// Song @Description
// @Accept json
// @Produce json
// @Success 200 {object} Song
type Song struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Group       string `gorm:"size:255" json:"group"`
	Song        string `gorm:"size:255" json:"song"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}
