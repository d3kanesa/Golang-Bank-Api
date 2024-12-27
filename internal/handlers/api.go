package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"apiProject/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
		router.Post("/deposit", AddCoins)
		router.Post("/transfer", TransferCoins)
	})

	r.Post("/createAccount", CreateAccount)
}