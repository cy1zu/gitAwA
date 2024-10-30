package postgres

import "errors"

var (
	ErrorLangInsertType     = errors.New("insType must be 'users' or 'repos'")
	ErrorDeveloperNotStored = errors.New("developer not stored")
	ErrorInitDatabaseFailed = errors.New("init postgres failed")
)
