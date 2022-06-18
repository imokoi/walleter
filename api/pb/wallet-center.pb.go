// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: wallet-center.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameClient int32

const (
	GameClient_NecoFishing GameClient = 0
	GameClient_NecoLand    GameClient = 1
)

// Enum value maps for GameClient.
var (
	GameClient_name = map[int32]string{
		0: "NecoFishing",
		1: "NecoLand",
	}
	GameClient_value = map[string]int32{
		"NecoFishing": 0,
		"NecoLand":    1,
	}
)

func (x GameClient) Enum() *GameClient {
	p := new(GameClient)
	*p = x
	return p
}

func (x GameClient) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameClient) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_center_proto_enumTypes[0].Descriptor()
}

func (GameClient) Type() protoreflect.EnumType {
	return &file_wallet_center_proto_enumTypes[0]
}

func (x GameClient) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameClient.Descriptor instead.
func (GameClient) EnumDescriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{0}
}

type ERC20Token int32

const (
	ERC20Token_NFISH ERC20Token = 0
	ERC20Token_BUSD  ERC20Token = 1
)

// Enum value maps for ERC20Token.
var (
	ERC20Token_name = map[int32]string{
		0: "NFISH",
		1: "BUSD",
	}
	ERC20Token_value = map[string]int32{
		"NFISH": 0,
		"BUSD":  1,
	}
)

func (x ERC20Token) Enum() *ERC20Token {
	p := new(ERC20Token)
	*p = x
	return p
}

func (x ERC20Token) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERC20Token) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_center_proto_enumTypes[1].Descriptor()
}

func (ERC20Token) Type() protoreflect.EnumType {
	return &file_wallet_center_proto_enumTypes[1]
}

func (x ERC20Token) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERC20Token.Descriptor instead.
func (ERC20Token) EnumDescriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{1}
}

type WalletActionType int32

const (
	WalletActionType_Initialize WalletActionType = 0
	WalletActionType_Income     WalletActionType = 1
	WalletActionType_Spend      WalletActionType = 2
	WalletActionType_Deposit    WalletActionType = 3
	WalletActionType_Withdraw   WalletActionType = 4
	WalletActionType_ChargeFee  WalletActionType = 5
)

// Enum value maps for WalletActionType.
var (
	WalletActionType_name = map[int32]string{
		0: "Initialize",
		1: "Income",
		2: "Spend",
		3: "Deposit",
		4: "Withdraw",
		5: "ChargeFee",
	}
	WalletActionType_value = map[string]int32{
		"Initialize": 0,
		"Income":     1,
		"Spend":      2,
		"Deposit":    3,
		"Withdraw":   4,
		"ChargeFee":  5,
	}
)

func (x WalletActionType) Enum() *WalletActionType {
	p := new(WalletActionType)
	*p = x
	return p
}

func (x WalletActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WalletActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_center_proto_enumTypes[2].Descriptor()
}

func (WalletActionType) Type() protoreflect.EnumType {
	return &file_wallet_center_proto_enumTypes[2]
}

func (x WalletActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WalletActionType.Descriptor instead.
func (WalletActionType) EnumDescriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{2}
}

type ActionStatus int32

const (
	ActionStatus_Pending ActionStatus = 0
	ActionStatus_Succeed ActionStatus = 1
)

// Enum value maps for ActionStatus.
var (
	ActionStatus_name = map[int32]string{
		0: "Pending",
		1: "Succeed",
	}
	ActionStatus_value = map[string]int32{
		"Pending": 0,
		"Succeed": 1,
	}
)

func (x ActionStatus) Enum() *ActionStatus {
	p := new(ActionStatus)
	*p = x
	return p
}

func (x ActionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_center_proto_enumTypes[3].Descriptor()
}

func (ActionStatus) Type() protoreflect.EnumType {
	return &file_wallet_center_proto_enumTypes[3]
}

func (x ActionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionStatus.Descriptor instead.
func (ActionStatus) EnumDescriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{3}
}

type AssetType int32

