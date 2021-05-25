package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error
// const (
// 	tableNameUser = "users"
// 	tableNameTodo = "todos"
// 	tableNameSession = "sessions"
// )

func init()  {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// //tableを無かったら作成
	// //idがINTEGER型でプライマリーキー　自動で増分
	// //null値を禁止　重複を禁止

	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME)`, tableNameUser)

	// 	Db.Exec(cmdU)

	// //todoのテーブル
	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	content TEXT,
	// 	user_id INTEGER,
	// 	created_at DATETIME)` , tableNameTodo)

	// 	Db.Exec(cmdT)

	// 	//セッションテーブルの作成
	// 	cmdS := fmt.Sprintf(`CREATE TABLE  IF NOT EXISTS %s (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		email STRING,
	// 		user_id INTEGER,
	// 		created_at DATETIME)`, tableNameSession)
	// 	Db.Exec(cmdS)	
}

//uuidを作成する関数
func createUUID() (uuidobj uuid.UUID)  {//uuidのUUID型を使っている
	uuidobj, _ = uuid.NewUUID()//作成
	return uuidobj//返す
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))//byteのスライスでplaintextを渡す
	return cryptext
}
