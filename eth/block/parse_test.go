package block

import (
	"ethservice/eth/models"
	"testing"
)

func TestParseBlockData(t *testing.T) {
	tests := map[string]struct {
		testCase          []byte
		expectedBlockData *models.BlockData
		throwsError       bool
	}{
		"bad input": {testCase: []byte(`{"status":"0","message":"NOTOK","result":"Max rate limit reached, please use API Key for higher rate limit"}`),
			expectedBlockData: nil,
			throwsError:       true,
		},

		"non-existing block": {testCase: []byte(`{"jsonrpc":"2.0","id":1,"va2323r":null}`),
			expectedBlockData: &models.BlockData{},
			throwsError:       false,
		},

		"block with no transactions": {testCase: []byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0xf93d1faf8","extraData":"0x476574682f4c5649562f76312e302e302f6c696e75782f676f312e342e32","gasLimit":"0x1388","gasUsed":"0x0","hash":"0x83f46170d084506ff3d7c62cf83bcfac6779ab3be4f83ff4847456a3c0ef5b40","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0xbb7b8287f3f0a933474a79eae42cbca977791171","mixHash":"0xaf720894805855798e054c5b3c2c11f7701a70495a9a8817890c97975abafa96","nonce":"0xd4e93cd18afe50e3","number":"0xafb","parentHash":"0xedbdc26498288dd86f649b5d177986e4e67bfd64b7b041a7ee7c7f09893a5b84","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x220","stateRoot":"0xabeac6b29da1e1d03772307243a446ed61fffc70878bd550da58e36b2b0d2914","timestamp":"0x55ba5dc6","totalDifficulty":"0x5d778e44c85d","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}}`),
			expectedBlockData: &models.BlockData{Transactions: []models.Transaction{}},
			throwsError:       false,
		},

		"block with 1 transaction": {testCase: []byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x409e254f6f7","extraData":"0x476574682f76312e302e312d38326566323666362f6c696e75782f676f312e34","gasLimit":"0x2fefd8","gasUsed":"0x5208","hash":"0x485af7ca0c64ca13f2c829fd1c6ae2dcdd8a185e22d39612f75741372b9fa7fe","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x20c6efe179c293690a5b3375f5bd81c2f07ed037","mixHash":"0x766f394002109afef4b7e46c70e4c61499344351a1cc99f15584ad02c2f4a7b9","nonce":"0xe94c564608f12898","number":"0x1acdd","parentHash":"0xc27ddfe052e15e231a0c7a4d2977c37513b873930116ee226a63f10ac46ef0e4","receiptsRoot":"0x34ab5d0280c2fa7ccb5b23640c3d6b65a811147358d93efe4ed4b7a7388347b1","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x298","stateRoot":"0x765cc1fcf10e45f90da70a577796a22d0a9a9825fc32d69f8d3bee5de3f67c3e","timestamp":"0x55d41544","totalDifficulty":"0x2eccbef944e93a2","transactions":[{"blockHash":"0x485af7ca0c64ca13f2c829fd1c6ae2dcdd8a185e22d39612f75741372b9fa7fe","blockNumber":"0x1acdd","from":"0xa70a9a7040d542c632d8e0be7c5513a8a5bbc752","gas":"0x5208","gasPrice":"0xdc600a748","hash":"0xc161ffdb15423d1ed48673743c6067f365c3ff09fa923401fd6ae8f9ee2829f0","input":"0x","nonce":"0x117","to":"0x32be343b94f860124dc4fee278fdcbd38c102d88","transactionIndex":"0x0","value":"0x455f32e9884c2400","type":"0x0","v":"0x1b","r":"0xa49efa1aca3b5380ab94ea0dec422a6c1e41631dc0791f79730a2e8767c1af02","s":"0x704ff18a004201ce1fa2297a74e3c0da1bb92dbe8e05dc8176ed2f563441b010"}],"transactionsRoot":"0x89275c43dfe1fea0c32f6f881c5d541aa9aff3bd64ee4f87bdf68f484db1624d","uncles":[]}}`),
			expectedBlockData: &models.BlockData{Transactions: []models.Transaction{{Value: "0x455f32e9884c2400"}}},
			throwsError:       false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := ParseBlockData(test.testCase)
			if (err != nil) != test.throwsError {
				t.Error(err)
			}
			if err == nil {
				if len(result.Transactions) != len(test.expectedBlockData.Transactions) {
					t.Error("Expected: ", test.expectedBlockData, "Got: ", result)
				}
			}

		})
	}
}
