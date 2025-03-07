package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	intenttypes "github.com/warden-protocol/wardenprotocol/warden/x/intent/types"
	"github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta2"
)

func (k msgServer) NewSignTransactionRequest(goCtx context.Context, msg *v1beta2.MsgNewSignTransactionRequest) (*intenttypes.MsgActionCreated, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	key, err := k.keys.Get(ctx, msg.KeyId)
	if err != nil {
		return nil, err
	}

	ws, err := k.spaces.Get(ctx, key.SpaceId)
	if err != nil {
		return nil, err
	}

	keychain, err := k.keychains.Get(ctx, key.KeychainId)
	if err != nil {
		return nil, err
	}

	if !keychain.IsActive {
		return nil, fmt.Errorf("keychain is not active")
	}

	act, err := k.intentKeeper.AddAction(ctx, msg.Creator, msg, ws.SignIntentId, msg.Btl)
	if err != nil {
		return nil, err
	}

	return &intenttypes.MsgActionCreated{Action: act}, nil
}

func (k msgServer) NewSignTransactionRequestIntentGenerator(ctx sdk.Context, act intenttypes.Action) (intenttypes.Intent, error) {
	msg, err := intenttypes.GetActionMessage[*v1beta2.MsgNewSignTransactionRequest](k.cdc, act)
	if err != nil {
		return intenttypes.Intent{}, err
	}

	key, err := k.keys.Get(ctx, msg.KeyId)
	if err != nil {
		return intenttypes.Intent{}, err
	}

	ws, err := k.spaces.Get(ctx, key.SpaceId)
	if err != nil {
		return intenttypes.Intent{}, err
	}

	pol := ws.IntentNewSignTransactionRequest()
	return pol, nil
}

func (k msgServer) NewSignTransactionRequestActionHandler(ctx sdk.Context, act intenttypes.Action) (proto.Message, error) {
	msg, err := intenttypes.GetActionMessage[*v1beta2.MsgNewSignTransactionRequest](k.cdc, act)
	if err != nil {
		return nil, err
	}

	key, err := k.keys.Get(ctx, msg.KeyId)
	if err != nil {
		return nil, err
	}

	// use wallet to parse unsigned transaction
	w, err := v1beta2.NewWallet(&key, msg.WalletType)
	if err != nil {
		return nil, err
	}

	parser, ok := w.(v1beta2.TxParser)
	if !ok {
		return nil, fmt.Errorf("wallet does not implement TxParser")
	}

	var meta v1beta2.Metadata
	if err := k.cdc.UnpackAny(msg.Metadata, &meta); err != nil {
		return nil, fmt.Errorf("failed to unpack metadata: %w", err)
	}
	tx, err := parser.ParseTx(msg.UnsignedTransaction, meta)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tx: %w", err)
	}

	ctx.Logger().Debug("parsed layer 1 tx", "wallet", w, "tx", tx)

	// generate signature request
	signatureRequest := &v1beta2.SignRequest{
		Creator:        msg.Creator,
		KeyId:          msg.KeyId,
		KeyType:        key.Type,
		DataForSigning: tx.DataForSigning,
		Status:         v1beta2.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
	}
	signRequestID, err := k.signatureRequests.Append(ctx, signatureRequest)
	if err != nil {
		return nil, err
	}

	id, err := k.signTransactionRequests.Append(ctx, &v1beta2.SignTransactionRequest{
		Creator:             msg.Creator,
		SignRequestId:       signRequestID,
		KeyId:               msg.KeyId,
		WalletType:          msg.WalletType,
		UnsignedTransaction: msg.UnsignedTransaction,
	})
	if err != nil {
		return nil, err
	}

	return &v1beta2.MsgNewSignTransactionRequestResponse{Id: id, SignatureRequestId: signRequestID}, nil
}
