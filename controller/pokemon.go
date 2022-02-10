package controller

import (
  // "encoding/json"
  "errors"
  "fmt"
  "log"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)
// validate and transform data from request


type usecase interface {
  GetAllPokemon() (*model.Pokemons, error)
  GetPokemonById(id int) (*model.Pokemon, error)
}

type pokemonController struct {
  usecase usecase
}

func New(uc usecase) *pokemonController {
  return &pokemonController{
    usercase: uc,
  }
}

func (pc *pokemonController) GetAllPokemon(w http.ResponseWriter, r *http.Request) {
  pokemon, err := pc.usecaseGetAllPokemon()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "error getting pokemon")

    log.Fatalf("getting all pokemon %v", err)
  }

  if len(pokemon) == 0 {
    log.Println("no pokemon found")
    w.WriteHeader(http.StatusNotFound)
  }

  jsonData, err := json.Marshal(pokemon)
  if err != nil {
    log.Println("error marshalling pokemon")
    w.WriteHeader(http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusTeapot)
  w.Write(jsonData)
}

func (pc *pokemonController) GetPokemonById(w http.ResponseWriter, r *http.Request) {
  log.Println("in controller - GetPokemonById")

  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "invalid id: %v", err)

    log.Fatalf("converting id param into an int: %v", err)
  }

  pokemon, err := pc.usecase.GetPokemonById(id)
  if err != nil {
    switch {
      case errors.Is(err, "Not Found"), errors.Is(err, "EmptyData")
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "pokemon not found")

      case errors.Is(err, "Not Init")
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "data not initialized")

        log.Fatalf("getting pokemon: %v", err)
    }
  }

  if pokemon == nil {
    log.Println("no pokemon found")
    w.WriteHeader(http.StatusNotFound)
  }

  jsonData, err := json.Marshal(pokemon)
  if err != nil {
    log.Println("error marshalling pokemon")
    w.WriteHeader(http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusTeapot)
  w.Write(jsonData)
}