const (
	AssetType_ERC20AssetType   AssetType = 0
	AssetType_ERC1155AssetType AssetType = 1
	AssetType_Other            AssetType = 2
)

// Enum value maps for AssetType.
var (
	AssetType_name = map[int32]string{
		0: "ERC20AssetType",
		1: "ERC1155AssetType",
		2: "Other",
	}
	AssetType_value = map[string]int32{
		"ERC20AssetType":   0,
		"ERC1155AssetType": 1,
		"Other":            2,
	}
)

func (x AssetType) Enum() *AssetType {
	p := new(AssetType)
	*p = x
	return p
}

func (x AssetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetType) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_center_proto_enumTypes[4].Descriptor()
}

func (AssetType) Type() protoreflect.EnumType {
	return &file_wallet_center_proto_enumTypes[4]
}

func (x AssetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetType.Descriptor instead.
func (AssetType) EnumDescriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{4}
}

type UpdateUserWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId        uint64            `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GameClient       GameClient        `protobuf:"varint,2,opt,name=GameClient,proto3,enum=GameClient" json:"GameClient,omitempty"`
	BusinessModule   string            `protobuf:"bytes,3,opt,name=BusinessModule,proto3" json:"BusinessModule,omitempty"`
	AssetType        AssetType         `protobuf:"varint,4,opt,name=AssetType,proto3,enum=AssetType" json:"AssetType,omitempty"`
	ActionType       WalletActionType  `protobuf:"varint,5,opt,name=ActionType,proto3,enum=WalletActionType" json:"ActionType,omitempty"`
	ERC20TokenData   []*ERC20TokenData `protobuf:"bytes,6,rep,name=ERC20TokenData,proto3" json:"ERC20TokenData,omitempty"`
	ERC1155TokenData *ERC1155TokenData `protobuf:"bytes,7,opt,name=ERC1155TokenData,proto3" json:"ERC1155TokenData,omitempty"`
	FeeData          []*ERC20TokenData `protobuf:"bytes,8,rep,name=FeeData,proto3" json:"FeeData,omitempty"`
}

func (x *UpdateUserWalletRequest) Reset() {
	*x = UpdateUserWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserWalletRequest) ProtoMessage() {}

func (x *UpdateUserWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserWalletRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserWalletRequest) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *UpdateUserWalletRequest) GetGameClient() GameClient {
	if x != nil {
		return x.GameClient
	}
	return GameClient_NecoFishing
}

func (x *UpdateUserWalletRequest) GetBusinessModule() string {
	if x != nil {
		return x.BusinessModule
	}
	return ""
}

func (x *UpdateUserWalletRequest) GetAssetType() AssetType {
	if x != nil {
		return x.AssetType
	}
	return AssetType_ERC20AssetType
}

func (x *UpdateUserWalletRequest) GetActionType() WalletActionType {
	if x != nil {
		return x.ActionType
	}
	return WalletActionType_Initialize
}

func (x *UpdateUserWalletRequest) GetERC20TokenData() []*ERC20TokenData {
	if x != nil {
		return x.ERC20TokenData
	}
	return nil
}

func (x *UpdateUserWalletRequest) GetERC1155TokenData() *ERC1155TokenData {
	if x != nil {
		return x.ERC1155TokenData
	}
	return nil
}

func (x *UpdateUserWalletRequest) GetFeeData() []*ERC20TokenData {
	if x != nil {
		return x.FeeData
	}
	return nil
}

type GetUserWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  uint64     `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GameClient GameClient `protobuf:"varint,2,opt,name=GameClient,proto3,enum=GameClient" json:"GameClient,omitempty"`
}

func (x *GetUserWalletRequest) Reset() {
	*x = GetUserWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWalletRequest) ProtoMessage() {}

func (x *GetUserWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWalletRequest.ProtoReflect.Descriptor instead.
func (*GetUserWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserWalletRequest) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *GetUserWalletRequest) GetGameClient() GameClient {
	if x != nil {
		return x.GameClient
	}
	return GameClient_NecoFishing
}

type UserWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64            `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GameClient    GameClient        `protobuf:"varint,2,opt,name=GameClient,proto3,enum=GameClient" json:"GameClient,omitempty"`
	AccountId     uint64            `protobuf:"varint,3,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	PublicAddress string            `protobuf:"bytes,4,opt,name=PublicAddress,proto3" json:"PublicAddress,omitempty"`
	ERC20Tokens   []*ERC20TokenData `protobuf:"bytes,5,rep,name=ERC20Tokens,proto3" json:"ERC20Tokens,omitempty"`
	ERC1155Token  *ERC1155TokenData `protobuf:"bytes,6,opt,name=ERC1155Token,proto3" json:"ERC1155Token,omitempty"`
}

func (x *UserWallet) Reset() {
	*x = UserWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWallet) ProtoMessage() {}

func (x *UserWallet) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWallet.ProtoReflect.Descriptor instead.
func (*UserWallet) Descriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{2}
}

func (x *UserWallet) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserWallet) GetGameClient() GameClient {
	if x != nil {
		return x.GameClient
	}
	return GameClient_NecoFishing
}

func (x *UserWallet) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *UserWallet) GetPublicAddress() string {
	if x != nil {
		return x.PublicAddress
	}
	return ""
}

func (x *UserWallet) GetERC20Tokens() []*ERC20TokenData {
	if x != nil {
		return x.ERC20Tokens
	}
	return nil
}

func (x *UserWallet) GetERC1155Token() *ERC1155TokenData {
	if x != nil {
		return x.ERC1155Token
	}
	return nil
}

type ERC1155TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Ids    []uint64 `protobuf:"varint,2,rep,packed,name=Ids,proto3" json:"Ids,omitempty"`
	Values []uint64 `protobuf:"varint,3,rep,packed,name=Values,proto3" json:"Values,omitempty"`
}

func (x *ERC1155TokenData) Reset() {
	*x = ERC1155TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC1155TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC1155TokenData) ProtoMessage() {}

func (x *ERC1155TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC1155TokenData.ProtoReflect.Descriptor instead.
func (*ERC1155TokenData) Descriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{3}
}

