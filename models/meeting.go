package models

import (
	"time"
)

type Meeting struct {
	Title string
	Originator string
	Participants string
	StartTime	time.Time
	EndTime		time.Time
}