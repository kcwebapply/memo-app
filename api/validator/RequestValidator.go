package requestValidator

import (
	"errors"
	"regexp"
)

var idExp = "^[0-9]+$"

func IdValidator(memo_id string) error {
	var err error

	result := regexp.MustCompile(idExp).Match([]byte(memo_id))
	if !result {
		err = errors.New("invalid memo_id : " + memo_id)
	}
	return err
}
