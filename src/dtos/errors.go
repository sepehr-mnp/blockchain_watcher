package dtos

import "errors"

var (
	AddressUnderWatcherError = errors.New("Address under watcher error!")
	GetAddressPaymentIDError = errors.New("Get address payment id error!")
	CreateTransactionError   = errors.New("Create transaction error!")

	FailedGetBlockRequestError = errors.New("Failed get block request error!")
	FailedTXInfoRequestError   = errors.New("Failed get transaction info request error!")

	FailedPaymentDepositeRequestError = errors.New("Failed payment deposite request error!")

	TxNotFoundError           = errors.New("Transaction not found!")
	TransactionIsDuplicateErr = errors.New("Transaction is duplicate!")
)
