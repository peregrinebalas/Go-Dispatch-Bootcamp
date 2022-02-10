package model

type Pokemon struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Type1 string `json:"type1"`
  Type2 string `json:"type2"`
  Total int `json:"total"`
  Hp int `json:"hp"`
  Attack int `json:"attack"`
  Defense int `json:"defense"`
  SpAtk int `json:"spAtk"`
  SpDef int `json:"spDef"`
  Speed int `json:"speed"`
  Generation int `json:"generation"`
  Legendary bool `json:"legendary"`
}
