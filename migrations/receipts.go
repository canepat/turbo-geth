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
	Name: "receipts_cbor_encode_experiment_mdbx4",
	Up: func(db ethdb.Database, datadir string, OnLoadCommit etl.LoadCommitHandler) error {
		logEvery := time.NewTicker(30 * time.Second)
		defer logEvery.Stop()

		if err := db.Walk(dbutils.BlockReceiptsPrefix, nil, 0, func(k, v []byte) (bool, error) {
			blockNum := binary.BigEndian.Uint64(k[:8])
			select {
			default:
			case <-logEvery.C:
				log.Info("Migration progress", "blockNum", blockNum)
			}
			k, v = common.CopyBytes(k), common.CopyBytes(v)
			if err := db.Delete(dbutils.BlockReceiptsPrefix, k); err != nil {
				return false, err
			}
			if err := db.Put(dbutils.BlockReceiptsPrefix, k, v); err != nil {
				return false, err
			}
			return true, nil
		}); err != nil {
			return err
		}

		return OnLoadCommit(db, nil, true)
	},
}
