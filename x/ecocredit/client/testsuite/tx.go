package testsuite

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	testAccount sdk.AccAddress
	classInfo   *ecocredit.ClassInfo
	batchInfo   *ecocredit.BatchInfo
}

const (
	validCreditType = "carbon"
	validMetadata   = "AQ=="
	classId         = "C01"
	batchDenom      = "C01-20210101-20210201-001"
)

var validMetadataBytes = []byte{0x1}

func RunCLITests(t *testing.T, cfg network.Config) {
	suite.Run(t, NewIntegrationTestSuite(cfg))

	// setup another cfg for testing ecocredit enabled class creators list.
	genesisState := ecocredit.DefaultGenesisState()
	genesisState.Params.AllowlistEnabled = true
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState[ecocredit.ModuleName] = bz
	suite.Run(t, NewAllowListEnabledTestSuite(cfg))
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

// Write a MsgCreateBatch to a new temporary file and return the filename
func (s *IntegrationTestSuite) writeMsgCreateBatchJSON(msg *ecocredit.MsgCreateBatch) string {
	bytes, err := s.network.Validators[0].ClientCtx.Codec.MarshalJSON(msg)
	s.Require().NoError(err)

	return testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]

	// create an account for val
	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator0", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	_, a1pub, a1 := testdata.KeyTestPubAddr()
	val.ClientCtx.Keyring.SavePubKey("throwaway", a1pub, hd.Secp256k1Type)

	account := sdk.AccAddress(info.GetPubKey().Address())
	for _, acc := range []sdk.AccAddress{account, a1} {
		_, err = banktestutil.MsgSendExec(
			val.ClientCtx,
			val.Address,
			acc,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(2000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		)
		s.Require().NoError(err)
	}

	s.testAccount = account

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	// Create a few credit classes
	for i := 0; i < 4; i++ {
		out, err := cli.ExecTestCLICmd(val.ClientCtx, client.TxCreateClassCmd(),
			append(
				[]string{
					val.Address.String(),
					validCreditType,
					validMetadata,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				},
				commonFlags...,
			),
		)

		s.Require().NoError(err, out.String())
		var txResp = sdk.TxResponse{}
		s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
		s.Require().Equal(uint32(0), txResp.Code, out.String())
	}

	// Store the first one in the test suite
	s.classInfo = &ecocredit.ClassInfo{
		ClassId:    classId,
		Admin:      val.Address.String(),
		Issuers:    []string{val.Address.String()},
		CreditType: ecocredit.DefaultParams().CreditTypes[0],
		Metadata:   validMetadataBytes,
	}

	startDate, err := client.ParseDate("start date", "2021-01-01")
	s.Require().NoError(err)
	endDate, err := client.ParseDate("end date", "2021-02-01")
	s.Require().NoError(err)

	msgCreateBatch := ecocredit.MsgCreateBatch{
		ClassId: classId,
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          val.Address.String(),
				TradableAmount:     "100",
				RetiredAmount:      "0.000001",
				RetirementLocation: "AB",
			},
		},
		Metadata:        validMetadataBytes,
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}

	// Write MsgCreateBatch to a temporary file
	batchFile := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Create a few credit batches
	for i := 0; i < 4; i++ {
		out, err := cli.ExecTestCLICmd(val.ClientCtx, client.TxCreateBatchCmd(),
			append(
				[]string{
					batchFile,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				},
				commonFlags...,
			),
		)

		s.Require().NoError(err, out.String())
		txResp := sdk.TxResponse{}
		s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
		s.Require().Equal(uint32(0), txResp.Code, out.String())
	}

	// Store the first one in the test suite
	s.batchInfo = &ecocredit.BatchInfo{
		ClassId:         classId,
		BatchDenom:      batchDenom,
		Issuer:          val.Address.String(),
		TotalAmount:     "100.000001",
		Metadata:        []byte{0x01},
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) commonTxFlags() []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
}

var flagOutputJSON = fmt.Sprintf("--%s=json", tmcli.OutputFlag)

func makeFlagFrom(from string) string {
	return fmt.Sprintf("--%s=%s", flags.FlagFrom, from)
}

