package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main()  {
  serveWeb()
}

func serveWeb()  {
  r := mux.NewRouter()

  r.HandleFunc("/", serveContent)
  r.HandleFunc("/{my_custom_param}", serveContent)

  http.Handle("/", r)
  http.ListenAndServe(":3000", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request)  {
  urlParams := mux.Vars(r)
  myParameter := urlParams["my_custom_param"]

  if myParameter == "" {
    myParameter = "Home"
  }

  w.Write([]byte("Hello " + myParameter))
}
