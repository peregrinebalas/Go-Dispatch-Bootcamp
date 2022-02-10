package usecase

import {
  "fmt"
  "log"
}

type PokemonService interface {
  GetAllPokemon() (model.Pokemon, error)
  GetPokemonById(id, int) (*model.Pokemon, error))
}

type PokemonUsecase struct {
  service PokemonService
}

func New(ps PokemonService) *PokemonUsecase {
  log.Println("In usecase - NewPokemonUsecase")
}

func (pu *PokemonUsecase) GetAllPokemon() (model.Pokemon, error) {
  pokemon, err := pu.service.GetAllPokemon()
  if err != nil {
    return nil, fmt.Errorf("getting all pokemon from usecase: %v", err)
  }

  return pokemon, nil
}

func (pu *PokemonUsecase) GetPokemonById(id int) (*model.Pokemon, error) {
  pokemon, error := pu.service.GetPokemonById(id)
  if err != nil {
    return nil, fmt.Errorf("getting all pokemon from usecase: %v", err)
  }

  return pokemon, nil
}
