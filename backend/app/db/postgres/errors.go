package postgres

import "errors"

var (
	ErrorLangInsertType = errors.New("InsType must be 'users' or 'repos'")
)
