package models

import "time"

type Contributions struct {
	Developer  string
	Repository string
	Lines      string
	Date       time.Time
}
