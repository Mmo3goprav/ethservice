package transport

import (
	"ethservice/eth/models"
	"testing"
)

func TestGetBlock(t *testing.T) {
	//w := new(http.ResponseWriter)

	tests := map[string]struct {
		blockNumber        int
		expectedResponse   *models.BlockDataResponse
		expectToThrowError bool
	}{
		// "Block normal": {
		// 	blockNumber:      11508993,
		// 	expectedResponse: &models.BlockDataResponse{Transactions: 241, Amount: 1130.9870854468265},
		// },

		"Block with 0 transactions": {
			blockNumber:        10,
			expectedResponse:   &models.BlockDataResponse{Transactions: 0, Amount: 0},
			expectToThrowError: false,
		},

		"Non-existing block": {
			blockNumber:        100000000000000,
			expectedResponse:   nil,
			expectToThrowError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resp, err := GetBlockData(test.blockNumber, "ANDAVVBEEUCWUR3TEP8DTA2DPRIAB796WI")
			if (err != nil) != test.expectToThrowError {
				t.Errorf("Returned error doesn't match expected error in test %s", name)
			}

			if err == nil {
				if resp.Amount != test.expectedResponse.Amount && resp.Transactions != test.expectedResponse.Transactions {
					t.Errorf("Expected response %v, got %v", test.expectedResponse, resp)
				}
			}

		})
	}

}
