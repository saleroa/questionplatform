package serve

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

func InsertAnswer(uid int, qid string, answer string) (err error) {
	err = dao.InsertAnswer(uid, qid, answer)
	return
}
func FindUidByAid(aid int) (err error, userid int) {
	err, userid = dao.FindUidByAid(aid)
	return
}

func DeleteAnswer(aid int) (err error) {
	err = dao.DeleteAnswer(aid)
	return
}
func ChangeAnswer(aid int, answer string) (err error) {
	err = dao.ChangeAnswer(aid, answer)
	return
}
func FindAnswerByUid(uid int) (err error, slices []model.Answer) {
	err, slices = dao.FindAnswerByUid(uid)
	return
}
