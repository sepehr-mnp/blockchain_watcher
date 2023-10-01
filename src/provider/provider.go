package provider

import (
	"context"
	"encoding/json"
	"evmbase/src/dtos"
	"evmbase/src/dtos/trongrid"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

type WatcherProvider interface {
	GetLatestBlockNumber() (uint64, error)
	GetBlockByNumber(blockId uint64) ([]byte, error)
}
type tp struct {
	url string
}

type ep struct {
	url string
}

func WatcherFactory(network string) (WatcherProvider, error) {
	if network == "TRON" {
		return &tp{
			url: "https://nile.trongrid.io/walletsolidity/getnowblock",
		}, nil
	}

	// ethereum base

	if network == "POLYGON" {
		return &ep{
			url: "https://rpc-mumbai.maticvigil.com/",
		}, nil
	}

	if network == "BNB" {
		return &ep{}, nil
	}

	return nil, dtos.InvalidNetwork
}

func (t *tp) GetLatestBlockNumber() (uint64, error) {

	url := t.url

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer res.Body.Close()
	var resType trongrid.Block
	json.NewDecoder(res.Body).Decode(&resType)

	return uint64(resType.BlockHeader.RawData.Number), nil

}

func (t *tp) GetBlockByNumber(blockId uint64) ([]byte, error) {
	url := t.url
	strUrl := fmt.Sprintf("%s%s%s", "{\"detail\":true,\"id_or_num\":\"", strconv.FormatUint(blockId, 10), "\"}")

	payload := strings.NewReader(strUrl)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	var resType trongrid.Block

	json.NewDecoder(res.Body).Decode(&resType) ///checkthis
	jsonString, err := json.Marshal(resType.Transactions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return (jsonString), nil
}

func (e *ep) GetLatestBlockNumber() (uint64, error) {

	client, err := ethclient.Dial(e.url)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	numberToUint, err := strconv.ParseUint(header.Number.String(), 10, 64)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return numberToUint, nil

}

func (e *ep) GetBlockByNumber(blockId uint64) ([]byte, error) {

	client, _ := ethclient.Dial(e.url)

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockId)))
	if err != nil {
		log.Fatal(err)
	}
	l := 0
	if block != nil && block.Transactions() != nil {
		l = block.Transactions().Len()
	}
	transactions := make([]string, l)
	for i := 0; i < len(transactions); i++ {
		if block.Transactions()[i].To() != nil { // all this nil checking happend because transaction to create a new contract doesn't have TO_Address
			transactions[i] = block.Transactions()[i].To().Hex()
		}
	}
	return json.Marshal(transactions)
}
