package controllers

import (
	"log"
	"net/http"
	"todo/app/models"
)


//ハンドラ
func signup(w http.ResponseWriter, r *http.Request) {
	//リクエストr　のメソットMethodを取得できる
	if r.Method == "GET" {
			_, err := session(w, r)
			if err != nil {
				generateHTML(w, nil, "layout", "public_navbar", "signup")
			} else {
				http.Redirect(w, r, "/todos", 302)
			}
	} else if r.Method == "POST" { //新しいユーザーとして登録
		err := r.ParseForm()//覚える
		if err != nil {
			log.Println(err)
		}
		user := models.User{//htmlのname = "名前"とリンク
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		//topページにリダイレクト
		http.Redirect(w, r, "/", 302)
		
	}
}
//一度データベースを消す


func login(w http.ResponseWriter, r *http.Request)  {
	_, err := session(w, r)
			if err != nil {
				generateHTML(w, nil, "layout", "public_navbar", "login")
			} else {
				http.Redirect(w, r, "/todos", 302)
			}
	
}

//ハンドラ　ユーザーの認証
func authtenticate(w http.ResponseWriter, r *http.Request)  {
	//loginの入力欄から取得
	err := r.ParseForm() 
	user, err := models.GetUserByEmail(r.PostFormValue("email"))//ここはemail
	if err != nil {
		log.Panicln(err)
		http.Redirect(w,r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		//passwordが一致したユーザーでセッションを作る
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}
		//作成したセッションを元にクッキーを作る
		cookie := http.Cookie {
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
	//ここまで書いたらhtmlのformのactionに登録
}


func logout(w http.ResponseWriter, r *http.Request)  {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}


