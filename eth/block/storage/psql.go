package storage

import (
	"database/sql"
	"log"
	"strconv"

	"ethservice/eth/models"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func InitTableDB(db *sql.DB) error {
	if db == nil {
		return errors.New("Database handle is nil")
	}

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS BLOCKCACHE (Blocknum INT, Transactions INT, Amount FLOAT);`)
	if err != nil {
		return errors.Wrap(err, "Database table creation failed")
	}

	log.Println("Database ready")
	return nil
}

func SaveBlockToDB(db *sql.DB, blockResponse *models.BlockDataResponse, blocknum int) error {
	if db == nil {
		return errors.New("Database handle is nil")
	}
	if blockResponse == nil {
		return errors.New("BlockData is nil")
	}

	sqlStatement := `
	INSERT INTO BLOCKCACHE (Blocknum, Transactions, Amount)
	VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, blocknum, blockResponse.Transactions, blockResponse.Amount)
	return err
}

func GetBlockFromDB(blocknum int, db *sql.DB) (*models.BlockDataResponse, error) {
	if db == nil {
		return nil, errors.New("Database handle is nil")
	}

	rows, err := db.Query("SELECT Transactions,Amount FROM BLOCKCACHE WHERE Blocknum = $1", strconv.Itoa(blocknum))
	if err != nil {
		return nil, errors.Wrap(err, "Database query failed")
	}
	defer rows.Close()

	if rows.Next() {
		var blockDataResponse models.BlockDataResponse
		err = rows.Scan(&blockDataResponse.Transactions, &blockDataResponse.Amount)
		if err != nil {
			return nil, errors.Wrap(err, "Row scan failed")
		}
		return &blockDataResponse, nil

	} else {
		return nil, errors.New("No such block in database")
	}
}
