package repository

import (
	"evmbase/src/repository/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type Address interface {
	CreateInvoiceAddress(*models.InvoiceAddress) (*models.InvoiceAddress, error)
	CreateUserAddress(*models.UserAddress) (*models.UserAddress, error)

	GetInvoiceAddress(invoiceID uint32, network string, seedVersion string) (*models.InvoiceAddress, error)
	GetUserAddress(userID uint32, network string, seedVersion string) (*models.UserAddress, error)

	UpdateInvoiceAddress(*models.InvoiceAddress) (*models.InvoiceAddress, error)
	Migrate()
}

type address struct {
	DB *gorm.DB
}

func NewHDWalletRepository(db *gorm.DB) Address {
	repo := address{DB: db}
	return &repo
}

func (repo *address) GetUserAddress(userID uint32, network string, seedVersion string) (*models.UserAddress, error) {
	address := models.UserAddress{}
	err := repo.DB.Model(models.UserAddress{}).
		Where("user_id = ? AND network = ? AND seed_version = ?", userID, network, seedVersion).
		First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repo *address) GetInvoiceAddress(invoiceID uint32, network string, seedVersion string) (*models.InvoiceAddress, error) {
	address := models.InvoiceAddress{}
	err := repo.DB.Model(models.InvoiceAddress{}).
		Where("invoice_id = ? AND network = ? AND seed_version = ?", invoiceID, network, seedVersion).
		First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repo *address) UpdateInvoiceAddress(address *models.InvoiceAddress) (*models.InvoiceAddress, error) {
	err := repo.DB.Save(address).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return address, nil
}

func (repo *address) CreateInvoiceAddress(data *models.InvoiceAddress) (*models.InvoiceAddress, error) {
	err := repo.DB.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (repo *address) CreateUserAddress(data *models.UserAddress) (*models.UserAddress, error) {
	err := repo.DB.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *address) Migrate() {
	err := repo.DB.AutoMigrate(&models.InvoiceAddress{})
	if err != nil {
		log.Fatal(err)
	}
	err = repo.DB.AutoMigrate(&models.UserAddress{})
	if err != nil {
		log.Fatal(err)
	}
}
