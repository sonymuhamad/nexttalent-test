package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/sonymuhamad/nexttalent-test"
	"github.com/sonymuhamad/nexttalent-test/model"
)

type Person struct {
	BaseRepository
}

func NewPerson(db *gorm.DB) *Person {
	return &Person{
		BaseRepository{
			db: db,
		},
	}
}

func (p *Person) GetByName(ctx context.Context, name string) (model.Person, error) {
	db := p.getDB(ctx)

	var person model.Person

	err := db.Where("name = ?", name).
		First(&person).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return person, nexttalent.ErrPersonNotFound
		}

		return person, err
	}

	return person, nil
}
