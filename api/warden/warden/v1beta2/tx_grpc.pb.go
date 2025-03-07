// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: warden/warden/v1beta2/tx.proto

package wardenv1beta2

import (
	context "context"
	intent "github.com/warden-protocol/wardenprotocol/api/warden/intent"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName              = "/warden.warden.v1beta2.Msg/UpdateParams"
	Msg_NewSpace_FullMethodName                  = "/warden.warden.v1beta2.Msg/NewSpace"
	Msg_AddSpaceOwner_FullMethodName             = "/warden.warden.v1beta2.Msg/AddSpaceOwner"
	Msg_RemoveSpaceOwner_FullMethodName          = "/warden.warden.v1beta2.Msg/RemoveSpaceOwner"
	Msg_NewKeychain_FullMethodName               = "/warden.warden.v1beta2.Msg/NewKeychain"
	Msg_AddKeychainParty_FullMethodName          = "/warden.warden.v1beta2.Msg/AddKeychainParty"
	Msg_UpdateSpace_FullMethodName               = "/warden.warden.v1beta2.Msg/UpdateSpace"
	Msg_UpdateKeychain_FullMethodName            = "/warden.warden.v1beta2.Msg/UpdateKeychain"
	Msg_NewKeyRequest_FullMethodName             = "/warden.warden.v1beta2.Msg/NewKeyRequest"
	Msg_UpdateKeyRequest_FullMethodName          = "/warden.warden.v1beta2.Msg/UpdateKeyRequest"
	Msg_NewSignatureRequest_FullMethodName       = "/warden.warden.v1beta2.Msg/NewSignatureRequest"
	Msg_FulfilSignatureRequest_FullMethodName    = "/warden.warden.v1beta2.Msg/FulfilSignatureRequest"
	Msg_NewSignTransactionRequest_FullMethodName = "/warden.warden.v1beta2.Msg/NewSignTransactionRequest"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// Create a new Space. The creator will be the first owner of the Space.
	NewSpace(ctx context.Context, in *MsgNewSpace, opts ...grpc.CallOption) (*MsgNewSpaceResponse, error)
	// Add a new owner to a space.
	AddSpaceOwner(ctx context.Context, in *MsgAddSpaceOwner, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
	// Remove an owner from the space. The user can remove itself, but at
	// least one owner must be left.
	RemoveSpaceOwner(ctx context.Context, in *MsgRemoveSpaceOwner, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
	// Create a new keychain. The user will be the first admin of the keychain.
	NewKeychain(ctx context.Context, in *MsgNewKeychain, opts ...grpc.CallOption) (*MsgNewKeychainResponse, error)
	// Add a new party to a keychain. Transactions coming from this party will
	// be considered trusted by the keychain.
	AddKeychainParty(ctx context.Context, in *MsgAddKeychainParty, opts ...grpc.CallOption) (*MsgAddKeychainPartyResponse, error)
	// Update a space, e.g. changing the intents in use.
	UpdateSpace(ctx context.Context, in *MsgUpdateSpace, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
	// Update a keychain, e.g. update the status or description.
	UpdateKeychain(ctx context.Context, in *MsgUpdateKeychain, opts ...grpc.CallOption) (*MsgUpdateKeychainResponse, error)
	// Request a new key to a keychain, the key will belong to the specified
	// space.
	NewKeyRequest(ctx context.Context, in *MsgNewKeyRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
	// Update an existing request by writing a result into it. This method is
	// called by a keychain party.
	UpdateKeyRequest(ctx context.Context, in *MsgUpdateKeyRequest, opts ...grpc.CallOption) (*MsgUpdateKeyRequestResponse, error)
	// Request a new signature
	NewSignatureRequest(ctx context.Context, in *MsgNewSignatureRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
	// Fulfill a signature request
	FulfilSignatureRequest(ctx context.Context, in *MsgFulfilSignatureRequest, opts ...grpc.CallOption) (*MsgFulfilSignatureRequestResponse, error)
	// Request a new signature for a layer 1 transaction, using the specified
	// wallet.
	// The difference with NewSignatureRequest is that this message will be
	// parsed by the wallet to apply specific intents that depends on
	// informations contained in the transaction itself (e.g. amount, recipient).
	NewSignTransactionRequest(ctx context.Context, in *MsgNewSignTransactionRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewSpace(ctx context.Context, in *MsgNewSpace, opts ...grpc.CallOption) (*MsgNewSpaceResponse, error) {
	out := new(MsgNewSpaceResponse)
	err := c.cc.Invoke(ctx, Msg_NewSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddSpaceOwner(ctx context.Context, in *MsgAddSpaceOwner, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_AddSpaceOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveSpaceOwner(ctx context.Context, in *MsgRemoveSpaceOwner, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_RemoveSpaceOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewKeychain(ctx context.Context, in *MsgNewKeychain, opts ...grpc.CallOption) (*MsgNewKeychainResponse, error) {
	out := new(MsgNewKeychainResponse)
	err := c.cc.Invoke(ctx, Msg_NewKeychain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddKeychainParty(ctx context.Context, in *MsgAddKeychainParty, opts ...grpc.CallOption) (*MsgAddKeychainPartyResponse, error) {
	out := new(MsgAddKeychainPartyResponse)
	err := c.cc.Invoke(ctx, Msg_AddKeychainParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSpace(ctx context.Context, in *MsgUpdateSpace, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_UpdateSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateKeychain(ctx context.Context, in *MsgUpdateKeychain, opts ...grpc.CallOption) (*MsgUpdateKeychainResponse, error) {
	out := new(MsgUpdateKeychainResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateKeychain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewKeyRequest(ctx context.Context, in *MsgNewKeyRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_NewKeyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateKeyRequest(ctx context.Context, in *MsgUpdateKeyRequest, opts ...grpc.CallOption) (*MsgUpdateKeyRequestResponse, error) {
	out := new(MsgUpdateKeyRequestResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateKeyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewSignatureRequest(ctx context.Context, in *MsgNewSignatureRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_NewSignatureRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FulfilSignatureRequest(ctx context.Context, in *MsgFulfilSignatureRequest, opts ...grpc.CallOption) (*MsgFulfilSignatureRequestResponse, error) {
	out := new(MsgFulfilSignatureRequestResponse)
	err := c.cc.Invoke(ctx, Msg_FulfilSignatureRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewSignTransactionRequest(ctx context.Context, in *MsgNewSignTransactionRequest, opts ...grpc.CallOption) (*intent.MsgActionCreated, error) {
	out := new(intent.MsgActionCreated)
	err := c.cc.Invoke(ctx, Msg_NewSignTransactionRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// Create a new Space. The creator will be the first owner of the Space.
	NewSpace(context.Context, *MsgNewSpace) (*MsgNewSpaceResponse, error)
	// Add a new owner to a space.
	AddSpaceOwner(context.Context, *MsgAddSpaceOwner) (*intent.MsgActionCreated, error)
	// Remove an owner from the space. The user can remove itself, but at
	// least one owner must be left.
	RemoveSpaceOwner(context.Context, *MsgRemoveSpaceOwner) (*intent.MsgActionCreated, error)
	// Create a new keychain. The user will be the first admin of the keychain.
	NewKeychain(context.Context, *MsgNewKeychain) (*MsgNewKeychainResponse, error)
	// Add a new party to a keychain. Transactions coming from this party will
	// be considered trusted by the keychain.
	AddKeychainParty(context.Context, *MsgAddKeychainParty) (*MsgAddKeychainPartyResponse, error)
	// Update a space, e.g. changing the intents in use.
	UpdateSpace(context.Context, *MsgUpdateSpace) (*intent.MsgActionCreated, error)
	// Update a keychain, e.g. update the status or description.
	UpdateKeychain(context.Context, *MsgUpdateKeychain) (*MsgUpdateKeychainResponse, error)
	// Request a new key to a keychain, the key will belong to the specified
	// space.
	NewKeyRequest(context.Context, *MsgNewKeyRequest) (*intent.MsgActionCreated, error)
	// Update an existing request by writing a result into it. This method is
	// called by a keychain party.
	UpdateKeyRequest(context.Context, *MsgUpdateKeyRequest) (*MsgUpdateKeyRequestResponse, error)
	// Request a new signature
	NewSignatureRequest(context.Context, *MsgNewSignatureRequest) (*intent.MsgActionCreated, error)
	// Fulfill a signature request
	FulfilSignatureRequest(context.Context, *MsgFulfilSignatureRequest) (*MsgFulfilSignatureRequestResponse, error)
	// Request a new signature for a layer 1 transaction, using the specified
	// wallet.
	// The difference with NewSignatureRequest is that this message will be
	// parsed by the wallet to apply specific intents that depends on
	// informations contained in the transaction itself (e.g. amount, recipient).
	NewSignTransactionRequest(context.Context, *MsgNewSignTransactionRequest) (*intent.MsgActionCreated, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) NewSpace(context.Context, *MsgNewSpace) (*MsgNewSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSpace not implemented")
}
func (UnimplementedMsgServer) AddSpaceOwner(context.Context, *MsgAddSpaceOwner) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSpaceOwner not implemented")
}
func (UnimplementedMsgServer) RemoveSpaceOwner(context.Context, *MsgRemoveSpaceOwner) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSpaceOwner not implemented")
}
func (UnimplementedMsgServer) NewKeychain(context.Context, *MsgNewKeychain) (*MsgNewKeychainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewKeychain not implemented")
}
func (UnimplementedMsgServer) AddKeychainParty(context.Context, *MsgAddKeychainParty) (*MsgAddKeychainPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKeychainParty not implemented")
}
func (UnimplementedMsgServer) UpdateSpace(context.Context, *MsgUpdateSpace) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
func (UnimplementedMsgServer) UpdateKeychain(context.Context, *MsgUpdateKeychain) (*MsgUpdateKeychainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeychain not implemented")
}
func (UnimplementedMsgServer) NewKeyRequest(context.Context, *MsgNewKeyRequest) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewKeyRequest not implemented")
}
func (UnimplementedMsgServer) UpdateKeyRequest(context.Context, *MsgUpdateKeyRequest) (*MsgUpdateKeyRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeyRequest not implemented")
}
func (UnimplementedMsgServer) NewSignatureRequest(context.Context, *MsgNewSignatureRequest) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSignatureRequest not implemented")
}
func (UnimplementedMsgServer) FulfilSignatureRequest(context.Context, *MsgFulfilSignatureRequest) (*MsgFulfilSignatureRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FulfilSignatureRequest not implemented")
}
func (UnimplementedMsgServer) NewSignTransactionRequest(context.Context, *MsgNewSignTransactionRequest) (*intent.MsgActionCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSignTransactionRequest not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewSpace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewSpace(ctx, req.(*MsgNewSpace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddSpaceOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddSpaceOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddSpaceOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddSpaceOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddSpaceOwner(ctx, req.(*MsgAddSpaceOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveSpaceOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveSpaceOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveSpaceOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveSpaceOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveSpaceOwner(ctx, req.(*MsgRemoveSpaceOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewKeychain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewKeychain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewKeychain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewKeychain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewKeychain(ctx, req.(*MsgNewKeychain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddKeychainParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddKeychainParty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddKeychainParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddKeychainParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddKeychainParty(ctx, req.(*MsgAddKeychainParty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSpace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSpace(ctx, req.(*MsgUpdateSpace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateKeychain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateKeychain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateKeychain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateKeychain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateKeychain(ctx, req.(*MsgUpdateKeychain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewKeyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewKeyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewKeyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewKeyRequest(ctx, req.(*MsgNewKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateKeyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateKeyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateKeyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateKeyRequest(ctx, req.(*MsgUpdateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewSignatureRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewSignatureRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewSignatureRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewSignatureRequest(ctx, req.(*MsgNewSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FulfilSignatureRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFulfilSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FulfilSignatureRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FulfilSignatureRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FulfilSignatureRequest(ctx, req.(*MsgFulfilSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewSignTransactionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewSignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewSignTransactionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewSignTransactionRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewSignTransactionRequest(ctx, req.(*MsgNewSignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warden.warden.v1beta2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "NewSpace",
			Handler:    _Msg_NewSpace_Handler,
		},
		{
			MethodName: "AddSpaceOwner",
			Handler:    _Msg_AddSpaceOwner_Handler,
		},
		{
			MethodName: "RemoveSpaceOwner",
			Handler:    _Msg_RemoveSpaceOwner_Handler,
		},
		{
			MethodName: "NewKeychain",
			Handler:    _Msg_NewKeychain_Handler,
		},
		{
			MethodName: "AddKeychainParty",
			Handler:    _Msg_AddKeychainParty_Handler,
		},
		{
			MethodName: "UpdateSpace",
			Handler:    _Msg_UpdateSpace_Handler,
		},
		{
			MethodName: "UpdateKeychain",
			Handler:    _Msg_UpdateKeychain_Handler,
		},
		{
			MethodName: "NewKeyRequest",
			Handler:    _Msg_NewKeyRequest_Handler,
		},
		{
			MethodName: "UpdateKeyRequest",
			Handler:    _Msg_UpdateKeyRequest_Handler,
		},
		{
			MethodName: "NewSignatureRequest",
			Handler:    _Msg_NewSignatureRequest_Handler,
		},
		{
			MethodName: "FulfilSignatureRequest",
			Handler:    _Msg_FulfilSignatureRequest_Handler,
		},
		{
			MethodName: "NewSignTransactionRequest",
			Handler:    _Msg_NewSignTransactionRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/warden/v1beta2/tx.proto",
}
