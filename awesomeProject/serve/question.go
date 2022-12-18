package serve

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

func AddQuestion(id int, question string) (err error) {
	err = dao.InsertQuestion(id, question)
	return

}

func DeleteQuestion(qid int) (err error) {
	err = dao.SubQuestion(qid)
	return
}

func ChangeQuestion(qid int, question string) (err error) {
	err = dao.ChangeQuestion(qid, question)
	return
}

func FindUidByQid(qid int) (err error, userid int) {
	err, userid = dao.FindUidByQid(qid)
	return
}

//

func FindQuestionByUid(uid int) (err error, slices []model.Question) {
	err, slices = dao.FindQuestionByUid(uid)
	return
}
func FindAllQuestion() (err error, slices []model.Question) {
	err, slices = dao.FindAllQuestion()
	return
}
