package service 

import (
	"sync"

	"gorm2.0/model"
)

var once sync.Once

type supplierService struct {
	supplierRepository model.SupplierRepository
}

var instance *supplierService

//Constructor function
func NewSupplierService(repository model.SupplierRepository) model.SupplierService {
	once.Do(func(){
		instance = &supplierService{
			supplierRepository: repository,
		}
	})
	return instance
}

func (s *supplierService) Create(supplier *model.Supplier) (*model.Supplier, error) {
	return s.supplierRepository.Save(supplier)
}

func (s *supplierService) FindAll() ([]model.Supplier, error) {
	return s.supplierRepository.FindAll()
}

func (s *supplierService) Delete(supplier *model.Supplier) error {
	return s.supplierRepository.Delete(supplier)
}

