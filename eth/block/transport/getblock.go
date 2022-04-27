package transport

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"ethservice/eth/block"
	"ethservice/eth/models"

	"github.com/pkg/errors"
)

func GetBlockData(blockNum int, apikey string) (*models.BlockDataResponse, error) {
	tag := fmt.Sprintf("%x", blockNum)

	url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=0x%s&boolean=true&apikey=%s", tag, apikey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "etherscan.io: GET request failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`"result":null`)
	if re.Match(body) {
		return nil, errors.New("block not found")
	}

	blockData, err := block.ParseBlockData(body)
	if err != nil {
		return nil, err
	}

	transactions := block.TotalTransactions(blockData)
	amount, err := block.TotalAmount(blockData)
	if err != nil {
		return nil, err
	}

	blockResponse := models.BlockDataResponse{
		Transactions: transactions,
		Amount:       amount,
	}

	return &blockResponse, nil
}
