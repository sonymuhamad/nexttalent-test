package model

import "time"

type Person struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Person) TableName() string {
	return "persons"
}
