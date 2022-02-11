package service

import (
  "log"
  "encoding/csv"
  "os"
  "io"
  "strconv"

  "github.com/peregrinebalas/go-dispatch-bootcamp/model"
)

type PokemonService struct {
}

func New() *PokemonService {
  return &PokemonService{

  }
}

func (ps PokemonService) GetAllPokemon() ([]model.Pokemon, error) {
  pokedex, err := Pokedex()
  if err != nil {
    log.Fatal("could not fetch pokemon", err)
  }

  return pokedex, nil
}

func (ps PokemonService) GetPokemonById(id int) (model.Pokemon, error) {
  pokedex, err := Pokedex()
  if err != nil {
    log.Fatal("could not fetch pokemon", err)
  }

  var pokemon model.Pokemon
  for _, p := range pokedex {
        if id == p.Id {
            pokemon = p
        }
  }

  if (pokemon == model.Pokemon{}) {
    log.Fatal("could not find pokemon with id: %v", id)
  }

  return pokemon, nil
}

func Pokedex() (pokemon []model.Pokemon, error error) {
	headers := make(map[string]int)

  f, _ := os.Open("pokemon.csv")
  r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Unable to read input file", err)
		}
    if len(headers) == 0 {
      for i,v := range record {
        headers[v] = i
      }

      continue
    }
    pokemon = append(pokemon, model.Pokemon{
      Id: strToInt(record[headers["ID"]]),
      Name: record[headers["Name"]],
      Type1: record[headers["Type 1"]],
      Type2: record[headers["Type 2"]],
      Total: strToInt(record[headers["Total"]]),
      Hp: strToInt(record[headers["HP"]]),
      Attack: strToInt(record[headers["Attack"]]),
      Defense: strToInt(record[headers["Defense"]]),
      SpAtk: strToInt(record[headers["Sp. Atk"]]),
      SpDef: strToInt(record[headers["Sp. Def"]]),
      Speed: strToInt(record[headers["Speed"]]),
      Generation: strToInt(record[headers["Generation"]]),
      Legendary: strToBool(record[headers["Legendary"]]),
		})
	}
	return pokemon, nil
}


func strToInt(val string) int {
  x, _ := strconv.Atoi(val)

  return x
}

func strToBool(val string) bool {
  x, _ := strconv.ParseBool(val)

  return x
}
