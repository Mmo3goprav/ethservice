package models

// BlockDataResponse is the response for a block request
type BlockDataResponse struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}

// BlockData is the data for a block
type BlockData struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Value string `json:"value"`
}
