package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"privyID/config"
	"privyID/internal/cake"
	"privyID/internal/constant"
)

func main() {
	cfg := config.New()

	db, err := sql.Open("mysql", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	validator := validator.New()
	router := mux.NewRouter()

	cakeRepo := cake.NewCakeRepository(db, constant.TableCake)
	cakeUseCase := cake.NewCakeUseCase(cakeRepo)

	cake.NewCakeHandler(router, validator, cakeUseCase)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	port := os.Getenv("PORT")

	fmt.Println("SERVER ON")
	fmt.Println("PORT :", port)
	log.Fatal(server.ListenAndServe())
}
