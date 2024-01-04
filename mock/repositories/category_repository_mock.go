package repositories

import (
	"github.com/stretchr/testify/mock"
	"golang-unitTest/mock/entities"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entities.Category {
	arguments := repository.Mock.Called(id)
	//Get(0) because the FindById Function only returns 1
	if arguments.Get(0) == nil {
		return nil
	} else {
		//parse the category into entities.Category
		category := arguments.Get(0).(entities.Category)
		//because return type is pointer , so need to use &
		return &category
	}
}
