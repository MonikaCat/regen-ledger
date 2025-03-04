package regen_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	cmd "github.com/regen-network/regen-ledger/v2/app/regen/cmd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"regenapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	err := cmd.Execute(rootCmd)
	require.NoError(t, err)
}
