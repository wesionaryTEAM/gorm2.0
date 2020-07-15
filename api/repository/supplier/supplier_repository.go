package repository

import (
	"gorm2.0/model"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	DB *gorm.DB
}

//Constructor function
func NewSupplierRepository (db *gorm.DB) model.SupplierRepository {
	return &SupplierRepository {
		DB: db,
	}
}

func (s *SupplierRepository) Save(supplier *model.Supplier) (*model.Supplier, error) {
	return supplier, s.DB.Create(&supplier).Error
}

func (s *SupplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := s.DB.Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *SupplierRepository) Delete(supplier *model.Supplier) error {
	return s.DB.Delete(&supplier).Error
}

func (s *SupplierRepository) Migrate() error {
	err := s.DB.AutoMigrate(&model.Supplier{})
	return err
}