func (x *ERC1155TokenData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ERC1155TokenData) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ERC1155TokenData) GetValues() []uint64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type ERC20TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Token   ERC20Token `protobuf:"varint,2,opt,name=Token,proto3,enum=ERC20Token" json:"Token,omitempty"`
	Balance float32    `protobuf:"fixed32,3,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Decimal uint64     `protobuf:"varint,4,opt,name=Decimal,proto3" json:"Decimal,omitempty"`
}

func (x *ERC20TokenData) Reset() {
	*x = ERC20TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC20TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC20TokenData) ProtoMessage() {}

func (x *ERC20TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC20TokenData.ProtoReflect.Descriptor instead.
func (*ERC20TokenData) Descriptor() ([]byte, []int) {
	return file_wallet_center_proto_rawDescGZIP(), []int{4}
}

func (x *ERC20TokenData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ERC20TokenData) GetToken() ERC20Token {
	if x != nil {
		return x.Token
	}
	return ERC20Token_NFISH
}

func (x *ERC20TokenData) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ERC20TokenData) GetDecimal() uint64 {
	if x != nil {
		return x.Decimal
	}
	return 0
}

var File_wallet_center_proto protoreflect.FileDescriptor

var file_wallet_center_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x37, 0x0a, 0x0e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x52, 0x43, 0x32,
	0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x45, 0x52, 0x43, 0x32,
	0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x10, 0x45, 0x52,
	0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x07, 0x46, 0x65, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x52, 0x43,
	0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x46, 0x65, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x61,
	0x6d, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x52, 0x43, 0x32, 0x30,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45,
	0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x45,
	0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x45, 0x52,
	0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x4c, 0x0a, 0x10, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x03, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0x77, 0x0a, 0x0e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2a, 0x2b, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x65, 0x63, 0x6f, 0x46, 0x69,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x65, 0x63, 0x6f, 0x4c,
	0x61, 0x6e, 0x64, 0x10, 0x01, 0x2a, 0x21, 0x0a, 0x0a, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x46, 0x49, 0x53, 0x48, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x55, 0x53, 0x44, 0x10, 0x01, 0x2a, 0x63, 0x0a, 0x10, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x46, 0x65, 0x65, 0x10, 0x05, 0x2a, 0x28, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x10, 0x01, 0x2a, 0x40, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x43, 0x31,
	0x31, 0x35, 0x35, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x10, 0x02, 0x32, 0x86, 0x01, 0x0a, 0x10, 0x4e, 0x65,
	0x63, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wallet_center_proto_rawDescOnce sync.Once
	file_wallet_center_proto_rawDescData = file_wallet_center_proto_rawDesc
)

func file_wallet_center_proto_rawDescGZIP() []byte {
	file_wallet_center_proto_rawDescOnce.Do(func() {
		file_wallet_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_center_proto_rawDescData)
	})
	return file_wallet_center_proto_rawDescData
}

var file_wallet_center_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_wallet_center_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wallet_center_proto_goTypes = []interface{}{
	(GameClient)(0),                 // 0: GameClient
	(ERC20Token)(0),                 // 1: ERC20Token
	(WalletActionType)(0),           // 2: WalletActionType
	(ActionStatus)(0),               // 3: ActionStatus
	(AssetType)(0),                  // 4: AssetType
	(*UpdateUserWalletRequest)(nil), // 5: UpdateUserWalletRequest
	(*GetUserWalletRequest)(nil),    // 6: GetUserWalletRequest
	(*UserWallet)(nil),              // 7: UserWallet
	(*ERC1155TokenData)(nil),        // 8: ERC1155TokenData
	(*ERC20TokenData)(nil),          // 9: ERC20TokenData
}
var file_wallet_center_proto_depIdxs = []int32{
	0,  // 0: UpdateUserWalletRequest.GameClient:type_name -> GameClient
	4,  // 1: UpdateUserWalletRequest.AssetType:type_name -> AssetType
	2,  // 2: UpdateUserWalletRequest.ActionType:type_name -> WalletActionType
	9,  // 3: UpdateUserWalletRequest.ERC20TokenData:type_name -> ERC20TokenData
	8,  // 4: UpdateUserWalletRequest.ERC1155TokenData:type_name -> ERC1155TokenData
	9,  // 5: UpdateUserWalletRequest.FeeData:type_name -> ERC20TokenData
	0,  // 6: GetUserWalletRequest.GameClient:type_name -> GameClient
	0,  // 7: UserWallet.GameClient:type_name -> GameClient
	9,  // 8: UserWallet.ERC20Tokens:type_name -> ERC20TokenData
	8,  // 9: UserWallet.ERC1155Token:type_name -> ERC1155TokenData
	1,  // 10: ERC20TokenData.Token:type_name -> ERC20Token
	5,  // 11: NecoWalletCenter.UpdateUserWallet:input_type -> UpdateUserWalletRequest
	6,  // 12: NecoWalletCenter.GetUserWallet:input_type -> GetUserWalletRequest
	7,  // 13: NecoWalletCenter.UpdateUserWallet:output_type -> UserWallet
	7,  // 14: NecoWalletCenter.GetUserWallet:output_type -> UserWallet
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_wallet_center_proto_init() }
func file_wallet_center_proto_init() {
	if File_wallet_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWallet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC1155TokenData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_center_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC20TokenData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_center_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_center_proto_goTypes,
		DependencyIndexes: file_wallet_center_proto_depIdxs,
		EnumInfos:         file_wallet_center_proto_enumTypes,
		MessageInfos:      file_wallet_center_proto_msgTypes,
	}.Build()
	File_wallet_center_proto = out.File
	file_wallet_center_proto_rawDesc = nil
	file_wallet_center_proto_goTypes = nil
	file_wallet_center_proto_depIdxs = nil
}
