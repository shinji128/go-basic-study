package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/shinji128/go-basic-study/myapi/models"
	"github.com/shinji128/go-basic-study/myapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

  // 1. テスト結果として期待する値を定義
	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	}

  // 2. テスト対象となる関数を実行
  // -> 結果がgot に格納される
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
    // 関数の実行そのものに失敗した場合には、テストも失敗させる
		t.Fatal(err)
	}


  // 3. 2 の結果と1 の値を比べる
  // Errorは実行されても他のテストを続行する
	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
  // t.Fatal もt.Errorf も実行されずに関数が終わった場合にはテスト成功
}
