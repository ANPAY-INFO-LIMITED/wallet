// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: invoicesrpc/invoices.proto

package invoicesrpc

import (
	lnrpc "github.com/lightningnetwork/lnd/lnrpc"
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

type LookupModifier int32

const (
	// The default look up modifier, no look up behavior is changed.
	LookupModifier_DEFAULT LookupModifier = 0
	// Indicates that when a look up is done based on a set_id, then only that set
	// of HTLCs related to that set ID should be returned.
	LookupModifier_HTLC_SET_ONLY LookupModifier = 1
	// Indicates that when a look up is done using a payment_addr, then no HTLCs
	// related to the payment_addr should be returned. This is useful when one
	// wants to be able to obtain the set of associated setIDs with a given
	// invoice, then look up the sub-invoices "projected" by that set ID.
	LookupModifier_HTLC_SET_BLANK LookupModifier = 2
)

// Enum value maps for LookupModifier.
var (
	LookupModifier_name = map[int32]string{
		0: "DEFAULT",
		1: "HTLC_SET_ONLY",
		2: "HTLC_SET_BLANK",
	}
	LookupModifier_value = map[string]int32{
		"DEFAULT":        0,
		"HTLC_SET_ONLY":  1,
		"HTLC_SET_BLANK": 2,
	}
)

func (x LookupModifier) Enum() *LookupModifier {
	p := new(LookupModifier)
	*p = x
	return p
}

func (x LookupModifier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LookupModifier) Descriptor() protoreflect.EnumDescriptor {
	return file_invoicesrpc_invoices_proto_enumTypes[0].Descriptor()
}

func (LookupModifier) Type() protoreflect.EnumType {
	return &file_invoicesrpc_invoices_proto_enumTypes[0]
}

func (x LookupModifier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LookupModifier.Descriptor instead.
func (LookupModifier) EnumDescriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{0}
}

type CancelInvoiceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash corresponding to the (hold) invoice to cancel. When using
	// REST, this field must be encoded as base64.
	PaymentHash []byte `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3" json:"payment_hash,omitempty"`
}

func (x *CancelInvoiceMsg) Reset() {
	*x = CancelInvoiceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelInvoiceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInvoiceMsg) ProtoMessage() {}

func (x *CancelInvoiceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInvoiceMsg.ProtoReflect.Descriptor instead.
func (*CancelInvoiceMsg) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{0}
}

func (x *CancelInvoiceMsg) GetPaymentHash() []byte {
	if x != nil {
		return x.PaymentHash
	}
	return nil
}

type CancelInvoiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelInvoiceResp) Reset() {
	*x = CancelInvoiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelInvoiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInvoiceResp) ProtoMessage() {}

func (x *CancelInvoiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInvoiceResp.ProtoReflect.Descriptor instead.
func (*CancelInvoiceResp) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{1}
}

type AddHoldInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional memo to attach along with the invoice. Used for record keeping
	// purposes for the invoice's creator, and will also be set in the description
	// field of the encoded payment request if the description_hash field is not
	// being used.
	Memo string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	// The hash of the preimage
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The value of this invoice in satoshis
	//
	// The fields value and value_msat are mutually exclusive.
	Value int64 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	// The value of this invoice in millisatoshis
	//
	// The fields value and value_msat are mutually exclusive.
	ValueMsat int64 `protobuf:"varint,10,opt,name=value_msat,json=valueMsat,proto3" json:"value_msat,omitempty"`
	// Hash (SHA-256) of a description of the payment. Used if the description of
	// payment (memo) is too long to naturally fit within the description field
	// of an encoded payment request.
	DescriptionHash []byte `protobuf:"bytes,4,opt,name=description_hash,json=descriptionHash,proto3" json:"description_hash,omitempty"`
	// Payment request expiry time in seconds. Default is 86400 (24 hours).
	Expiry int64 `protobuf:"varint,5,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// Fallback on-chain address.
	FallbackAddr string `protobuf:"bytes,6,opt,name=fallback_addr,json=fallbackAddr,proto3" json:"fallback_addr,omitempty"`
	// Delta to use for the time-lock of the CLTV extended to the final hop.
	CltvExpiry uint64 `protobuf:"varint,7,opt,name=cltv_expiry,json=cltvExpiry,proto3" json:"cltv_expiry,omitempty"`
	// Route hints that can each be individually used to assist in reaching the
	// invoice's destination.
	RouteHints []*lnrpc.RouteHint `protobuf:"bytes,8,rep,name=route_hints,json=routeHints,proto3" json:"route_hints,omitempty"`
	// Whether this invoice should include routing hints for private channels.
	Private bool `protobuf:"varint,9,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *AddHoldInvoiceRequest) Reset() {
	*x = AddHoldInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHoldInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHoldInvoiceRequest) ProtoMessage() {}

