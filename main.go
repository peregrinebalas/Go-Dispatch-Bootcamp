package main

import(
  "net/http"
  "log"

  "github.com/peregrinebalas/go-dispatch-bootcamp/controller"
  "github.com/peregrinebalas/go-dispatch-bootcamp/router"
  "github.com/peregrinebalas/go-dispatch-bootcamp/usecase"
  "github.com/peregrinebalas/go-dispatch-bootcamp/service"
)

func main() {
  pokemonService := service.New()
  pokemonUsecase := usecase.New(pokemonService)
  pokemonController := controller.New(pokemonUsecase)
  httpRouter := router.Setup(pokemonController)

  http.ListenAndServe(":8080", httpRouter)

  log.Printf("starting server on port :8080")
}
