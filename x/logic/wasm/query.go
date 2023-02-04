package wasm

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okp4/okp4d/x/logic/keeper"
	"github.com/okp4/okp4d/x/logic/types"
)

// AskQuery contains Ask gRPC request parameters, it is redefined to keep control in case of eventual breaking change
// in the logic module definition.
type AskQuery struct {
	Program string `json:"program"`
	Query   string `json:"query"`
}

// LogicQuerier ease the bridge between the logic module with the wasm CustomQuerier to allow wasm contracts to query
// the logic module.
type LogicQuerier struct {
	k keeper.Keeper
}

// MakeLogicQuerier creates a new LogicQuerier based on the logic keeper.
func MakeLogicQuerier(keeper keeper.Keeper) LogicQuerier {
	return LogicQuerier{
		k: keeper,
	}
}

// Ask is a proxy method with the gRPC request, returning the result in the json format.
func (querier LogicQuerier) Ask(ctx sdk.Context, query *AskQuery) ([]byte, error) {
	resp, err := querier.k.Ask(ctx, &types.QueryServiceAskRequest{
		Program: query.Program,
		Query:   query.Query,
	})
	if err != nil {
		return nil, err
	}

	return json.Marshal(resp)
}
