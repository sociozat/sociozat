package models

import "time"

type UserTodaysChannels struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UserHeaderChannels struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	LastRead time.Time `json:"lastRead"`
}
