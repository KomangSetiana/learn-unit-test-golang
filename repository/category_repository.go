package repository

import "learn-unit-test-golang/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
