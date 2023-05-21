// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/p2p.proto

package api

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

type HealthStatus int32

const (
	HealthStatus_HEALTH_STATUS_UNSPECIFIED HealthStatus = 0
	HealthStatus_HEALTH_STATUS_OK          HealthStatus = 1
	HealthStatus_HEALTH_STATUS_ERROR       HealthStatus = 2
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_UNSPECIFIED",
		1: "HEALTH_STATUS_OK",
		2: "HEALTH_STATUS_ERROR",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_UNSPECIFIED": 0,
		"HEALTH_STATUS_OK":          1,
		"HEALTH_STATUS_ERROR":       2,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_p2p_proto_enumTypes[0].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_api_p2p_proto_enumTypes[0]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{2}
}

func (x *DisconnectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{4}
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Event    string `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{5}
}

func (x *EventRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *EventRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{6}
}

func (x *EventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Body     string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{7}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Event) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthStatus `protobuf:"varint,1,opt,name=status,proto3,enum=p2p.api.HealthStatus" json:"status,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_p2p_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_p2p_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_api_p2p_proto_rawDescGZIP(), []int{8}
}

func (x *HealthResponse) GetStatus() HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthStatus_HEALTH_STATUS_UNSPECIFIED
}

var File_api_p2p_proto protoreflect.FileDescriptor

var file_api_p2p_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x2e, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x41, 0x0a, 0x0c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x3f, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x5c, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x32, 0xbe, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x3e, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a,
	0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x32, 0x70,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x32, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x32, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x0e, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x67, 0x79, 0x75, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_p2p_proto_rawDescOnce sync.Once
	file_api_p2p_proto_rawDescData = file_api_p2p_proto_rawDesc
)

func file_api_p2p_proto_rawDescGZIP() []byte {
	file_api_p2p_proto_rawDescOnce.Do(func() {
		file_api_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_p2p_proto_rawDescData)
	})
	return file_api_p2p_proto_rawDescData
}

var file_api_p2p_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_p2p_proto_goTypes = []interface{}{
	(HealthStatus)(0),          // 0: p2p.api.HealthStatus
	(*ConnectRequest)(nil),     // 1: p2p.api.ConnectRequest
	(*ConnectResponse)(nil),    // 2: p2p.api.ConnectResponse
	(*DisconnectRequest)(nil),  // 3: p2p.api.DisconnectRequest
	(*DisconnectResponse)(nil), // 4: p2p.api.DisconnectResponse
	(*Empty)(nil),              // 5: p2p.api.Empty
	(*EventRequest)(nil),       // 6: p2p.api.EventRequest
	(*EventResponse)(nil),      // 7: p2p.api.EventResponse
	(*Event)(nil),              // 8: p2p.api.Event
	(*HealthResponse)(nil),     // 9: p2p.api.HealthResponse
}
var file_api_p2p_proto_depIdxs = []int32{
	0, // 0: p2p.api.HealthResponse.status:type_name -> p2p.api.HealthStatus
	1, // 1: p2p.api.PeerToPeer.Connect:input_type -> p2p.api.ConnectRequest
	3, // 2: p2p.api.PeerToPeer.Disconnect:input_type -> p2p.api.DisconnectRequest
	6, // 3: p2p.api.PeerToPeer.ProduceEvent:input_type -> p2p.api.EventRequest
	6, // 4: p2p.api.PeerToPeer.ProduceEventStream:input_type -> p2p.api.EventRequest
	8, // 5: p2p.api.PeerToPeer.RelayEvent:input_type -> p2p.api.Event
	8, // 6: p2p.api.PeerToPeer.RelayEventStream:input_type -> p2p.api.Event
	5, // 7: p2p.api.PeerToPeer.Health:input_type -> p2p.api.Empty
	2, // 8: p2p.api.PeerToPeer.Connect:output_type -> p2p.api.ConnectResponse
	4, // 9: p2p.api.PeerToPeer.Disconnect:output_type -> p2p.api.DisconnectResponse
	7, // 10: p2p.api.PeerToPeer.ProduceEvent:output_type -> p2p.api.EventResponse
	7, // 11: p2p.api.PeerToPeer.ProduceEventStream:output_type -> p2p.api.EventResponse
	5, // 12: p2p.api.PeerToPeer.RelayEvent:output_type -> p2p.api.Empty
	5, // 13: p2p.api.PeerToPeer.RelayEventStream:output_type -> p2p.api.Empty
	9, // 14: p2p.api.PeerToPeer.Health:output_type -> p2p.api.HealthResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_p2p_proto_init() }
func file_api_p2p_proto_init() {
	if File_api_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_api_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_api_p2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_api_p2p_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
		file_api_p2p_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_p2p_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_api_p2p_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_api_p2p_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_api_p2p_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_api_p2p_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_p2p_proto_goTypes,
		DependencyIndexes: file_api_p2p_proto_depIdxs,
		EnumInfos:         file_api_p2p_proto_enumTypes,
		MessageInfos:      file_api_p2p_proto_msgTypes,
	}.Build()
	File_api_p2p_proto = out.File
	file_api_p2p_proto_rawDesc = nil
	file_api_p2p_proto_goTypes = nil
	file_api_p2p_proto_depIdxs = nil
}
