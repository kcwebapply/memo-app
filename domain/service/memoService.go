package memoservice

import (
	"time"

	. "github.com/kcwebapply/examination/api/resource"
	. "github.com/kcwebapply/examination/infrastructure/model"
	repository "github.com/kcwebapply/examination/infrastructure/repository"
)

func GetMemo(memo_id string) Memo {
	return repository.GetOneMemo(memo_id)
}

func PostMemo(memoRequest MemoRequest) bool {
	var newId = repository.GetNewId()
	memoObject := Memo{Id: newId, Title: memoRequest.Title, Text: memoRequest.Text, Flag: false, Date: timeGenerator()}
	return repository.SaveMemo(memoObject)
}

func timeGenerator() string {
	return time.Now().Format("2006-01-02")
}
