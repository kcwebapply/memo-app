package memoservice

import (
	. "github.com/kcwebapply/examination/infrastructure/model"
	repository "github.com/kcwebapply/examination/infrastructure/repository"
)

func GetMemo(memo_id string) Memo {
	return repository.GetOneMemo(memo_id)
}

func postMemo(memo Memo) bool {
	return repository.SaveMemo(memo)
}
