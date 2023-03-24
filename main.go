package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rastros-da-mata/database"
	"syscall"
	"time"
)

func main() {

	// carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar as variáveis de ambiente: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Inicializando roteador
	app := &App{
		DB: map[string]*mongo.Collection{
			"fruits":     db.Collection("fruits"),
			"vegetables": db.Collection("vegetables"),
			"greens":     db.Collection("greens"),
		},
	}

	router := mux.NewRouter()

	// Criando rotas
	router.HandleFunc("/api/fruits", app.createFruitHandler).Methods("POST")
	router.HandleFunc("/api/fruits/{id}", app.readFruitHandler).Methods("GET")
	router.HandleFunc("/api/fruits/{id}", app.updateFruitHandler).Methods("PUT")
	router.HandleFunc("/api/fruits/{id}", app.deleteFruitHandler).Methods("DELETE")
	router.HandleFunc("/api/fruits", app.GetAllFruitsHandler).Methods("GET")

	router.HandleFunc("/api/vegetables", app.createVegetableHandler).Methods("POST")
	router.HandleFunc("/api/vegetables/{id}", app.readVegetableHandler).Methods("GET")
	router.HandleFunc("/api/vegetables/{id}", app.updateVegetableHandler).Methods("PUT")
	router.HandleFunc("/api/vegetables/{id}", app.deleteVegetableHandler).Methods("DELETE")
	router.HandleFunc("/api/vegetables", app.GetAllVegetablesHandler).Methods("GET")

	router.HandleFunc("/api/greens", app.createGreenHandler).Methods("POST")
	router.HandleFunc("/api/greens/{id}", app.readGreenHandler).Methods("GET")
	router.HandleFunc("/api/greens/{id}", app.updateGreenHandler).Methods("PUT")
	router.HandleFunc("/api/greens/{id}", app.deleteGreenHandler).Methods("DELETE")
	router.HandleFunc("/api/greens", app.GetAllGreensHandler).Methods("GET")

	srv := &http.Server{
		Handler:      handlers.CORS()(router),
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Server started on port :%s\n", os.Getenv("PORT"))

	// Capturando sinal para finalizar servidor
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	log.Println("Shutting down server...")

	// Encerrando servidor
	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Println("Server stopped.")

}