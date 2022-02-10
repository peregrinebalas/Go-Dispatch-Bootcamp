package service

import (
  "log"
  "encoding/csv"
)

type PokemonMap map[int]model.Pokemon

var pokemonOrder []int = []int{1,2}

var db PokemonMap map[int]model.Pokemon
