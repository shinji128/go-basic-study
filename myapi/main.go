package main

import (
  "log"
  "net/http"
  "github.com/shinji128/go-basic-study/myapi/handlers"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
  r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

  log.Println("server start at port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}
