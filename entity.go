package main

import "time"

type message struct {
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Name          string    `json:"name"`
	Body          string    `json:"body"`
	WrittenRegion string    `json:"written_region"`
	ServerRegion  string    `json:"server_region"`
}
