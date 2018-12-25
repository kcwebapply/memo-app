package memoRepository

import (
	"fmt"
	"strconv"

	"github.com/gocraft/dbr"
	. "github.com/kcwebapply/memo-app/infrastructure/config"
	. "github.com/kcwebapply/memo-app/infrastructure/model"
	_ "github.com/lib/pq"
)

var Db *dbr.Connection

func init() {
	Db = GetConnection()
	Db.SetMaxOpenConns(10)
}

func GetOneMemo(memo_id string) Memo {
	sess := Db.NewSession(nil)
	var memo Memo
	var err error
	_, err = sess.Select("id,title,text,flag,date").From("memo").Where("id = ?", memo_id).Load(&memo)
	if err != nil {
		fmt.Print("error:", err)
	}
	return memo
}

func GetAllEffectiveMemo() (memos []Memo, err error) {
	sess := Db.NewSession(nil)
	var rows []Memo
	_, err = sess.Select("*").From("memo").Where("flag = ?", false).Load(&rows)
	if err != nil {
		fmt.Print(err)
	}
	return rows, err
}

func GetNewId() string {
	var newId int
	var err error
	sess := Db.NewSession(nil)
	sess.Select("id").From("memo").OrderDesc("id").Limit(1).Load(&newId)
	newId++
	if err != nil {
		fmt.Print("error:", err)
	}
	return strconv.Itoa(newId)
}

func SaveMemo(memo Memo) bool {
	var err error
	fmt.Print("title", memo.Title, "text", memo.Text)
	sess := Db.NewSession(nil)
	sess.InsertInto("memo").Columns("id", "title", "text", "flag", "date").Record(memo).Exec()
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}

func DeleteMemo(memo_id string) bool {
	var err error
	sess := Db.NewSession(nil)
	sess.Update("memo").Set("flag", true).Where("id = ?", memo_id).Exec()
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}
