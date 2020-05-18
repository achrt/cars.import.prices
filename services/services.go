package services

import (
	"cars.import.prices/services/cars_catalog"
	"cars.import.prices/domain"
)

type Services struct {
	CarsCatalogService *cars_catalog.Service
}

func (s *Services) CarsCatalog() domain.CarsCatalog {
	return s.CarsCatalogService.WithContext()
}
