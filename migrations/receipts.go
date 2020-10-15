package migrations

import (
	"encoding/binary"
	"time"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/common/etl"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
)

var receiptsCborEncode = Migration{
	Name: "receipts_cbor_encode_experiment_mdbx3",
	Up: func(db ethdb.Database, datadir string, OnLoadCommit etl.LoadCommitHandler) error {
		logEvery := time.NewTicker(30 * time.Second)
		defer logEvery.Stop()

		buf := make([]byte, 0, 100_000)
		if err := db.Walk(dbutils.BlockReceiptsPrefix, nil, 0, func(k, v []byte) (bool, error) {
			blockNum := binary.BigEndian.Uint64(k[:8])
			select {
			default:
			case <-logEvery.C:
				log.Info("Migration progress", "blockNum", blockNum)
			}

			if err := db.Delete(dbutils.BlockReceiptsPrefix, common.CopyBytes(k)); err != nil {
				return false, err
			}

			if err := db.Put(dbutils.BlockReceiptsPrefix, common.CopyBytes(k), common.CopyBytes(buf)); err != nil {
				return false, err
			}
			return true, nil
		}); err != nil {
			return err
		}

		return OnLoadCommit(db, nil, true)
	},
}
