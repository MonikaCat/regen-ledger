//go:build !experimental
// +build !experimental

// DONTCOVER

package app

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/authz"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ecocredittypes "github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/regen-network/regen-ledger/types/module/server"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler,
			upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
	}
}

// setCustomModules registers new modules with the server module manager.
// It does nothing here and returns an empty manager since we're not using experimental mode.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {
	return server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))
}
func setCustomKVStoreKeys() []string {
	return []string{}
}

func (app *RegenApp) registerUpgradeHandlers() {
	// This is the upgrade plan name we used in the gov proposal.
	upgradeName := "v2.0-upgrade"
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
		// 1st-time running in-store migrations, using 1 as fromVersion to
		// avoid running InitGenesis.
		// Explicitly skipping x/auth migrations. It is already patched in regen-ledger v1.0.
		fromVM := map[string]uint64{
			"auth":         auth.AppModule{}.ConsensusVersion(),
			"bank":         1,
			"capability":   1,
			"crisis":       1,
			"distribution": 1,
			"evidence":     1,
			"gov":          1,
			"mint":         1,
			"params":       1,
			"slashing":     1,
			"staking":      1,
			"upgrade":      1,
			"vesting":      1,
			"ibc":          1,
			"genutil":      1,
			"transfer":     1,
			"ecocredit":    1, // we don't run InitGenesis for ecocredit in `RunMigrations`, but manually instead.
		}

		gen := ecocredittypes.DefaultGenesisState()
		gen.Params.AllowlistEnabled = true
		gen.Params.CreditClassFee = sdk.NewCoins(sdk.NewCoin("uregen", ecocredittypes.DefaultCreditClassFeeTokens))

		modules := make(map[string]json.RawMessage)
		modules[ecocredittypes.ModuleName] = app.cdc.MustMarshalJSON(gen)
		app.smm.InitGenesis(ctx, modules, []abci.ValidatorUpdate{})

		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{authz.ModuleName, feegrant.ModuleName, ecocredittypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomKeeprs(_ *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec, _ govtypes.Router, _ string) {
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{}
}

func initCustomParamsKeeper(_ *paramskeeper.Keeper) {}