func (s *IntegrationTestSuite) TestTxCreateClass() {
	val0 := s.network.Validators[0]
	val1 := s.network.Validators[1]
	clientCtx := val0.ClientCtx

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		respCode          uint32
		expectedClassInfo *ecocredit.ClassInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "accepts 3 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "accepts 3 arg(s), received 4",
		},
		{
			name: "invalid issuer",
			args: append(
				[]string{
					"abcde",
					validCreditType,
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid metadata",
			args: append(
				[]string{
					val0.Address.String(),
					validCreditType,
					"=",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "metadata is malformed, proper base64 string is required",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					val0.Address.String(),
					validCreditType,
					validMetadata,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "invalid credit type",
			args: append(
				[]string{
					val0.Address.String(),
					"caarbon",
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      false,
			expectedErrMsg: "caarbon is not a valid credit type",
			respCode:       29,
		},
		{
			name: "single issuer",
			args: append(
				[]string{
					val0.Address.String(),
					validCreditType,
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Admin:    val0.Address.String(),
				Issuers:  []string{val0.Address.String()},
				Metadata: []byte{0x1},
			},
		},
		{
			name: "single issuer with from key-name",
			args: append(
				[]string{
					val0.Address.String(),
					validCreditType,
					validMetadata,
					makeFlagFrom("node0"),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Admin:    val0.Address.String(),
				Issuers:  []string{val0.Address.String()},
				Metadata: []byte{0x1},
			},
		},
		{
			name: "multiple issuers",
			args: append(
				[]string{
					strings.Join(
						[]string{
							val0.Address.String(),
							val1.Address.String(),
						},
						",",
					),
					validCreditType,
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Admin:    val0.Address.String(),
				Issuers:  []string{val0.Address.String(), val1.Address.String()},
				Metadata: []byte{0x1},
			},
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					val0.Address.String(),
					validCreditType,
					validMetadata,
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Admin:    val0.Address.String(),
				Issuers:  []string{val0.Address.String()},
				Metadata: []byte{0x1},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := client.TxCreateClassCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.respCode, res.Code)
				if tc.respCode == 0 {
					classIdFound := false
					for _, e := range res.Logs[0].Events {
						if e.Type == proto.MessageName(&ecocredit.EventCreateClass{}) {
							for _, attr := range e.Attributes {
								if attr.Key == "class_id" {
									classIdFound = true
									classId := strings.Trim(attr.Value, "\"")

									queryCmd := client.QueryClassInfoCmd()
									queryArgs := []string{classId, flagOutputJSON}
									queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
									s.Require().NoError(err, queryOut.String())
									var queryRes ecocredit.QueryClassInfoResponse
									s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))

									s.Require().Equal(tc.expectedClassInfo.Admin, queryRes.Info.Admin)
									s.Require().Equal(tc.expectedClassInfo.Issuers, queryRes.Info.Issuers)
									s.Require().Equal(tc.expectedClassInfo.Metadata, queryRes.Info.Metadata)
								}
							}
						}
					}
					s.Require().True(classIdFound)
				} else {
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCreateBatch() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// Write some invalid JSON to a file
	invalidJsonFile := testutil.WriteToNewTempFile(s.T(), "{asdljdfklfklksdflk}")

	// Create a valid MsgCreateBatch
	startDate, err := client.ParseDate("start date", "2021-01-01")
	s.Require().NoError(err)
	endDate, err := client.ParseDate("end date", "2021-02-01")
	s.Require().NoError(err)

	msgCreateBatch := ecocredit.MsgCreateBatch{
		ClassId: s.classInfo.ClassId,
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          s.network.Validators[1].Address.String(),
				TradableAmount:     "100",
				RetiredAmount:      "0.000001",
				RetirementLocation: "AB",
			},
		},
		Metadata:        validMetadataBytes,
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}

	validBatchJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid class id
	msgCreateBatch.ClassId = "abcde"
	invalidClassIdJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing start date
	msgCreateBatch.ClassId = s.classInfo.ClassId
	msgCreateBatch.StartDate = nil
	missingStartDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing end date
	msgCreateBatch.StartDate = &startDate
	msgCreateBatch.EndDate = nil
	missingEndDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing project location
	msgCreateBatch.EndDate = &endDate
	msgCreateBatch.ProjectLocation = ""
	missingProjectLocationJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance recipient
	msgCreateBatch.ProjectLocation = "AB"
	msgCreateBatch.Issuance[0].Recipient = "abcde"
	invalidRecipientJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance tradable amount
	msgCreateBatch.Issuance[0].Recipient = s.network.Validators[1].Address.String()
	msgCreateBatch.Issuance[0].TradableAmount = "abcde"
	invalidTradableAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance retired amount
	msgCreateBatch.Issuance[0].TradableAmount = "100"
	msgCreateBatch.Issuance[0].RetiredAmount = "abcde"
	invalidRetiredAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance retirement location
	msgCreateBatch.Issuance[0].RetiredAmount = "0.000001"
	msgCreateBatch.Issuance[0].RetirementLocation = "abcde"
	invalidRetirementLocationJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		errInTxResponse   bool
		expectedErrMsg    string
		expectedBatchInfo *ecocredit.BatchInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"r", "e", "g", "e", "n"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 5",
		},
		{
			name: "invalid json",
			args: append(
				[]string{
					invalidJsonFile.Name(),
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid character",
		},
		{
			name: "invalid class id",
			args: append(
				[]string{
					invalidClassIdJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:       true,
			errInTxResponse: false,
			expectedErrMsg:  "class ID didn't match the format",
		},
		{
			name: "missing start date",
			args: append(
				[]string{
					missingStartDateJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "must provide a start date for the credit batch: invalid request",
		},
		{
			name: "missing end date",
			args: append(
				[]string{
					missingEndDateJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "must provide an end date for the credit batch: invalid request",
		},
		{
			name: "missing project location",
			args: append(
				[]string{
					missingProjectLocationJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid location",
		},
		{
			name: "invalid issuance recipient",
			args: append(
				[]string{
					invalidRecipientJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid issuance tradable amount",
			args: append(
				[]string{
					invalidTradableAmountJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid decimal string",
		},
		{
			name: "invalid issuance retired amount",
			args: append(
				[]string{
					invalidRetiredAmountJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid decimal string",
		},
		{
			name: "invalid issuance retirement location",
			args: append(
				[]string{
					invalidRetirementLocationJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid location: abcde",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					validBatchJson,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid batch",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatchInfo: &ecocredit.BatchInfo{
				ClassId:         s.classInfo.ClassId,
				Issuer:          val.Address.String(),
				TotalAmount:     "100.000001",
				Metadata:        []byte{0x1},
				AmountCancelled: "0",
			},
		},
		{
			name: "valid batch with from key-name",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom("node0"),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatchInfo: &ecocredit.BatchInfo{
				ClassId:         s.classInfo.ClassId,
				Issuer:          val.Address.String(),
				TotalAmount:     "100.000001",
				Metadata:        []byte{0x1},
				AmountCancelled: "0",
			},
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom(val.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatchInfo: &ecocredit.BatchInfo{
				ClassId:         s.classInfo.ClassId,
				Issuer:          val.Address.String(),
				TotalAmount:     "100.000001",
				Metadata:        []byte{0x1},
				AmountCancelled: "0",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := client.TxCreateBatchCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(res.Code, 0)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				batchDenomFound := false
				for _, e := range res.Logs[0].Events {
					if e.Type == proto.MessageName(&ecocredit.EventCreateBatch{}) {
						for _, attr := range e.Attributes {
							if attr.Key == "batch_denom" {
								batchDenomFound = true
								batchDenom := strings.Trim(attr.Value, "\"")

								queryCmd := client.QueryBatchInfoCmd()
								queryArgs := []string{batchDenom, flagOutputJSON}
								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
								s.Require().NoError(err, queryOut.String())
								var queryRes ecocredit.QueryBatchInfoResponse
								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))

								s.Require().Equal(tc.expectedBatchInfo.ClassId, queryRes.Info.ClassId)
								s.Require().Equal(tc.expectedBatchInfo.Issuer, queryRes.Info.Issuer)
								s.Require().Equal(tc.expectedBatchInfo.TotalAmount, queryRes.Info.TotalAmount)
								s.Require().Equal(tc.expectedBatchInfo.Metadata, queryRes.Info.Metadata)
								s.Require().Equal(tc.expectedBatchInfo.AmountCancelled, queryRes.Info.AmountCancelled)
							}
						}
					}
				}
				s.Require().True(batchDenomFound)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxSend() {
	val0 := s.network.Validators[0]
	val1 := s.network.Validators[1]
	clientCtx := val0.ClientCtx

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := "[{batch_denom: abcde, tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]"
	invalidTradableAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"abcde\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidRetiredAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"abcde\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidRetirementLocationCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"abcde\"}]", s.batchInfo.BatchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name: "invalid recipient",
			args: append(
				[]string{
					"abcde",
					validCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid batch denom",
			args: append(
				[]string{
					val1.Address.String(),
					invalidBatchDenomCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid denom",
		},
		{
			name: "invalid tradable amount",
			args: append(
				[]string{
					val1.Address.String(),
					invalidTradableAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid decimal string",
		},
		{
			name: "invalid retired amount",
			args: append(
				[]string{
					val1.Address.String(),
					invalidRetiredAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid decimal string",
		},
		{
			name: "invalid retirement location",
			args: append(
				[]string{
					val1.Address.String(),
					invalidRetirementLocationCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid location: abcde",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := client.TxSendCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxRetire() {
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"5\"}]", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := "[{batch_denom: abcde, amount: \"5\"}]"
	invalidAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"abcde\"}]", s.batchInfo.BatchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name: "invalid batch denom",
			args: append(
				[]string{
					invalidBatchDenomCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid denom",
		},
		{
			name: "invalid amount",
			args: append(
				[]string{
					invalidAmountCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid decimal string",
		},
		{
			name: "invalid retirement location",
			args: append(
				[]string{
					validCredits,
					"abcde",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid location: abcde",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := client.TxRetireCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCancel() {
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	validCredits := fmt.Sprintf("5 %s", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := "5 abcde"
	invalidAmountCredits := fmt.Sprintf("abcde %s", s.batchInfo.BatchDenom)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name: "invalid batch denom",
			args: append(
				[]string{
					invalidBatchDenomCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid credit expression",
		},
		{
			name: "invalid amount",
			args: append(
				[]string{
					invalidAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid credit expression",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					validCredits,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					validCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validCredits,
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := client.TxCancelCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateAdmin() {
	// use this classId as to not corrupt other tests
	const classId = "C02"
	_, _, a1 := testdata.KeyTestPubAddr()
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid request: not enough args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:      "invalid request: no id",
			args:      []string{"", a1.String()},
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:      "invalid request: no admin address",
			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "new admin address is required",
		},
		{
			name:   "valid request",
			args:   append([]string{classId, a1.String(), makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr: false,
		},
		{
			name:   "valid test: from key-name",
			args:   append([]string{classId, a1.String(), makeFlagFrom("node0")}, s.commonTxFlags()...),
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassAdminCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query the class info
				query := client.QueryClassInfoCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res ecocredit.QueryClassInfoResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// check the admin has been changed
				s.Require().Equal(res.Info.Admin, tc.args[1])
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateMetadata() {
	// use C03 here as C02 will be corrupted by the admin change test
	const classId = "C03"
	newMetaData := base64.StdEncoding.EncodeToString([]byte("hello"))
	_, _, a1 := testdata.KeyTestPubAddr()
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid request: not enough args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:      "invalid request: bad id",
			args:      []string{"", a1.String()},
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:      "invalid request: no metadata",
			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "base64_metadata is required",
		},
		{
			name:      "invalid request: bad metadata",
			args:      append([]string{classId, "test", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "metadata is malformed, proper base64 string is required",
		},
		{
			name:   "valid request",
			args:   append([]string{classId, newMetaData, makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr: false,
		},
		{
			name:   "valid test: from key-name",
			args:   append([]string{classId, newMetaData, makeFlagFrom("node0")}, s.commonTxFlags()...),
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassMetadataCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query the credit class info
				query := client.QueryClassInfoCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res ecocredit.QueryClassInfoResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// check metadata changed
				b, err := base64.StdEncoding.DecodeString(newMetaData)
				s.Require().NoError(err)
				s.Require().Equal(res.Info.Metadata, b)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateIssuers() {
	const classId = "C03"
	_, _, a2 := testdata.KeyTestPubAddr()
	newIssuers := []string{s.testAccount.String(), a2.String()}
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid request: not enough args",
			args:      append([]string{makeFlagFrom(s.testAccount.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:      "invalid request: no id",
			args:      append([]string{"", s.testAccount.String(), makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:      "invalid request: bad issuer addresses",
			args:      append([]string{classId, "hello,world", makeFlagFrom(s.testAccount.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "invalid address",
		},
		{
			name:   "valid request",
			args:   append([]string{classId, fmt.Sprintf("%s,%s", newIssuers[0], newIssuers[1]), makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr: false,
		},
		{
			name:   "valid test: from key-name",
			args:   append([]string{classId, fmt.Sprintf("%s,%s", newIssuers[0], newIssuers[1]), makeFlagFrom("node0")}, s.commonTxFlags()...),
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassIssuersCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)

				// query the credit class info
				query := client.QueryClassInfoCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res ecocredit.QueryClassInfoResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// check issuers list was changed
				s.Require().NoError(err)
				s.Require().Equal(res.Info.Issuers, newIssuers)
			}
		})
	}
}
