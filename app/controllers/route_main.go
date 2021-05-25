package controllers

import (
	"fmt"
	"log"
	"net/http"
	"todo/app/models"
)

//ハンドラー
func top(w http.ResponseWriter, r *http.Request)  { //パターン
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// t.Execute(w, "hi")//実行
	//第二引数をhtmlで{{.}}として呼び出せる

	_, err := session(w, r)//クッキーの取得
	if err != nil {
		//レスポンスライター　//セッションがない時だけtopへ
		generateHTML(w, "hi", "layout", "public_navbar", "top")

	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request)  {
	sess, err := session(w,r)//セッションでクッキーを取得・クッキーの値とセッションが一致すればtodoへ
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {//そうでなければtopへ
		user, err := sess.GetUserBySession()//セッションのuserIDを使って一致するuserを取得
		if err != nil {
			log.Panicln(err)
		}
		todos, _ := user.GetTodoByUser()//上で作成したtodoの一覧を取得。
		//todosをuserのストラクトに入れてまとめる => users.go
		user.Todos = todos
		//userの情報をテンプレートに渡したので第二引数をuserにする
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

//todoハンドラ
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}


func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)

		}
		//userの取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//nameで指定されているcontentを取得
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)

		}
		fmt.Println(content)
		http.Redirect(w, r, "/todos", 302)
	}
}


func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)

		}
		//todoを取得
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(t)
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int)  {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t. UpdateTodo(); err != nil {
			log.Panicln(err)
		}
		http.Redirect(w, r, "/login", 302)
	}

}
 
func todoDelete(w http.ResponseWriter, r *http.Request, id int)  {
	sess, err := session(w,r)

	if err != nil {
		http.Redirect(w, r, "/login", 302)

	}else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}

}