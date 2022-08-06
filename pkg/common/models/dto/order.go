package dto

import (
	"time"
)

type Order struct {
	Id            uint
	Number        string
	Date          *time.Time
	CustomerRefer uint
	Comment       string
}
