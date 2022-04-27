package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"ethservice/eth/block/storage"
	"ethservice/eth/block/transport"
	"syscall"
	"time"

	"ethservice/eth/config"

	"github.com/gorilla/mux"
)

type App struct {
	db     *sql.DB
	apikey string
}

func NewApp() *App {
	key, err := config.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	return &App{db: db, apikey: key}
}

func (app *App) Run(port string) {
	err := storage.InitTableDB(app.db)
	if err != nil {
		log.Fatal(err)
	}
	defer app.db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/blocks/{number}/total", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transport.Handler(w, vars["number"], app.db, app.apikey)
	})

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	// Listen for Ctrl+C to terminate server properly
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("Server Exited Properly")

}

func initDB() (*sql.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "mypassword", "user")

	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, err
	}

	start := time.Now()
	for db.Ping() != nil {
		if time.Now().After(start.Add(5 * time.Second)) {
			log.Print("failed to connect after 5 secs.")
			break
		}
	}

	if db.Ping() != nil {
		return nil, errors.New("Database connection failed")
	}

	return db, nil
}
