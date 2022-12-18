package dao

//此处大有名堂，深究

import (
	"awesomeProject/model"
	"log"
)

func InsertUser(user model.Users) (err error) {

	_, err = DB.Exec("insert into user (username,password) values (?,?)", user.Username, user.Password)
	if err != nil {
		log.Println(err)
		return
	}
	return

}

func SelectUser(username string) (err error, user model.Users) {
	//
	//query 和 scan 的 几个数据必须一样
	//
	row := DB.QueryRow("SELECT  id,username,password FROM user WHERE username =?", username)
	if err = row.Err(); err != nil {
		return
	}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	return

}

func Changepass(username interface{}, newpass string) (err error) {
	_, err = DB.Exec("update user set password =? where username = ? ", newpass, username)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
