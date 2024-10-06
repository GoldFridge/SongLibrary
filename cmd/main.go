package main

import (
	_ "SongLibrary/docs"
	"SongLibrary/internal/database"
	"SongLibrary/internal/handlers"
	"SongLibrary/internal/handlers/SongHandlers"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title SongLibraryAPI
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

func main() {
	database.InitDatabase()
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/songs", SongHandlers.GetSong)
	http.HandleFunc("/songs/create", SongHandlers.CreateSong)
	http.HandleFunc("/songs/delete/{id}", SongHandlers.DeleteSong)
	http.HandleFunc("/songs/update/{id}", SongHandlers.UpdateSong)
	http.HandleFunc("/songs/lyrics/{id}", SongHandlers.GetSongLyrics)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error start server")
	}
}
