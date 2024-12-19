package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wafieabiyya/goropher-social.git/internal/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	store := store.NewStorage(nil)

	cfg := config{
		addr: port,
	}

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
