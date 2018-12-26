package memoservice

import (
	"time"

	. "github.com/kcwebapply/memo-app/api/resource"
	. "github.com/kcwebapply/memo-app/infrastructure/model"
	repository "github.com/kcwebapply/memo-app/infrastructure/repository"
)

func GetMemo(memo_id string) Memo {
	return repository.GetOneMemo(memo_id)
}

func GetAllMemo() []Memo {
	var memos, _ = repository.GetAllEffectiveMemo()
	return memos
}

func PostMemo(memoRequest MemoRequest) bool {
	var newId string = repository.GetNewId()
	memoObject := Memo{Id: newId, Title: memoRequest.Title, Text: memoRequest.Text, Flag: false, Date: timeGenerator()}
	return repository.SaveMemo(memoObject)
}

func DeleteMemo(memo_id string) {
	repository.DeleteMemo(memo_id)
}

func timeGenerator() string {
	return time.Now().Format("2006-01-02")
}
