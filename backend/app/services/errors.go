package services

import "errors"

var (
	ErrorDataNeedFetch  = errors.New("developer need fetch")
	ErrorDataProcessing = errors.New("developer processing")
)
