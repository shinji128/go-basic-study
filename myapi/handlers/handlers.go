package handlers

import (
  "fmt"
  "io"
  "net/http"
  "strconv"
  "errors"
  "encoding/json"
  "github.com/shinji128/go-basic-study/myapi/models"
  "github.com/gorilla/mux"
)
// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello, world!\n")
}
// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
  length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    http.Error(w, "cannot get content length\n", http.StatusBadRequest)
    return
  }
  reqBodybuffer := make([]byte, length)

  // ファイルの読み込みが終わるとEOF(EndOfFile)というエラーが発生するがそれは正常に処理が終了したということ
  // errors.Isは第一引数のエラーが第二引数のエラーと同じかどうかを判定する
  if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
    http.Error(w, "fail to get request body\n", http.StatusBadRequest)
    return
  }
  defer req.Body.Close()

  article := models.Article{}
  jsonData, err := json.Marshal(article)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }

  w.Write(jsonData)
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
