package controller

import (
  "encoding/json"
  "encoding/csv"
  "errors"
  "fmt"
  "log"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)
// validate and transform data from request


type usecase interface {

}

type csvController struct {
  usecase usecase
}

func New(uc usecase) *csvController {
  return &csvController{
    usercase: uc,
  }
}

func (cc *csvController) PostCsv(w http.ResponseWriter, r *http.Request) {
  log.Println("in controller - PostCsv")

  vars := mux.Vars(r)
  body := mux.Body(r)

  id, err := strconv.Atoi(vars["id"])
  if vars["id"] && err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "invalid id: %v", err)

    log.Fatalf("converting id param into an int: %v", err)
  }

  csv, err := cc.usecase.ReadCsv(body)
}
