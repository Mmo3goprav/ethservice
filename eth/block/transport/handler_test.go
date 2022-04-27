package transport

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	//w := new(http.ResponseWriter)

	tests := map[string]struct {
		blockNumber      string
		expectedResponse string
	}{
		"Block normal": {
			blockNumber:      "11508993",
			expectedResponse: `{"transactions":241,"amount":1130.9870854468265}`,
		},

		"Block not a number": {
			blockNumber:      "&63#%^@#",
			expectedResponse: "",
		},

		"Block with 0 transactions": {
			blockNumber:      "10",
			expectedResponse: `{"transactions":0,"amount":0}`,
		},

		"Non-existing block": {
			blockNumber:      "100000000000000",
			expectedResponse: `{"transactions":0,"amount":0}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// my own secret key, use with care :)
			Handler(w, test.blockNumber, nil, "ANDAVVBEEUCWUR3TEP8DTA2DPRIAB796WI")

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if string(body) != test.expectedResponse {
				t.Errorf("Expected response %s, got %s", test.expectedResponse, string(body))
			}

			resp.Body.Close()

			time.Sleep(time.Second)
		})
	}

}
