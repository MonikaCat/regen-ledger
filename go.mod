module github.com/regen-network/regen-ledger

go 1.17

require (
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/cosmos/ibc-go v1.0.1
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/regen-ledger/types v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/x/data v0.0.0-20210602121340-fa967f821a6e
	github.com/regen-network/regen-ledger/x/ecocredit v0.0.0-20210602121340-fa967f821a6e
	github.com/regen-network/regen-ledger/x/group v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.13
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211117155847-120650a500bb // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.44.2-regen-1

replace github.com/regen-network/regen-ledger/types => ./types

replace github.com/regen-network/regen-ledger/orm => ./orm

replace github.com/regen-network/regen-ledger/x/data => ./x/data

replace github.com/regen-network/regen-ledger/x/ecocredit => ./x/ecocredit

replace github.com/regen-network/regen-ledger/x/group => ./x/group
