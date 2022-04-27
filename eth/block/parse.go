package block

import (
	"encoding/json"
	"ethservice/eth/models"
	"fmt"
	"log"
	"math/big"

	"github.com/pkg/errors"
)

func CreateBlockResponse(blockResponse *models.BlockDataResponse) []byte {
	resp, err := json.Marshal(blockResponse)
	if err != nil {
		log.Print(err)
	}

	return resp
}

func ParseBlockData(rawBlockData []byte) (*models.BlockData, error) {
	Block := struct {
		Result models.BlockData `json:"result"`
	}{}

	err := json.Unmarshal(rawBlockData, &Block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal block data")
	}
	fmt.Println(Block.Result)
	return &Block.Result, nil
}

func TotalAmount(b *models.BlockData) (float64, error) {
	totalWei := big.NewInt(0)
	for _, transaction := range b.Transactions {
		if len(transaction.Value) < 3 {
			return 0, errors.New("invalid transaction value")
		}

		amountWei := &big.Int{}
		amountWei, success := amountWei.SetString(transaction.Value[2:], 16)
		if !success {
			return 0, errors.New("Failed to parse amount")
		}
		totalWei = totalWei.Add(totalWei, amountWei)
	}

	totalWeiFloat := new(big.Float).SetInt(totalWei)
	// multiply by 1x10^-18 to convert from Wei to ETH
	totalWeiFloat = totalWeiFloat.Mul(totalWeiFloat, big.NewFloat(1e-18))
	totalEth, _ := totalWeiFloat.Float64()

	return totalEth, nil
}

func TotalTransactions(b *models.BlockData) int {
	return len(b.Transactions)
}
