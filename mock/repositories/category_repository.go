package repositories

import "golang-unitTest/mock/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}
