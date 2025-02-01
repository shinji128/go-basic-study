package main

import (
  "dbsample/models"
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

	const sqlStr = `
		SELECT title, contents, username, nice
		FROM articles;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
