package main

import (
	"context"
	"evmbase/src/core"
	"evmbase/src/repository"
	"evmbase/src/service"
	"strconv"
	"time"
)

func main() {
	// "5bb93439c874832ab3bb8923d20e2e824954716dd8d5e2fbcdeeef477aa8c96888124173f8559fa69485709af73c7236ce620f94a4294b9210dad3e76e83d417"
	// panic: seed length must be between 128 and 512 bits
	// master, err := hdwallet.NewKey(
	// 	hdwallet.Seed([]byte("6a2f85cb3a5ef0a7d7747bbfe1e439f0")),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// openDB := core.OpenPostgresql("postgresql://root:FndO360ldzRNYDgltzQqfkFE@may.iran.liara.ir:32662/postgres")
	// repo := repository.NewHDWalletRepository(openDB)
	// repo.Migrate()

	// svc := service.NewHDWalletServic(repo, master, master, "1", "1")

	// addr, err := svc.GenerateInvoiceAddress(1, "BNB")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, err = svc.GenerateInvoiceAddress(1, "ETHEREUM")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, b, p, _ := svc.GetUserKeys(1, "BNB")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)
	// fmt.Println(b)
	//fmt.Println(p)

	// addr, err = svc.GenerateInvoiceAddress(1, "BNB")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, err = svc.GenerateInvoiceAddress(2, "TRON")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, err = svc.GenerateUserAddress(2, "TRON")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, err = svc.GenerateUserAddress(2, "ETHEREUM")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	// addr, err = svc.GenerateUserAddress(2, "POLYGON")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(addr)

	///sep
	// redisClient := core.OpenRedis("localhost:6379", "", 0)
	// redisRepo := repository.NewRedisRepo(redisClient)
	// redisRepo.AddAddressToUnderWatch("TQtCP5BsVAiMF1TJJmuKZw72UhAX91xuHJ")
	// redisRepo.AddAddressToUnderWatch("TQBi5Ki8DoyJue9a66Dkvvh3AUJW7n2exZ")
	// tronWatcher, _ := service.IWatcherFactory("TRON", redisRepo)
	// lastCheckedBlock, _ := tronWatcher.GetProvider().GetLatestBlockNumber()
	// redisClient.Set(context.Background(), "lastCheckedBlock", strconv.FormatUint(lastCheckedBlock, 10), 0)
	// for {
	// 	tronWatcher.ProcesseUnProcessedBlocks()
	// 	time.Sleep(time.Second * 2)
	// }

	redisClient := core.OpenRedis("localhost:6379", "", 0)
	redisRepo := repository.NewRedisRepo(redisClient)
	redisRepo.AddAddressToUnderWatch("0xa43ca591292aC8903FA492e1EaD511EF3aFF5289")
	ethereumWatcher, _ := service.IWatcherFactory("POLYGON", redisRepo)
	lastCheckedBlock, _ := ethereumWatcher.GetProvider().GetLatestBlockNumber()
	redisClient.Set(context.Background(), "lastCheckedBlock", strconv.FormatUint(lastCheckedBlock, 10), 0)
	for {
		go ethereumWatcher.ProcesseUnProcessedBlocks()
		time.Sleep(time.Second * 1)
	}

}

// package main

// import (
// 	"fmt"
// 	"github.com/btcsuite/btcd/chaincfg"
// 	"github.com/btcsuite/btcutil/hdkeychain"
// )

// func main() {
// 	// Define your seed string. Make sure to keep it secret and secure.
// 	seed := "6a2f85cb3a5ef0a7d7747bbfe1e439f0"

// 	// Define the network parameters for the Bitcoin mainnet
// 	params := &chaincfg.MainNetParams

// 	// Create a master key from the seed
// 	masterKey, err := hdkeychain.NewMaster([]byte(seed), params)
// 	if err != nil {
// 		fmt.Println("Error creating master key:", err)
// 		return
// 	}

// 	// Derive a child key (e.g., for an account or purpose)
// 	// You can derive more keys as needed for your use case
// 	childKey, err := masterKey.Child(0) // Derive the first child key
// 	if err != nil {
// 		fmt.Println("Error deriving child key:", err)
// 		return
// 	}

