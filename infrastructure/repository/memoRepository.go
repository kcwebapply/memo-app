package memoRepository

import (
	"database/sql"
	"fmt"
	"strconv"

	. "github.com/kcwebapply/memo-app/infrastructure/config"
	. "github.com/kcwebapply/memo-app/infrastructure/model"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	Db = GetConnection()
}

func GetOneMemo(memo_id string) Memo {
	memo := Memo{}
	var err error
	err = Db.QueryRow("select id,title,text,flag,date from memo where id = $1", memo_id).Scan(&memo.Id, &memo.Title, &memo.Text, &memo.Flag, &memo.Date)
	if err != nil {
		fmt.Print("error:", err)
	}
	return memo
}

func GetAllEffectiveMemo() (memos []Memo, err error) {
	rows, err := Db.Query("select id,title,text,flag,date from memo where flag=false")
	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		memo := Memo{}
		err = rows.Scan(&memo.Id, &memo.Title, &memo.Text, &memo.Flag, &memo.Date)
		if err != nil {
			fmt.Print(err)
		}
		memos = append(memos, memo)
	}
	rows.Close()
	return memos, err
}

func GetNewId() string {
	var newId int
	var err error
	err = Db.QueryRow("select id from memo order by id desc limit 1").Scan(&newId)
	newId++
	if err != nil {
		fmt.Print("error:", err)
	}
	return strconv.Itoa(newId)
}

func SaveMemo(memo Memo) bool {
	var err error
	_, err = Db.Exec("insert into memo values($1,$2,$3,$4,$5)", memo.Id, memo.Title, memo.Text, memo.Flag, memo.Date)
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}

func DeleteMemo(memo_id string) bool {
	var err error
	_, err = Db.Exec("update memo set flag = true where id = $1", memo_id)
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}
