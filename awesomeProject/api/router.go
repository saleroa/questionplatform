package api

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	u := r.Group("/user")
	// group的花括号必须放在下一行
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.PUT("/changepass", ChangePass)
	}

	q := r.Group("/question")
	{
		q.POST("/create", CreateQuestion)
		q.PUT("/change", ChangeQuestion)
		q.DELETE("/delete", DeleteQuestion)
		q.GET("/getmine", GetMyQuestion)
		q.GET("/getall", GetAllQuestion)
	}

	a := r.Group("/answer")
	{
		a.POST("/create", CreateAnswer)
		a.DELETE("/delete", DeleteAnswer)
		a.PUT("/change", ChangeAnswer)
		a.GET("/getmine", GetMyAnswer)
	}

	_ = r.Run()
}
