package httpEntity

import (
	"time"
)

type ApplicationGet struct {
	Id          string `json:"id"`
	ContactInfo string `json:"contact_info"`
	Name        string `json:"name"`
	Text        string `json:"text"`
	// Status must be "Создан" or "Закрыт"
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Link      string    `json:"link"`
}

type ApplicationWithInfo struct {
	Id          string `json:"id"`
	ContactInfo string `json:"contact_info"`
	Name        string `json:"name"`
	Text        string `json:"text"`
	// Status must be "Создан" or "Закрыт"
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Link      string    `json:"link"`
	FIO       string    `json:"fio"`
	Course    int       `json:"course"`
	Group     int       `json:"group"`
}
