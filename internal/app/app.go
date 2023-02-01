package app

import (
	"fmt"
	"forum/internal/controller"
	"forum/internal/controller/httpserver"
	"forum/internal/repository"
	"forum/internal/repository/tables"
	"forum/internal/usecase"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	CreateDB()

	db, err := repository.NewSQLiteDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	tables.CreateTables(db)
	repos := repository.NewRepository(db)
	usecases := usecase.NewUseCase(repos)
	handlers := controller.NewHandler(usecases)

	log.Print("Sever go http://localhost", httpserver.WritePORT())
	srv := new(httpserver.Server)
	err = srv.Run(httpserver.WritePORT(), handlers.Routes())
	if err != nil {
		log.Println(err.Error())
		return
	}
}
