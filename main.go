package main

import (
	"fmt"

	"todo/app/controllers"
	"todo/app/models"
)

func main()  {
	fmt.Println(models.Db)

	controllers.StartMainServer() //go run main.goでサーバー起動.

	// //userの参照
	// user, _ := models.GetUserByEmail("sa10sh12un08@icloud.com")
	// fmt.Println(user)

	// //セッションの作成と取得
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// //セッションの確認
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)


}
