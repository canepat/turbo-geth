package commands

import (
	"github.com/c2h5oh/datasize"
	"github.com/ledgerwatch/turbo-geth/cmd/utils"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/internal/debug"
	"github.com/ledgerwatch/turbo-geth/migrations"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "integration",
	Short: "long and heavy integration tests for turbo-geth",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := utils.SetupCobra(cmd); err != nil {
			panic(err)
		}

		if len(chaindata) > 0 {
			db := openDatabase()
			defer db.Close()
			if err := migrations.NewMigrator().Apply(db, datadir); err != nil {
				panic(err)
			}

			err := SetSnapshotKV(db, snapshotDir, snapshotMode)
			if err != nil {
				panic(err)
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		defer utils.StopDebug()
	},
}

func RootCommand() *cobra.Command {
	utils.CobraFlags(rootCmd, append(debug.Flags, utils.MetricFlags...))
	return rootCmd
}

func openDatabase() *ethdb.ObjectDatabase {
	if mapSizeStr != "" {
		var mapSize datasize.ByteSize
		must(mapSize.UnmarshalText([]byte(mapSizeStr)))
		return ethdb.NewObjectDatabase(ethdb.NewLMDB().Path(chaindata).MapSize(mapSize).MustOpen())
	}
	return ethdb.MustOpen(chaindata)
}