func (x *AddHoldInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHoldInvoiceRequest.ProtoReflect.Descriptor instead.
func (*AddHoldInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{2}
}

func (x *AddHoldInvoiceRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *AddHoldInvoiceRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *AddHoldInvoiceRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AddHoldInvoiceRequest) GetValueMsat() int64 {
	if x != nil {
		return x.ValueMsat
	}
	return 0
}

func (x *AddHoldInvoiceRequest) GetDescriptionHash() []byte {
	if x != nil {
		return x.DescriptionHash
	}
	return nil
}

func (x *AddHoldInvoiceRequest) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *AddHoldInvoiceRequest) GetFallbackAddr() string {
	if x != nil {
		return x.FallbackAddr
	}
	return ""
}

func (x *AddHoldInvoiceRequest) GetCltvExpiry() uint64 {
	if x != nil {
		return x.CltvExpiry
	}
	return 0
}

func (x *AddHoldInvoiceRequest) GetRouteHints() []*lnrpc.RouteHint {
	if x != nil {
		return x.RouteHints
	}
	return nil
}

func (x *AddHoldInvoiceRequest) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type AddHoldInvoiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A bare-bones invoice for a payment within the Lightning Network. With the
	// details of the invoice, the sender has all the data necessary to send a
	// payment to the recipient.
	PaymentRequest string `protobuf:"bytes,1,opt,name=payment_request,json=paymentRequest,proto3" json:"payment_request,omitempty"`
	// The "add" index of this invoice. Each newly created invoice will increment
	// this index making it monotonically increasing. Callers to the
	// SubscribeInvoices call can use this to instantly get notified of all added
	// invoices with an add_index greater than this one.
	AddIndex uint64 `protobuf:"varint,2,opt,name=add_index,json=addIndex,proto3" json:"add_index,omitempty"`
	// The payment address of the generated invoice. This value should be used
	// in all payments for this invoice as we require it for end to end
	// security.
	PaymentAddr []byte `protobuf:"bytes,3,opt,name=payment_addr,json=paymentAddr,proto3" json:"payment_addr,omitempty"`
}

func (x *AddHoldInvoiceResp) Reset() {
	*x = AddHoldInvoiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHoldInvoiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHoldInvoiceResp) ProtoMessage() {}

func (x *AddHoldInvoiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHoldInvoiceResp.ProtoReflect.Descriptor instead.
func (*AddHoldInvoiceResp) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{3}
}

func (x *AddHoldInvoiceResp) GetPaymentRequest() string {
	if x != nil {
		return x.PaymentRequest
	}
	return ""
}

func (x *AddHoldInvoiceResp) GetAddIndex() uint64 {
	if x != nil {
		return x.AddIndex
	}
	return 0
}

func (x *AddHoldInvoiceResp) GetPaymentAddr() []byte {
	if x != nil {
		return x.PaymentAddr
	}
	return nil
}

type SettleInvoiceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Externally discovered pre-image that should be used to settle the hold
	// invoice.
	Preimage []byte `protobuf:"bytes,1,opt,name=preimage,proto3" json:"preimage,omitempty"`
}

func (x *SettleInvoiceMsg) Reset() {
	*x = SettleInvoiceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleInvoiceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleInvoiceMsg) ProtoMessage() {}

func (x *SettleInvoiceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleInvoiceMsg.ProtoReflect.Descriptor instead.
func (*SettleInvoiceMsg) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{4}
}

func (x *SettleInvoiceMsg) GetPreimage() []byte {
	if x != nil {
		return x.Preimage
	}
	return nil
}

type SettleInvoiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SettleInvoiceResp) Reset() {
	*x = SettleInvoiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleInvoiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleInvoiceResp) ProtoMessage() {}

func (x *SettleInvoiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleInvoiceResp.ProtoReflect.Descriptor instead.
func (*SettleInvoiceResp) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{5}
}

type SubscribeSingleInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash corresponding to the (hold) invoice to subscribe to. When using
	// REST, this field must be encoded as base64url.
	RHash []byte `protobuf:"bytes,2,opt,name=r_hash,json=rHash,proto3" json:"r_hash,omitempty"`
}

func (x *SubscribeSingleInvoiceRequest) Reset() {
	*x = SubscribeSingleInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeSingleInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeSingleInvoiceRequest) ProtoMessage() {}

