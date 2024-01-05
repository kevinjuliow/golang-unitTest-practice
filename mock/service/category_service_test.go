package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unitTest/mock/repositories"
	"testing"
)

var categoryRepository = repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repositories: &categoryRepository}

func TestCategoryServiceGet(t *testing.T) {
	//Mock
	//When called the function FindById with "1" as the argument , expects to return nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.GET("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
