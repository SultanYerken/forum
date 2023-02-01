package main

import (
	"forum/internal/app"
	"forum/internal/controller"
	"forum/internal/controller/httpserver"
	"forum/internal/repository"
	"forum/internal/repository/tables"
	"forum/internal/usecase"
	"log"
)

func main() {
	app.CreateDB()
	db, err := repository.NewSQLiteDB()
	if err != nil {
		log.Println(err)
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