func (x *SubscribeSingleInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeSingleInvoiceRequest.ProtoReflect.Descriptor instead.
func (*SubscribeSingleInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeSingleInvoiceRequest) GetRHash() []byte {
	if x != nil {
		return x.RHash
	}
	return nil
}

type LookupInvoiceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to InvoiceRef:
	//
	//	*LookupInvoiceMsg_PaymentHash
	//	*LookupInvoiceMsg_PaymentAddr
	//	*LookupInvoiceMsg_SetId
	InvoiceRef     isLookupInvoiceMsg_InvoiceRef `protobuf_oneof:"invoice_ref"`
	LookupModifier LookupModifier                `protobuf:"varint,4,opt,name=lookup_modifier,json=lookupModifier,proto3,enum=invoicesrpc.LookupModifier" json:"lookup_modifier,omitempty"`
}

func (x *LookupInvoiceMsg) Reset() {
	*x = LookupInvoiceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicesrpc_invoices_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupInvoiceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupInvoiceMsg) ProtoMessage() {}

func (x *LookupInvoiceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_invoicesrpc_invoices_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupInvoiceMsg.ProtoReflect.Descriptor instead.
func (*LookupInvoiceMsg) Descriptor() ([]byte, []int) {
	return file_invoicesrpc_invoices_proto_rawDescGZIP(), []int{7}
}

func (m *LookupInvoiceMsg) GetInvoiceRef() isLookupInvoiceMsg_InvoiceRef {
	if m != nil {
		return m.InvoiceRef
	}
	return nil
}

func (x *LookupInvoiceMsg) GetPaymentHash() []byte {
	if x, ok := x.GetInvoiceRef().(*LookupInvoiceMsg_PaymentHash); ok {
		return x.PaymentHash
	}
	return nil
}

func (x *LookupInvoiceMsg) GetPaymentAddr() []byte {
	if x, ok := x.GetInvoiceRef().(*LookupInvoiceMsg_PaymentAddr); ok {
		return x.PaymentAddr
	}
	return nil
}

func (x *LookupInvoiceMsg) GetSetId() []byte {
	if x, ok := x.GetInvoiceRef().(*LookupInvoiceMsg_SetId); ok {
		return x.SetId
	}
	return nil
}

func (x *LookupInvoiceMsg) GetLookupModifier() LookupModifier {
	if x != nil {
		return x.LookupModifier
	}
	return LookupModifier_DEFAULT
}

type isLookupInvoiceMsg_InvoiceRef interface {
	isLookupInvoiceMsg_InvoiceRef()
}

type LookupInvoiceMsg_PaymentHash struct {
	// When using REST, this field must be encoded as base64.
	PaymentHash []byte `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3,oneof"`
}

type LookupInvoiceMsg_PaymentAddr struct {
	PaymentAddr []byte `protobuf:"bytes,2,opt,name=payment_addr,json=paymentAddr,proto3,oneof"`
}

type LookupInvoiceMsg_SetId struct {
	SetId []byte `protobuf:"bytes,3,opt,name=set_id,json=setId,proto3,oneof"`
}

func (*LookupInvoiceMsg_PaymentHash) isLookupInvoiceMsg_InvoiceRef() {}

func (*LookupInvoiceMsg_PaymentAddr) isLookupInvoiceMsg_InvoiceRef() {}

func (*LookupInvoiceMsg_SetId) isLookupInvoiceMsg_InvoiceRef() {}

var File_invoicesrpc_invoices_proto protoreflect.FileDescriptor

var file_invoicesrpc_invoices_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x1a, 0x0f, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x10, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xca, 0x02, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x48, 0x6f,
	0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x73, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x73, 0x61, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x74, 0x76, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6c, 0x74, 0x76,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x68, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x22, 0x7d, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x64, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x3c, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x48, 0x61, 0x73, 0x68, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0xca, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x23, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x06, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a,
	0x0f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x2a, 0x44, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x54, 0x4c, 0x43, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x4f, 0x4e,
	0x4c, 0x59, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x54, 0x4c, 0x43, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x42, 0x4c, 0x41, 0x4e, 0x4b, 0x10, 0x02, 0x32, 0x9b, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x2a, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x30, 0x01, 0x12, 0x4e, 0x0a,
	0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x1e, 0x2e,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x55, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x22, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64,
	0x64, 0x48, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x4d, 0x73, 0x67, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x32, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x0e, 0x2e, 0x6c, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6e, 0x64, 0x2f, 0x6c, 0x6e, 0x72, 0x70, 0x63, 0x2f,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_invoicesrpc_invoices_proto_rawDescOnce sync.Once
	file_invoicesrpc_invoices_proto_rawDescData = file_invoicesrpc_invoices_proto_rawDesc
)

