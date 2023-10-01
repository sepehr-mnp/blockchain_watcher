package trongrid

type GetBlockBody struct {
	Num int64 `json:"num"`
}

type GetTransactionInfoByIDBody struct {
	Value string `json:"value"`
}

type GetLastBlockBody struct {
	Detail bool `json:"detail"`
}

type Block struct {
	BlockID      string        `json:"blockID"`
	BlockHeader  BlockHeader   `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type HeaderRawData struct {
	Number         int64  `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int32  `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

type BlockHeader struct {
	RawData          HeaderRawData `json:"raw_data"`
	WitnessSignature string        `json:"witness_signature"`
}

type Value struct {
	Amount          float64 `json:"amount"`
	OwnerAddress    string  `json:"owner_address"`
	ToAddress       string  `json:"to_address"`
	Data            string  `json:"data"`
	ContractAddress string  `json:"contract_address"`
}

type Parameter struct {
	Value   Value  `json:"value"`
	TypeUrl string `json:"type_url"`
}

type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}

type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	FeeLimit      int64      `json:"fee_limit"`
	Timestamp     int64      `json:"timestamp"`
}

type Ret struct {
	ContractRet string `json:"contractRet"`
}

type Transaction struct {
	Ret        []Ret    `json:"ret"`
	Signature  []string `json:"signature"`
	TxID       string   `json:"txID"`
	RawData    RawData  `json:"raw_data"`
	RawDataHex string   `json:"raw_data_hex"`
}

type TransactionInfo struct {
	ID             string
	Fee            float64
	BlockNumber    int64
	BlockTimeStamp int64
}
