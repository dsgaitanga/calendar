package main

import "time"

type Event struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Start_time  time.Time `json:"start_time"`
	End_time    time.Time `json:"end_time"`
	Status      bool      `json:"status"`
}