func file_invoicesrpc_invoices_proto_rawDescGZIP() []byte {
	file_invoicesrpc_invoices_proto_rawDescOnce.Do(func() {
		file_invoicesrpc_invoices_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoicesrpc_invoices_proto_rawDescData)
	})
	return file_invoicesrpc_invoices_proto_rawDescData
}

var file_invoicesrpc_invoices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_invoicesrpc_invoices_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_invoicesrpc_invoices_proto_goTypes = []interface{}{
	(LookupModifier)(0),                   // 0: invoicesrpc.LookupModifier
	(*CancelInvoiceMsg)(nil),              // 1: invoicesrpc.CancelInvoiceMsg
	(*CancelInvoiceResp)(nil),             // 2: invoicesrpc.CancelInvoiceResp
	(*AddHoldInvoiceRequest)(nil),         // 3: invoicesrpc.AddHoldInvoiceRequest
	(*AddHoldInvoiceResp)(nil),            // 4: invoicesrpc.AddHoldInvoiceResp
	(*SettleInvoiceMsg)(nil),              // 5: invoicesrpc.SettleInvoiceMsg
	(*SettleInvoiceResp)(nil),             // 6: invoicesrpc.SettleInvoiceResp
	(*SubscribeSingleInvoiceRequest)(nil), // 7: invoicesrpc.SubscribeSingleInvoiceRequest
	(*LookupInvoiceMsg)(nil),              // 8: invoicesrpc.LookupInvoiceMsg
	(*lnrpc.RouteHint)(nil),               // 9: lnrpc.RouteHint
	(*lnrpc.Invoice)(nil),                 // 10: lnrpc.Invoice
}
var file_invoicesrpc_invoices_proto_depIdxs = []int32{
	9,  // 0: invoicesrpc.AddHoldInvoiceRequest.route_hints:type_name -> lnrpc.RouteHint
	0,  // 1: invoicesrpc.LookupInvoiceMsg.lookup_modifier:type_name -> invoicesrpc.LookupModifier
	7,  // 2: invoicesrpc.Invoices.SubscribeSingleInvoice:input_type -> invoicesrpc.SubscribeSingleInvoiceRequest
	1,  // 3: invoicesrpc.Invoices.CancelInvoice:input_type -> invoicesrpc.CancelInvoiceMsg
	3,  // 4: invoicesrpc.Invoices.AddHoldInvoice:input_type -> invoicesrpc.AddHoldInvoiceRequest
	5,  // 5: invoicesrpc.Invoices.SettleInvoice:input_type -> invoicesrpc.SettleInvoiceMsg
	8,  // 6: invoicesrpc.Invoices.LookupInvoiceV2:input_type -> invoicesrpc.LookupInvoiceMsg
	10, // 7: invoicesrpc.Invoices.SubscribeSingleInvoice:output_type -> lnrpc.Invoice
	2,  // 8: invoicesrpc.Invoices.CancelInvoice:output_type -> invoicesrpc.CancelInvoiceResp
	4,  // 9: invoicesrpc.Invoices.AddHoldInvoice:output_type -> invoicesrpc.AddHoldInvoiceResp
	6,  // 10: invoicesrpc.Invoices.SettleInvoice:output_type -> invoicesrpc.SettleInvoiceResp
	10, // 11: invoicesrpc.Invoices.LookupInvoiceV2:output_type -> lnrpc.Invoice
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_invoicesrpc_invoices_proto_init() }
func file_invoicesrpc_invoices_proto_init() {
	if File_invoicesrpc_invoices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoicesrpc_invoices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelInvoiceMsg); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelInvoiceResp); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHoldInvoiceRequest); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHoldInvoiceResp); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleInvoiceMsg); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleInvoiceResp); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeSingleInvoiceRequest); i {
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
		file_invoicesrpc_invoices_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupInvoiceMsg); i {
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
	file_invoicesrpc_invoices_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*LookupInvoiceMsg_PaymentHash)(nil),
		(*LookupInvoiceMsg_PaymentAddr)(nil),
		(*LookupInvoiceMsg_SetId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_invoicesrpc_invoices_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoicesrpc_invoices_proto_goTypes,
		DependencyIndexes: file_invoicesrpc_invoices_proto_depIdxs,
		EnumInfos:         file_invoicesrpc_invoices_proto_enumTypes,
		MessageInfos:      file_invoicesrpc_invoices_proto_msgTypes,
	}.Build()
	File_invoicesrpc_invoices_proto = out.File
	file_invoicesrpc_invoices_proto_rawDesc = nil
	file_invoicesrpc_invoices_proto_goTypes = nil
	file_invoicesrpc_invoices_proto_depIdxs = nil
}
