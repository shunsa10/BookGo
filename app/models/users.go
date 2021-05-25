package models

import (
	"log"
	"time"
)

//ユーザーを作成する
//ユーザーのストラクト
type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	Todos []Todo
}

//セッションのストラクト
type Session struct {
	ID int
	UUID string
	Email string
	UserID string
	CreatedAt time.Time
}


func (u *User) CreateUser() (err error) {
	//コマンドを作る
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	//コマンドの実行
	_, err = Db.Exec(cmd,
		 createUUID(),
		  u.Name,
		   u.Email,
		    Encrypt(u.Password),//パスをハッシュ値にする
			 time.Now())
	
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

//ユーザーの取得
func GetUser(id int) (user User, err error) {
	user = User{}
		//QueryRowは常にnil以外の値を返します。
	 //cmd, idでnil以外だったものをスキャンに通す
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

//ユーザーのアップデート
//(u *User)メソット
func (u *User) UpdateUser() (err error)  {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd,u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}


//userの削除

func (u *User) DeleteUser() (err error)  {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	//emailを元に参照する
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

//セッションを作るメソット
func (u *User) CreateSession() (session Session, err error)  {
	session = Session{}//宣言
	//セッションを作成するコマンド
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)` 

		_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		//取得するためのコマンド
	cmd2 := `select id, uuid, email, user_id, created_at
		from sessions where user_id = ? and email = ?`

		err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
			&session.ID,
			&session.UUID,
			&session.Email,
			&session.UserID,
			&session.CreatedAt)

			return session, err
}

//セッションがデータベースにあるか確認
func (sess *Session) CheckSession() (valid bool, err error) {
	//uuidが一致するもの
	cmd := `select id, uuid, email, user_id, created_at
		from sessions where uuid = ?`

		err = Db.QueryRow(cmd, sess.UUID).Scan(
			&sess.ID,
			&sess.UUID,
			&sess.Email,
			&sess.UserID,
			&sess.CreatedAt)

			//セッションの判定
			if err != nil {
				valid = false
				return
			}
			if sess.ID != 0 {
				valid = true
			}
			return valid, err
 }

 func (sess *Session) DeleteSessionByUUID() (err error)  {
	 cmd := `delete from sessions where uuid = ?`
	 _, err = Db.Exec(cmd, sess.UUID)
	 if err != nil {
		 log.Fatalln(err)
	 }
	 return err
 }

 func (sess *Session) GetUserBySession() (user User, err error)  {
	 user = User{}//userのuserIDと
	 cmd := `select id, uuid, name, email, created_at FROM users
	 where id = ?`//usersのidが一致するのもにしたい
	 err = Db.QueryRow(cmd, sess.UserID).Scan(
		 &user.ID,
		 &user.UUID,
		 &user.Name,
		 &user.Email,
		 &user.CreatedAt)
		 return user, err
 }