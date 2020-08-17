package types

import (
	"fmt"
	"strings"

	"github.com/bhpnet/bhp-dev/x/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	DefaultFeeIssue  = "2500"
	DefaultFeeMint   = "10"
	DefaultFeeBurn   = "10"
	DefaultFeeModify = "0"
	DefaultFeeChown  = "10"
)

var (
	KeyFeeIssue  = []byte("FeeIssue")
	KeyFeeMint   = []byte("FeeMint")
	KeyFeeBurn   = []byte("FeeBurn")
	KeyFeeModify = []byte("FeeModify")
	KeyFeeChown  = []byte("FeeChown")
)

var _ params.ParamSet = &Params{}

// mint parameters
type Params struct {
	FeeIssue  sdk.Coin `json:"issue_fee"`
	FeeMint   sdk.Coin `json:"mint_fee"`
	FeeBurn   sdk.Coin `json:"burn_fee"`
	FeeModify sdk.Coin `json:"modify_fee"`
	FeeChown  sdk.Coin `json:"transfer_ownership_fee"`
}

// ParamKeyTable for auth module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyFeeIssue, &p.FeeIssue},
		{KeyFeeMint, &p.FeeMint},
		{KeyFeeBurn, &p.FeeBurn},
		{KeyFeeModify, &p.FeeModify},
		{KeyFeeChown, &p.FeeChown},
	}
}

//Bhp27
func MustNewIntFromStr(s string) sdk.Int {
	dec, err := sdk.NewIntFromString(s)
	if err != true {
		panic(err)
	}
	return dec
}

//Bhp26
// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		FeeIssue:  sdk.NewCoin(common.NativeToken, MustNewIntFromStr(DefaultFeeIssue)),
		FeeMint:   sdk.NewCoin(common.NativeToken, MustNewIntFromStr(DefaultFeeMint)),
		FeeBurn:   sdk.NewCoin(common.NativeToken, MustNewIntFromStr(DefaultFeeBurn)),
		FeeModify: sdk.NewCoin(common.NativeToken, MustNewIntFromStr(DefaultFeeModify)),
		FeeChown:  sdk.NewCoin(common.NativeToken, MustNewIntFromStr(DefaultFeeChown)),
	}
}

// String implements the stringer interface.
func (p Params) String() string {
	var sb strings.Builder
	sb.WriteString("Params: \n")
	sb.WriteString(fmt.Sprintf("FeeIssue: %s\n", p.FeeIssue))
	sb.WriteString(fmt.Sprintf("FeeMint: %s\n", p.FeeMint))
	sb.WriteString(fmt.Sprintf("FeeBurn: %s\n", p.FeeBurn))
	sb.WriteString(fmt.Sprintf("FeeModify: %s\n", p.FeeModify))
	sb.WriteString(fmt.Sprintf("FeeChown: %s\n", p.FeeChown))

	return sb.String()
}
