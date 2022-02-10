package router

import (
  "net/http"

  "github.com/gorilla/mux"
)

type PokemonController interface {
  GetAllPokemon(w http.ResponseWriter, r *http.Request)
  GetPokemonById(w http.ResponseWriter, r *http.Request)
}

func Setup(c PokemonController) *mux.Router {
  r := mux.NewRouter()

  api := r.PathPrefix("/api")

  api.HandleFunc("/pokemon", c.GetAllPokemon).
    Methods(http.MethodGet).Name("GetAllPokemon")

  api.HandleFunc("/pokemon/{id}", c.GetPokemonById).
    Methods(http.MethodGet).Name("GetPokemonById")

  return r
}
