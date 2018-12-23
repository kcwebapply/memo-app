package memoRepository

import (
	"database/sql"
	"fmt"
	"math/rand"

	. "github.com/kcwebapply/examination/infrastructure/model"
	_ "github.com/lib/pq"
)

var user = "wadakeishi"
var password = "qazwsx12e"
var host = "localhost"
var db_name = "memo"

var Db *sql.DB

func init() {
	rand.Float32()
	var err error
	Db, err = sql.Open("postgres", "postgres://"+user+":"+password+"@"+host+"/"+db_name+"?sslmode=disable")
	if err != nil {
		fmt.Println("error connection:", err)
		panic(err)
	}
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

func SaveMemo(memo Memo) bool {
	var err error
	_, err = Db.Exec("insert into memo values($1,$2,$3,$4,$5)", memo.Id, memo.Title, memo.Text, memo.Flag, memo.Date)
	if err != nil {
		fmt.Print("error:", err)
	}
	return err == nil
}

/*func GetOneMemo(memo_id string) Memo{
  memo := Memo{}
  var err error
  err = Db.QueryRow("select id,title,text,flag,date from memo where id = $1", memo_id).Scan(&memo.Id, &memo.Title, &memo.Text,&memo.Flag,&memo.Date)
  if(err != nil){
    fmt.Print("error:",err)
  }
  return memo
}*/
