package main

import "net/http"
import "log"

import "github.com/peregrinebalas/go-dispatch-bootcamp/controller"
import "github.com/peregrinebalas/go-dispatch-bootcamp/router"
import "github.com/peregrinebalas/go-dispatch-bootcamp/usecase"
import "github.com/peregrinebalas/go-dispatch-bootcamp/service"
// import "fmt"
// import "time"
// import "html"

func main() {
  pokemonService := service.New(nil)
  pokemonUsecase := usecase.New(pokemonService)
  pokemonController := controller.New(pokemonUsecase)
  httpRouter := router.Setup(pokemonController)

  http.ListenAndServe(":8080", httpRouter)
  // s := &http.Server{
  // 	Addr:           ":8080",
  // 	ReadTimeout:    10 * time.Second,
  // 	WriteTimeout:   10 * time.Second,
  // 	MaxHeaderBytes: 1 << 20,
  // }
  // http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
  //   fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  // })
  // http.ListenAndServe(s)
  log.Printf("starting server on port :8080")
}
