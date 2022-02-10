package router

import (
  "net/http"
  "github.com/gorilla/mux"
)

type CsvController interface {
  PostCsv(w http.ResponseWriter, r *http.Request)
}

func Setup(c CsvController) *mux.Router {
  r := mux.NewRouter()

  api := r.PathPrefix("/api")

  api.HandleFunc("/csv", c.PostCsv).
    Methods(http.MethodPost).Name("PostCsv")

  return r
}
