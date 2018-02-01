package main

import (
  "net/http"
)

func main()  {

  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/contact", contactHandler)

  http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("Hello World"))
}

func contactHandler(w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("contact Page"))
}
