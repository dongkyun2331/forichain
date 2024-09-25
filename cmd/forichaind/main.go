package main

import (
	"io"
	"os"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/dongkyun2331/forichain/app"
	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command, app.Configurator) {
    rootCmd := &cobra.Command{
        Use:   "forichaind",
        Short: "Forichain App Daemon (server)",
    }

    cfg := app.Configurator{}

    appCreator := func(logger log.Logger, db dbm.DB, traceStore io.Writer) types.Application {
        return app.NewForiChainApp(logger, db)
    }

    appExporter := func(logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string) (types.ExportedApp, error) {
        foriApp := app.NewForiChainApp(logger, db)
        return foriApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
    }

    // 서버 커맨드 추가
    server.AddCommands(rootCmd, "forichain", appCreator, appExporter, nil)

    return rootCmd, cfg
}

func main() {
    rootCmd, _ := NewRootCmd()  // NewRootCmd 호출

    if err := cmd.Execute(rootCmd, "", ""); err != nil {
        os.Exit(1)
    }
}
