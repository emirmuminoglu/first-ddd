package main

import (
	"github.com/emirmuminoglu/first-ddd/application"
	"github.com/emirmuminoglu/first-ddd/config"
	"github.com/emirmuminoglu/first-ddd/infrastructure/persistence"
	"github.com/emirmuminoglu/first-ddd/interfaces"
	"github.com/emirmuminoglu/lloyd"
)

var (
	userInteractor application.UserInteractor
	router         *lloyd.Lloyd
)

func initUserInteractor() {
	db, err := config.NewDBConnection()
	if err != nil {
		panic(err)
	}

	userInteractor.Repository = persistence.NewUserRepository(db, "first-ddd")
}

func initRouter() {
	router = lloyd.New(lloyd.Config{})
}

func init() {

	initUserInteractor()
	initRouter()
}

func main() {
	handler := interfaces.Handler{
		Interactor: userInteractor,
	}

	handler.Init(router)

	ln, err := config.NewListener()
	if err != nil {
		panic(err)
	}

	router.Serve(ln)
}
