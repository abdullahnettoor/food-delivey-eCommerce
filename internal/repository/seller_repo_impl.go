package repository

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(DB *gorm.DB) interfaces.ISellerRepository {
	return &SellerRepository{DB: DB}
}

func (repo *SellerRepository) FindAll() (*[]entities.Seller, error) {
	var sellerList []entities.Seller

	if err := repo.DB.Raw(`
	SELECT * 
	FROM sellers
	WHERE status <> 'Deleted'`).
		Scan(&sellerList).Error; err != nil {
		return nil, err
	}

	return &sellerList, nil
}

func (repo *SellerRepository) FindByID(id string) (*entities.Seller, error) {
	var seller entities.Seller

	res := repo.DB.Raw(`
	SELECT *
	FROM sellers
	WHERE id = ?`, id).
		Scan(&seller)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &seller, nil
}

func (repo *SellerRepository) FindByEmail(email string) (*entities.Seller, error) {
	var seller entities.Seller

	res := repo.DB.Raw(`
	SELECT *
	FROM sellers
	WHERE email = ?`, email).
		Scan(&seller)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &seller, nil
}

func (repo *SellerRepository) Create(seller *entities.Seller) error {
	query := repo.DB.Raw(`
	SELECT *
	FROM sellers 
	WHERE email = ?`, seller.Email)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected > 0 {
		return e.ErrConflict
	}
	if err := repo.DB.Create(&seller).Error; err != nil {
		return err
	}
	return nil
}
