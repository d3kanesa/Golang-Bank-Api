package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"apiProject/internal/middleware"
)

func Handler(r *chi.Mux) {
	//global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router){
		//Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
		router.Post("/deposit", AddCoins)
		router.Post("/transfer", TransferCoins)
	})

	r.Post("/createAccount", CreateAccount)
}