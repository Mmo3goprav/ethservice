package transport

import (
	"database/sql"
	"ethservice/eth/block"
	"ethservice/eth/block/storage"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, requestedBlock string, db *sql.DB, apikey string) {

	// make sure we get a valid block number
	blockNum, err := strconv.Atoi(requestedBlock)
	if err != nil {
		log.Printf("Invalid block number: %s", requestedBlock)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//search for block data in database first
	blockDataResponse, err := storage.GetBlockFromDB(blockNum, db)
	if err != nil {
		log.Print("block not found in db:", err)
		// if block not cached in database, look it up on etherscan.io
		blockDataResponse, err = GetBlockData(blockNum, apikey)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// cache newly retrieved block
		err = storage.SaveBlockToDB(db, blockDataResponse, blockNum)
		if err != nil {
			log.Print(err)
		}
	}

	// store block data in json and send as a response to client
	fmt.Fprint(w, string(block.CreateBlockResponse(blockDataResponse)))
}
