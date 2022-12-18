package api

import (
	"awesomeProject/model"
	"awesomeProject/serve"
	"awesomeProject/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 这一段，有点问题
func Register(c *gin.Context) {
	//账号和密码都不能为空
	username := c.PostForm("username")
	password := c.PostForm("password")
	if password == "" || username == "" {
		utils.ResParamErr(c)
		return
	}
	//
	err, user := serve.SearchUserByName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		utils.ResInterErr(c)
		return
	}
	//
	if user.Username != "" {
		utils.ResDesign(c, 300, "用户已经存在")
		return
	}
	err = serve.AddUser(model.Users{
		Username: username,
		Password: password,
	})
	if err != nil {
		utils.ResInterErr(c)
	}
	utils.ResOk(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if password == "" || username == "" {
		utils.ResParamErr(c)
		return
	}
	err, user := serve.SearchUserByName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		utils.ResInterErr(c)
		return
	}
	if user.Username == "" {
		utils.ResDesign(c, 300, "用户不存在")
		return
	}
	if user.Password != password {
		utils.ResDesign(c, 300, "密码错误")
		return
	}
	uid := strconv.Itoa(user.ID)
	c.SetCookie("name", uid, 3600, "", "/", false, true)
	//
	utils.ResOk(c)

}
func ChangePass(c *gin.Context) {
	newpass := c.PostForm("newpass")
	username := c.PostForm("username")
	_, err := c.Cookie("name")
	if err != nil {
		utils.ResDesign(c, 300, "未设置cookie")
		return
	}
	err = serve.FixPass(username, newpass)
	if err != nil {
		utils.ResInterErr(c)
		return
	}
	utils.ResOk(c)
}
