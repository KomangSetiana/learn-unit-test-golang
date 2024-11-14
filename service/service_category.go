package service

import (
	"errors"
	"learn-unit-test-golang/entity"
	"learn-unit-test-golang/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (s CategoryService) Get(id string) (*entity.Category, error) {
	category := s.Repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}

}
