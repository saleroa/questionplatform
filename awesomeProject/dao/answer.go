package dao

import (
	"awesomeProject/model"
	"database/sql"
	"log"
)

func InsertAnswer(uid int, qid string, answer string) (err error) {
	_, err = DB.Exec("insert into answer (uid,qid,answer) values (?,?,?)", uid, qid, answer)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func FindUidByAid(aid int) (err error, userid int) {
	row := DB.QueryRow("SELECT uid from question where aid = ?", aid)
	if err = row.Err(); err != nil {
		log.Println(err)
		return
	}
	err = row.Scan("&userid")
	return
}
func DeleteAnswer(aid int) (err error) {
	_, err = DB.Exec("delete from answer where aid = ?", aid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func ChangeAnswer(aid int, answer string) (err error) {
	_, err = DB.Exec("update answer set answer = ? where  aid =?", answer, aid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func FindAnswerByUid(uid int) (err error, slices []model.Answer) {
	rows, err := DB.Query("select qid ,answer from answer where uid=?", uid)
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
	//
	var slice model.Answer

	for rows.Next() {
		err = rows.Scan(&slice.Answer)
		if err != nil {
			log.Println(err)
			return
		}
		slices = append(slices, slice)
	}
	return

	return
}
