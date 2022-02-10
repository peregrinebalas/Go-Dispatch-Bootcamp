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
  pokemonService := service.New(nil)
  pokemonUsecase := usecase.New(pokemonService)
  pokemonController := controller.New(pokemonUsecase)
  httpRouter := router.Setup(pokemonController)

  http.ListenAndServe(":8080", httpRouter)

  log.Printf("starting server on port :8080")

  f, err := os.Open("pokemon.csv")
  if err != nil {
    log.Fatal("Unable to read input file", err)
  }
  defer f.Close()

  csvReader := csv.NewReader(f)
  records, err := csvReader.ReadAll()
  if err != nil {
    log.Fatal("Unable to parse file as CSV", err)
  }
}
