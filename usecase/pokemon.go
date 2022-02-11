package usecase

import (
  "fmt"

  "github.com/peregrinebalas/go-dispatch-bootcamp/model"
)

type PokemonService interface {
  GetAllPokemon() ([]model.Pokemon, error)
  GetPokemonById(id int) (model.Pokemon, error)
}

type PokemonUsecase struct {
  service PokemonService
}

func New(ps PokemonService) *PokemonUsecase {
  return &PokemonUsecase{
    service: ps,
  }
}

func (pu *PokemonUsecase) GetAllPokemon() ([]model.Pokemon, error) {
  pokemon, err := pu.service.GetAllPokemon()
  if err != nil {
    fmt.Errorf("getting all pokemon from usecase: %v", err)
  }

  return pokemon, err
}

func (pu *PokemonUsecase) GetPokemonById(id int) (model.Pokemon, error) {
  pokemon, err := pu.service.GetPokemonById(id)
  if err != nil {
    fmt.Errorf("getting all pokemon from usecase: %v", err)
  }

  return pokemon, err
}
