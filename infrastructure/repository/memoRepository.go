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
var Sess *dbr.Session

func init() {
	Db = GetConnection()
	Db.SetMaxOpenConns(10)
	Sess = Db.NewSession(nil)
}

func GetOneMemo(memo_id string) Memo {
	var memo Memo
	var err error
	_, err = Sess.Select("id,title,text,flag,date").From("memo").Where("id = ?", memo_id).Load(&memo)
	if err != nil {
		fmt.Print("error:", err)
	}
	return memo
}

func GetAllEffectiveMemo() (memos []Memo, err error) {
	var rows []Memo
	_, err = Sess.Select("*").From("memo").Where("flag = ?", false).Load(&rows)
	if err != nil {
		fmt.Print(err)
	}
	return rows, err
}

func GetNewId() string {
	var newId int
	var err error
	Sess.Select("id").From("memo").OrderDesc("id").Limit(1).Load(&newId)
	newId++
	if err != nil {
		fmt.Print("error:", err)
	}
	return strconv.Itoa(newId)
}

func SaveMemo(memo Memo) bool {
	var err error
	fmt.Print("title", memo.Title, "text", memo.Text)
	Sess.InsertInto("memo").Columns("id", "title", "text", "flag", "date").Record(memo).Exec()
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}

func DeleteMemo(memo_id string) bool {
	var err error
	Sess.Update("memo").Set("flag", true).Where("id = ?", memo_id).Exec()
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}
