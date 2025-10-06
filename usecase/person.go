package usecase

import (
	"context"

	"github.com/sonymuhamad/nexttalent-test/repository"
)

type Person struct {
	personRepository *repository.Person
}

func NewPerson(personRepository *repository.Person) *Person {
	return &Person{
		personRepository: personRepository,
	}
}

func (p *Person) GetCountryByPersonName(ctx context.Context, name string) (string, error) {
	person, err := p.personRepository.GetByName(ctx, name)
	if err != nil {
		return "", err
	}

	return person.Country, nil
}
