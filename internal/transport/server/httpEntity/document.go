package httpEntity

import "time"

type DocumentSelf struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
}
