package service

import (
	"errors"
	"golang-unitTest/mock/entities"
	"golang-unitTest/mock/repositories"
)

type CategoryService struct {
	Repositories repositories.CategoryRepository
}

// create a function thats in a CategoryService named GET with a string parameter that will return entitiesCategory and error
func (s *CategoryService) GET(id string) (*entities.Category, error) {
	foundedCategory := s.Repositories.FindById(id)
	if foundedCategory == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return foundedCategory, nil
	}
}
