package model

import "errors"

var ErrPartNotFound = errors.New("part not found")

type NotFound struct {
	// HTTP-код ошибки.
	Code int `json:"code"`
	// Описание ошибки.
	Message string `json:"message"`
}

func (*NotFound) getOrderByIdRes() {}
