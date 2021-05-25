package models

import (
	"log"
	"time"
)


type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

//todoの作成　table
func (u *User) CreateTodo (content string) (err error) {
	cmd := `insert into todos (
		content,
		 user_id,
		  created_at) values (?, ?, ?)`
		  
		  _, err = Db.Exec(cmd, content, u.ID, time.Now())
		  
		  if err != nil {
			  log.Fatalln(err)
		}
		return err
	}


//todoのidの取得

func GetTodo(id int) (todo Todo, err error)  {
	//id取得のコマンド
	cmd := 	`select id, content, user_id, created_at from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}



//複数のtodoを取得する　スライスはリスト
func GetTodos() (todos []Todo, err error)  {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

			if err != nil {
				log.Fatalln(err)
			}
			todos = append(todos, todo) //追加する
	}
	rows.Close()

	return todos, err
}

//特定のuserのtodoを取得する
func (u *User) GetTodoByUser() (todos []Todo, err error)  {
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
			
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//todoの更新
func (t *Todo) UpdateTodo() error  {
	//contentとuser idを更新
	cmd := `update todos set content = ?, user_id = ?
	where id = ?` 

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//todoの削除

func (t *Todo)DeleteTodo() error {
	cmd := `delete from todos where id = ?`

	_, err := Db.Exec(cmd, t.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}