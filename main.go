package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  dbUser := "docker"
  dbPassword := "docker"
  dbDatabase := "sampledb"
  dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

  // Open 関数の第二引数にDBのアドレスを渡して接続する
  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    fmt.Println(err)
  }

  defer db.Close()
  if err := db.Ping(); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("connect to DB")
  }
}
