package repository

import (
	"back-usm/internals/product/core/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductRepository struct {
	uri   string
	db    *gorm.DB
	table string
}

func NewProductRepository(uri string) (*ProductRepository, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &ProductRepository{
		uri:   uri,
		db:    db,
		table: "products",
	}, nil
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Table(r.table).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetOne(id string) (domain.Product, error) {
	var product domain.Product
	err := r.db.Table(r.table).Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	err := r.db.Table(r.table).Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Update(id string, product domain.Product) (domain.Product, error) {
	err := r.db.Table(r.table).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id string) error {
	err := r.db.Table(r.table).Where("id = ?", id).Delete(&domain.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}
