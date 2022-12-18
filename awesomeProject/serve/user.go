package serve

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

func SearchUserByName(username string) (err error, user model.Users) {
	//检查username是否存在
	err, user = dao.SelectUser(username)
	return
}

func AddUser(user model.Users) (err error) {
	err = dao.InsertUser(user)
	return
}

func FixPass(username interface{}, newpass string) (err error) {

	err = dao.Changepass(username, newpass)
	return
}
