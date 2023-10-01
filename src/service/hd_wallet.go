package service

import (
	"errors"
	"fmt"

	"evmbase/src/dtos"
	"evmbase/src/repository"
	"evmbase/src/repository/models"

	gotron "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/foxnut/go-hdwallet"

	"gorm.io/gorm"
)

type IHDWallet interface {
	GenerateAddress(addressIndex uint32) (string, string, string)
	GetKeys(addressIndex uint32) (string, string, string)
}

type tron struct {
	masterKey *hdwallet.Key
}

func (net *tron) GenerateAddress(addressIndex uint32) (string, string, string) {
	wallet, _ := net.masterKey.GetWallet(hdwallet.AddressIndex(addressIndex))
	publicECDSA := wallet.GetKey().PublicECDSA
	address := fmt.Sprint(gotron.PubkeyToAddress(*publicECDSA))
	privateHex := wallet.GetKey().PrivateHex()
	publicHex := wallet.GetKey().PublicHex(true)
	return address, privateHex, publicHex
}

func (net *tron) GetKeys(addressIndex uint32) (string, string, string) {
	wallet, _ := net.masterKey.GetWallet(hdwallet.AddressIndex(addressIndex))
	publicECDSA := wallet.GetKey().PublicECDSA
	address := fmt.Sprint(gotron.PubkeyToAddress(*publicECDSA))
	privateHex := wallet.GetKey().PrivateHex()
	publicHex := wallet.GetKey().PublicHex(true)
	return address, privateHex, publicHex
}

type ethereum struct {
	masterKey *hdwallet.Key
}

func (net *ethereum) GenerateAddress(addressIndex uint32) (string, string, string) {
	wallet, _ := net.masterKey.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(addressIndex))
	address, _ := wallet.GetAddress()

	privateHex := wallet.GetKey().PrivateHex()
	publicHex := wallet.GetKey().PublicHex(true)

	return address, privateHex, publicHex
}

func (net *ethereum) GetKeys(addressIndex uint32) (string, string, string) {
	wallet, _ := net.masterKey.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(addressIndex))
	address, _ := wallet.GetAddress()

	privateHex := wallet.GetKey().PrivateHex()
	publicHex := wallet.GetKey().PublicHex(true)

	return address, privateHex, publicHex
}

func HDWalletFactory(network string, masterKey *hdwallet.Key) (IHDWallet, error) {
	if network == "TRON" {
		return &tron{
			masterKey: masterKey,
		}, nil
	}

	// ethereum base
	if network == "ETHEREUM" {
		return &ethereum{
			masterKey: masterKey,
		}, nil
	}

	if network == "POLYGON" {
		return &ethereum{
			masterKey: masterKey,
		}, nil
	}

	if network == "BNB" {
		return &ethereum{
			masterKey: masterKey,
		}, nil
	}

	return nil, dtos.InvalidNetwork
}

type HDWallet interface {
	GenerateInvoiceAddress(invoiceID uint32, network string) (string, error)
	GenerateUserAddress(userID uint32, network string) (string, error)
	GetInvoiceKeys(invoiceID uint32, network string) (string, string, string, error)
	GetUserKeys(userID uint32, network string) (string, string, string, error)
}

type hdWallet struct {
	repo               repository.Address
	invoiceMasterKey   *hdwallet.Key
	userMasterKey      *hdwallet.Key
	invoiceSeedVersion string
	userSeedVersion    string
}

func NewHDWalletServic(repo repository.Address, invoiceMasterKey *hdwallet.Key, userMasterKey *hdwallet.Key, invoiceSeedVersion string, userSeedVersion string) HDWallet {
	svc := hdWallet{
		repo:               repo,
		invoiceMasterKey:   invoiceMasterKey,
		userMasterKey:      userMasterKey,
		userSeedVersion:    userSeedVersion,
		invoiceSeedVersion: invoiceSeedVersion,
	}
	return &svc
}

func (svc *hdWallet) GenerateInvoiceAddress(InvoiceID uint32, network string) (string, error) {
	addressDB, err := svc.repo.GetInvoiceAddress(InvoiceID, network, svc.invoiceSeedVersion)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		hdWallet, err := HDWalletFactory(network, svc.invoiceMasterKey)
		if err != nil {
			return "", err
		}

		addressHex, _, _ := hdWallet.GenerateAddress(InvoiceID)
		newAddress := models.InvoiceAddress{
			InvoiceID:   InvoiceID,
			Network:     network,
			Address:     addressHex,
			SeedVersion: svc.invoiceSeedVersion,
		}
		_, err = svc.repo.CreateInvoiceAddress(&newAddress)
		if err != nil {
			return "", err
		}

		return newAddress.Address, nil
	}

	return addressDB.Address, nil
}

func (svc *hdWallet) GenerateUserAddress(userID uint32, network string) (string, error) {
	addressDB, err := svc.repo.GetUserAddress(userID, network, svc.userSeedVersion)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		hdWallet, err := HDWalletFactory(network, svc.userMasterKey)
		if err != nil {
			return "", err
		}

		addressHex, _, _ := hdWallet.GenerateAddress(userID)
		newAddress := models.UserAddress{
			UserID:      userID,
			Network:     network,
			Address:     addressHex,
			SeedVersion: svc.userSeedVersion,
		}
		_, err = svc.repo.CreateUserAddress(&newAddress)
		if err != nil {
			return "", err
		}

		return newAddress.Address, nil
	}

	return addressDB.Address, nil
}

func (svc *hdWallet) GetInvoiceKeys(invoiceID uint32, network string) (string, string, string, error) {
	hdWallet, err := HDWalletFactory(network, svc.invoiceMasterKey)
	if err != nil {
		return "", "", "", err
	}

	addressHex, privateHex, publicHex := hdWallet.GenerateAddress(invoiceID)
	return addressHex, privateHex, publicHex, nil
}

func (svc *hdWallet) GetUserKeys(userID uint32, network string) (string, string, string, error) {
	hdWallet, err := HDWalletFactory(network, svc.userMasterKey)
	if err != nil {
		return "", "", "", err
	}

	addressHex, privateHex, publicHex := hdWallet.GenerateAddress(userID)
	return addressHex, privateHex, publicHex, nil
}
