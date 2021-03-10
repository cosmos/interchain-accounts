package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fmt"
	"github.com/interchainberlin/ica/x/inter-tx/types"
)

// IBCAccountFromData implements the Query/IBCAccount gRPC method
func (k Keeper) IBCAccountFromAddress(ctx context.Context, req *types.QueryIBCAccountFromAddressRequest) (*types.QueryIBCAccountFromAddressResponse, error) {
	fmt.Println("sean")
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Port == "" {
		return nil, status.Error(codes.InvalidArgument, "port cannot be empty")
	}

	if req.Channel == "" {
		return nil, status.Error(codes.InvalidArgument, "channel cannot be empty")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ibcAccount, err := k.GetIBCAccount(sdkCtx, req.Port, req.Channel, req.Address)
	if err != nil {
		return nil, err
	}

	return &ibcAccount, nil
}
