package api

import (
	"awesomeProject/midware"
	"awesomeProject/serve"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAnswer(c *gin.Context) {
	answer := c.PostForm("answer")
	qid := c.PostForm("qid")
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	err = serve.InsertAnswer(uid, qid, answer)
	if err != nil {
		utils.ResDesign(c, 300, "answer error")
		return
	}
	utils.ResOk(c)
}

func DeleteAnswer(c *gin.Context) {
	aidsting := c.PostForm("aid")
	aid, err := strconv.Atoi(aidsting)
	if err != nil {
		utils.ResParamErr(c)
		return
	}
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	//检验这个是你的
	err, userid := serve.FindUidByQid(aid)
	if err != nil {
		utils.ResDesign(c, 300, "find answer error")
		return
	}
	if uid != userid {
		utils.ResDesign(c, 300, "the answer is not yours")
		return
	}
	//删除
	err = serve.DeleteAnswer(aid)
	if err != nil {
		utils.ResDesign(c, 300, "delete error")
		return
	}
	utils.ResOk(c)
}

func ChangeAnswer(c *gin.Context) {
	answer := c.PostForm("answer")
	aidsting := c.PostForm("aid")
	aid, err := strconv.Atoi(aidsting)
	if err != nil {
		utils.ResParamErr(c)
		return
	}
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	//检验这个是你的
	err, userid := serve.FindUidByQid(aid)
	if err != nil {
		utils.ResDesign(c, 300, "find answer error")
		return
	}
	if uid != userid {
		utils.ResDesign(c, 300, "the answer is not yours")
		return
	}
	//开始修改
	err = serve.ChangeAnswer(aid, answer)
	if err != nil {
		utils.ResDesign(c, 300, "change error")
		return
	}
	utils.ResOk(c)
}

func GetMyAnswer(c *gin.Context) {
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	//获取
	err, answer := serve.FindAnswerByUid(uid)
	if err != nil {
		utils.ResDesign(c, 300, "get answer error")
		return
	}
	c.JSON(http.StatusOK, answer)
}
