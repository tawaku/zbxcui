package api

import (
	"github.com/pkg/errors"
)

type Error struct {
	Code    rune   `json:"code,string"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (self *Error) Error() error {
	return errors.Errorf("%s(%d): %s", self.Message, self.Code, self.Data)
}
