package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"todo/app/models"
	"todo/config"
)

//可変調引数
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string)  {
	var files []string
	//filenamesの値を取り出してfilesに格納する
	for _, file := range filenames { // fmt.Sprintfでしたのパスに入れてfilesに格納する
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
	//htmlから{{define "名前"}}として呼び出すときはExecuteTemplateで明示できに呼び出す"layout"
}

//route_auth.goのcookieとリンクさせる
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error)  {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}//valueにはuuidがある
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid sess")
		}
	} 
	return sess, err//uuidがあればerrが帰らない
}
//urlの正規表現のパターンをコンパイルする
var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")//数値の繰り返し

	//ハンドラ関数を返す関数　　  
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	//パターンとして覚える　　001ハンドラ関数を返す
	return func(w http.ResponseWriter, r *http.Request) {
		//validPathとURlがマッチしたところをスライスで取得　q
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		//数値型ならAtoiで変換してqiに入れる
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		//実行
		fn(w, r, qi)
	}
}

//サーバーの立ち上げ
func StartMainServer() error  {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
										//staticファイルはないので
										
	//HandleFuncはレスポンスライターと流クエストを引数として受け取るのでparseURLで...001
	http.HandleFunc("/", top) //urlの取得　/がurl topが接続先
	http.HandleFunc("/signup", signup)//signupのハンドラをurlに登録
	http.HandleFunc("/login", login)//loginのハンドラをurlに登録
	http.HandleFunc("/authenticate", authtenticate)//authtenticateのハンドラをurlに登録
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	//ハンドラ関数をチェインさせている
	//末尾の/はurlが一致しないといけないことを意味している

//herou psql
	port := os.Getenv("PORT")

	//httpのメソット立ち上げに使う　コロンとポート番号を渡す  2はデフォルトのマルチプレクサ
	return http.ListenAndServe(":" + port, nil)
}