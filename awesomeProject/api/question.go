package api

import (
	"awesomeProject/midware"
	"awesomeProject/serve"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateQuestion(c *gin.Context) {
	question := c.PostForm("question")
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "读取cookie错误")
		return
	}
	//添加问题
	err = serve.AddQuestion(uid, question)
	if err != nil {
		utils.ResInterErr(c)
		return
	}
	utils.ResOk(c)
}

func DeleteQuestion(c *gin.Context) {
	qidsting := c.PostForm("qid") //
	qid, err := strconv.Atoi(qidsting)
	if err != nil {
		utils.ResParamErr(c)
		return
	}
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	err, userid := serve.FindUidByQid(qid)
	if err != nil {
		utils.ResDesign(c, 300, "find question error")
		return
	}
	if userid != uid {
		utils.ResDesign(c, 300, "the question is not yours")
		return
	}
	//删除函数
	err = serve.DeleteQuestion(qid)
	if err != nil {
		utils.ResDesign(c, 300, "delete error")
		return
	}
	utils.ResOk(c)
}

func ChangeQuestion(c *gin.Context) {
	question := c.PostForm("question")
	qidsting := c.PostForm("qid")
	qid, err := strconv.Atoi(qidsting)
	if err != nil {
		utils.ResParamErr(c)
		return
	}
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	err, userid := serve.FindUidByQid(qid)
	if err != nil {
		utils.ResDesign(c, 300, "find question error")
		return
	}
	if userid != uid {
		utils.ResDesign(c, 300, "the question is not yours")
		return
	}

	//修改函数
	err = serve.ChangeQuestion(qid, question)
	if err != nil {
		utils.ResDesign(c, 300, "change error")
		return
	}
	utils.ResOk(c)
}

func GetMyQuestion(c *gin.Context) {
	uid, err := midware.Auth(c)
	if err != nil {
		utils.ResDesign(c, 300, "cookie error")
		return
	}
	//根据uid找问题
	err, slices := serve.FindQuestionByUid(uid)
	if err != nil {
		utils.ResDesign(c, 300, "get question error")
		return
	}
	c.JSON(http.StatusOK, slices)
}

func GetAllQuestion(c *gin.Context) {
	err, slices := serve.FindAllQuestion()
	if err != nil {
		utils.ResDesign(c, 300, "get question error")
		return
	}

	c.JSON(http.StatusOK, slices)

}