// 	// Get the extended public key (xpub) for this child key
// 	xpub, err := childKey.Neuter()
// 	if err != nil {
// 		fmt.Println("Error getting xpub:", err)
// 		return
// 	}

// 	fmt.Printf("Master Private Key: %s\n", masterKey.String())
// 	fmt.Printf("Master Extended Private Key (xprv): %s\n", masterKey.String())
// 	fmt.Printf("Child Extended Public Key (xpub): %s\n", xpub.String())
// }

// package main

// import (
//     "fmt"
// )

// func main() {
//     // Define an array of bytes
//     byteArray := []byte{58,60, 230, 205, 71 ,239, 166, 61 ,128 ,104, 95, 255, 127, 4, 131, 241, 226, 140, 10, 60, 75, 106, 228, 226, 12, 154, 246, 145, 89, 219, 16, 140, 124, 139, 101, 65, 61, 123, 12, 179, 189, 90, 143, 90, 97 ,48, 49 ,216 ,116 ,26 ,3 ,237, 45, 131, 133, 33, 38, 46, 11, 180 ,142 ,141 ,216 ,236}

//     // Convert the byte array to a string
//     str := string(byteArray)

//     // Print the resulting string
//     fmt.Println(str)
// }

// package main

// import (
// 	"fmt"

// 	"github.com/foxnut/go-hdwallet"
// 	MMM "github.com/fbsobreira/gotron-sdk/pkg/address"
// )

// var (
// 	mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
// )

// func main() {

// 	master, err := hdwallet.NewKey(
// 		hdwallet.Mnemonic(mnemonic),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	wallet, _ := master.GetWallet(hdwallet.AddressIndex(999999999))
// 	fmt.Println("------------------------------")
// 	fmt.Println(wallet.GetKey().PrivateHex())
// 	fmt.Println(wallet.GetKey().PublicHex(true))
// 	fmt.Println("------------------------------")
// 	Pub := wallet.GetKey().PublicECDSA
// 	addbyte := MMM.PubkeyToAddress(*Pub)

// 	fmt.Println( addbyte)

// 	fmt.Println("#################33")
// 	// BTC: 1AwEPfoojHnKrhgt1vfuZAhrvPrmz7Rh4
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1))
// 	address, _ := wallet.GetAddress()
// 	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
// 	addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()
// 	fmt.Println("BTC: ", address, addressP2WPKH, addressP2WPKHInP2SH)

// 	// BCH: 1CSBT18sjcCwLCpmnnyN5iqLc46Qx7CC91
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.BCH))
// 	address, _ = wallet.GetAddress()
// 	addressBCH, _ := wallet.GetKey().AddressBCH()
// 	fmt.Println("BCH: ", address, addressBCH)

// 	// LTC: LLCaMFT8AKjDTvz1Ju8JoyYXxuug4PZZmS
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.LTC))
// 	address, _ = wallet.GetAddress()
// 	fmt.Println("LTC: ", address)

// 	// DOGE: DHLA3rJcCjG2tQwvnmoJzD5Ej7dBTQqhHK
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.DOGE))
// 	address, _ = wallet.GetAddress()
// 	fmt.Println("DOGE:", address)

// 	// ETH: 0x37039021cBA199663cBCb8e86bB63576991A28C1
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
// 	address, _ = wallet.GetAddress()
// 	fmt.Println("ETH: ", address)

// 	// ETC: 0x480C69E014C7f018dAbF17A98273e90f0b0680cf
// 	wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETC))
// 	address, _ = wallet.GetAddress()
// 	fmt.Println("ETC: ", address)
// }
// mnemonic := "range sheriff try enroll deer over ten level bring display stamp recycle"
// newSeed, err :=  hdwallet.NewSeed(mnemonic, "4585d78541524d", "EN")
// if err != nil{
// 	panic(err)
// }
// fmt.Println(fmt.Sprint(newSeed))
// fmt.Println( string([]byte("6a2f85cb3a5ef0a7d7747bbfe1e439f0")))
