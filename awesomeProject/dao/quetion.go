package dao

import (
	"awesomeProject/model"
	"database/sql"
	"log"
)

func InsertQuestion(uid int, question string) (err error) {
	_, err = DB.Exec("insert into question (uid,question) values (?,?) ", uid, question)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func SubQuestion(qid int) (err error) {
	_, err = DB.Exec("delete from question where qid = ?", qid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func ChangeQuestion(qid int, question string) (err error) {
	_, err = DB.Exec("update question set question = ? where qid = ?", question, qid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// 根据qid查询uid
func FindUidByQid(qid int) (err error, userid int) {
	row := DB.QueryRow("select uid from question where qid = ?", qid)
	if err = row.Err(); err != nil {
		log.Println(err)
		return
	}
	err = row.Scan(&userid) //当这个只有一项的时候就不打引号
	if err != nil {
		log.Println(err)
	}
	return
}

// 三个get函数很可能有问题

func FindQuestionByUid(uid int) (err error, slices []model.Question) {
	//question model.Question
	rows, err := DB.Query("select qid ,question from question where uid =? ", uid)
	if err != nil {
		log.Println(err)
		return
	}
	//
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var slice model.Question

	for rows.Next() {
		err = rows.Scan(&slice.Qid, &slice.Question)
		if err != nil {
			log.Println(err)
			return
		}
		slices = append(slices, slice)
	}
	return
}

func FindAllQuestion() (err error, slices []model.Question) {
	rows, err := DB.Query("select qid ,question from question  ")
	if err != nil {
		log.Println(err)
		return
	}
	//
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var slice model.Question

	for rows.Next() {
		err = rows.Scan(&slice.Qid, &slice.Question)
		if err != nil {
			log.Println(err)
			return
		}
		slices = append(slices, slice)
	}
	return
}
