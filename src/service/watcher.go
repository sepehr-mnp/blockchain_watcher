package service

/*
* this version of code doesn't use task, after rabbitMQ is set up, task will be written
 */

import (
	"encoding/json"
	"errors"
	"evmbase/src/dtos/trongrid"
	"evmbase/src/provider"
	"evmbase/src/repository"
	"evmbase/src/utils"
	"fmt"
	"log"
)

type BlockForFactory struct {
	transactionIds []uint64
}
type TransactionForFactory struct {
	from   string
	amount uint64
	to     string
}

type IWatcher interface {
	FetchBlockById(uint64) (BlockForFactory, error)
	FetchTransactionById(uint64) (TransactionForFactory, error)
	ProcesseUnProcessedBlocks() error
	GetProvider() provider.WatcherProvider
}

type tronWatcher struct {
	redisRepo repository.RedisRepo
	Provider  provider.WatcherProvider
	contracts map[string]string
}

func IWatcherFactory(network string, redisClient repository.RedisRepo) (IWatcher, error) {
	if network == "TRON" {
		provider, err := provider.WatcherFactory("TRON")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &tronWatcher{
			redisRepo: redisClient,
			Provider:  provider,
		}, nil
	} else if network == "POLYGON" {
		provider, err := provider.WatcherFactory("POLYGON")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &ethereumWatcher{
			redisRepo: redisClient,
			Provider:  provider,
		}, nil
	}
	return nil, errors.New("Invalid Network!")
}

/*
* nothing returns yet because there is no Task
* after creating task.go, the functions in here will be modified so the task can get the transactions that we are paid by and alert the core with rabbitMQ
 */
func (tw *tronWatcher) processTransaction(transaction trongrid.Transaction) { ///nothing returns yet because there is no Task
	if transaction.Ret[0].ContractRet == "SUCCESS" {
		data := transaction.RawData.Contract[0]
		txDetail := data.Parameter.Value
		// TOKEN transfer
		if data.Type == "TriggerSmartContract" {
			contractAddress := txDetail.ContractAddress
			// Checking this contract address supported by us
			_, ok := tw.contracts[contractAddress]
			if ok {
				if txDetail.Data[:8] == "a9059cbb" {
					dataToAddress := txDetail.Data[8:72]
					toAddress := utils.ToBase58CheckAddress("41" + dataToAddress[len(dataToAddress)-40:])
					under, err := tw.redisRepo.AddressUnderWatcher(toAddress)
					if err != nil {
						log.Fatal(err)
					}
					if under {
						///check the transaction and give it to task

					}
				} else if txDetail.Data[:8] == "23b872dd" {
					dataToAddress := txDetail.Data[72:136]
					toAddress := utils.ToBase58CheckAddress("41" + dataToAddress[len(dataToAddress)-40:])
					under, err := tw.redisRepo.AddressUnderWatcher(toAddress)
					if err != nil {
						log.Fatal(err)
					}
					if under {
						///check the transaction and give it to task
					}
				}
			}
		}
		if data.Type == "TransferContract" { // TRX token
			toAddress := utils.ToBase58CheckAddress(txDetail.ToAddress)
			under, err := tw.redisRepo.AddressUnderWatcher(toAddress)
			if err != nil {
				log.Fatal(err)
			}
			if under {
				fmt.Println("address ", toAddress, " got paid")
				///check the transaction and give it to task
			}
		}
	}
}

func (tw *tronWatcher) proccessBlock(blockNumber uint64) error {
	jsonString, err := tw.Provider.GetBlockByNumber(blockNumber)
	if err != nil {
		log.Fatal(err)
		return err
	}
	transactions := []trongrid.Transaction{}
	err = json.Unmarshal((jsonString), &transactions)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for i := 0; i < len(transactions); i++ {
		go tw.processTransaction(transactions[i])
	}
	return nil
}

func (tw *tronWatcher) ProcesseUnProcessedBlocks() error {
	redisLastBlock := tw.redisRepo.GetLatestBlock()
	fmt.Println("started from:", redisLastBlock)
	networkLastBlock, err := tw.Provider.GetLatestBlockNumber()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("lasst network block: ", networkLastBlock)
	for ; redisLastBlock < networkLastBlock; redisLastBlock++ {
		go tw.proccessBlock(redisLastBlock)
	}
	tw.redisRepo.SetLatestBlock(networkLastBlock)
	return nil
}

func (tw *tronWatcher) GetProvider() provider.WatcherProvider {
	return tw.Provider
}

func (tw *tronWatcher) FetchBlockById(blockId uint64) (BlockForFactory, error) {
	return BlockForFactory{}, nil
}

func (tw *tronWatcher) FetchTransactionById(transactionId uint64) (TransactionForFactory, error) {
	return TransactionForFactory{}, nil
}

type ethereumWatcher struct {
	redisRepo repository.RedisRepo
	Provider  provider.WatcherProvider
}

func (eth *ethereumWatcher) processTransaction(toAddress string) {
	under, _ := eth.redisRepo.AddressUnderWatcher(toAddress)
	if under {
		fmt.Println(toAddress)
		////after we set up rabbit mq, we should modify here to get more data and set it through rabbitMQ
	}
}
func (eth *ethereumWatcher) proccessBlock(blockNumber uint64) error {

	jsonString, err := eth.Provider.GetBlockByNumber(blockNumber)

	if err != nil {
		log.Fatal(err)
		return err
	}

	transactions := []string{}
	json.Unmarshal(jsonString, &transactions)

	for i := 0; i < len(transactions); i++ {
		go eth.processTransaction(transactions[i])
	}
	return nil
}

func (eth *ethereumWatcher) ProcesseUnProcessedBlocks() error {
	redisLastBlock := eth.redisRepo.GetLatestBlock()
	fmt.Println("started from:", redisLastBlock)
	networkLastBlock, err := eth.Provider.GetLatestBlockNumber()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("last network block: ", networkLastBlock)
	for ; redisLastBlock < networkLastBlock; redisLastBlock++ {
		go eth.proccessBlock(redisLastBlock)
	}
	eth.redisRepo.SetLatestBlock(networkLastBlock)
	return nil
}

func (eth *ethereumWatcher) GetProvider() provider.WatcherProvider {
	return eth.Provider
}

func (eth *ethereumWatcher) FetchBlockById(blockId uint64) (BlockForFactory, error) {
	return BlockForFactory{}, nil
}

func (eth *ethereumWatcher) FetchTransactionById(transactionId uint64) (TransactionForFactory, error) {
	return TransactionForFactory{}, nil
}
