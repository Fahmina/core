// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/gantry/v1/gantry.proto

package v1

import (
	v1 "go.viam.com/rdk/proto/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPositionRequest) Reset() {
	*x = GetPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionRequest) ProtoMessage() {}

func (x *GetPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionRequest.ProtoReflect.Descriptor instead.
func (*GetPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{0}
}

func (x *GetPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionsMm []float64 `protobuf:"fixed64,1,rep,packed,name=positions_mm,json=positionsMm,proto3" json:"positions_mm,omitempty"`
}

func (x *GetPositionResponse) Reset() {
	*x = GetPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionResponse) ProtoMessage() {}

func (x *GetPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionResponse.ProtoReflect.Descriptor instead.
func (*GetPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{1}
}

func (x *GetPositionResponse) GetPositionsMm() []float64 {
	if x != nil {
		return x.PositionsMm
	}
	return nil
}

type MoveToPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PositionsMm []float64      `protobuf:"fixed64,2,rep,packed,name=positions_mm,json=positionsMm,proto3" json:"positions_mm,omitempty"`
	WorldState  *v1.WorldState `protobuf:"bytes,3,opt,name=world_state,json=worldState,proto3,oneof" json:"world_state,omitempty"`
}

func (x *MoveToPositionRequest) Reset() {
	*x = MoveToPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToPositionRequest) ProtoMessage() {}

func (x *MoveToPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToPositionRequest.ProtoReflect.Descriptor instead.
func (*MoveToPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{2}
}

func (x *MoveToPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveToPositionRequest) GetPositionsMm() []float64 {
	if x != nil {
		return x.PositionsMm
	}
	return nil
}

func (x *MoveToPositionRequest) GetWorldState() *v1.WorldState {
	if x != nil {
		return x.WorldState
	}
	return nil
}

type MoveToPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveToPositionResponse) Reset() {
	*x = MoveToPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToPositionResponse) ProtoMessage() {}

func (x *MoveToPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToPositionResponse.ProtoReflect.Descriptor instead.
func (*MoveToPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{3}
}

type GetLengthsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetLengthsRequest) Reset() {
	*x = GetLengthsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLengthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLengthsRequest) ProtoMessage() {}

func (x *GetLengthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLengthsRequest.ProtoReflect.Descriptor instead.
func (*GetLengthsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{4}
}

func (x *GetLengthsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetLengthsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LengthsMm []float64 `protobuf:"fixed64,1,rep,packed,name=lengths_mm,json=lengthsMm,proto3" json:"lengths_mm,omitempty"`
}

func (x *GetLengthsResponse) Reset() {
	*x = GetLengthsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLengthsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLengthsResponse) ProtoMessage() {}

func (x *GetLengthsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLengthsResponse.ProtoReflect.Descriptor instead.
func (*GetLengthsResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{5}
}

func (x *GetLengthsResponse) GetLengthsMm() []float64 {
	if x != nil {
		return x.LengthsMm
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionsMm []float64 `protobuf:"fixed64,1,rep,packed,name=positions_mm,json=positionsMm,proto3" json:"positions_mm,omitempty"`
	LengthsMm   []float64 `protobuf:"fixed64,2,rep,packed,name=lengths_mm,json=lengthsMm,proto3" json:"lengths_mm,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_gantry_v1_gantry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetPositionsMm() []float64 {
	if x != nil {
		return x.PositionsMm
	}
	return nil
}

func (x *Status) GetLengthsMm() []float64 {
	if x != nil {
		return x.LengthsMm
	}
	return nil
}

var File_proto_api_component_gantry_v1_gantry_proto protoreflect.FileDescriptor

var file_proto_api_component_gantry_v1_gantry_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6d, 0x22,
	0xa5, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6d,
	0x12, 0x45, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x4d, 0x6f, 0x76, 0x65, 0x54,
	0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x4d, 0x6d, 0x22,
	0x4a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x4d, 0x6d, 0x32, 0x9e, 0x04, 0x0a, 0x0d,
	0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xab, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x76,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xb4, 0x01, 0x0a, 0x0e,
	0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x1a, 0x2d, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0xa7, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x73, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c,
	0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x42, 0x5b, 0x0a, 0x2a,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x2d, 0x67, 0x6f, 0x2e, 0x76,
	0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_api_component_gantry_v1_gantry_proto_rawDescOnce sync.Once
	file_proto_api_component_gantry_v1_gantry_proto_rawDescData = file_proto_api_component_gantry_v1_gantry_proto_rawDesc
)

func file_proto_api_component_gantry_v1_gantry_proto_rawDescGZIP() []byte {
	file_proto_api_component_gantry_v1_gantry_proto_rawDescOnce.Do(func() {
		file_proto_api_component_gantry_v1_gantry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_gantry_v1_gantry_proto_rawDescData)
	})
	return file_proto_api_component_gantry_v1_gantry_proto_rawDescData
}

var file_proto_api_component_gantry_v1_gantry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_api_component_gantry_v1_gantry_proto_goTypes = []interface{}{
	(*GetPositionRequest)(nil),     // 0: proto.api.component.gantry.v1.GetPositionRequest
	(*GetPositionResponse)(nil),    // 1: proto.api.component.gantry.v1.GetPositionResponse
	(*MoveToPositionRequest)(nil),  // 2: proto.api.component.gantry.v1.MoveToPositionRequest
	(*MoveToPositionResponse)(nil), // 3: proto.api.component.gantry.v1.MoveToPositionResponse
	(*GetLengthsRequest)(nil),      // 4: proto.api.component.gantry.v1.GetLengthsRequest
	(*GetLengthsResponse)(nil),     // 5: proto.api.component.gantry.v1.GetLengthsResponse
	(*Status)(nil),                 // 6: proto.api.component.gantry.v1.Status
	(*v1.WorldState)(nil),          // 7: proto.api.common.v1.WorldState
}
var file_proto_api_component_gantry_v1_gantry_proto_depIdxs = []int32{
	7, // 0: proto.api.component.gantry.v1.MoveToPositionRequest.world_state:type_name -> proto.api.common.v1.WorldState
	0, // 1: proto.api.component.gantry.v1.GantryService.GetPosition:input_type -> proto.api.component.gantry.v1.GetPositionRequest
	2, // 2: proto.api.component.gantry.v1.GantryService.MoveToPosition:input_type -> proto.api.component.gantry.v1.MoveToPositionRequest
	4, // 3: proto.api.component.gantry.v1.GantryService.GetLengths:input_type -> proto.api.component.gantry.v1.GetLengthsRequest
	1, // 4: proto.api.component.gantry.v1.GantryService.GetPosition:output_type -> proto.api.component.gantry.v1.GetPositionResponse
	3, // 5: proto.api.component.gantry.v1.GantryService.MoveToPosition:output_type -> proto.api.component.gantry.v1.MoveToPositionResponse
	5, // 6: proto.api.component.gantry.v1.GantryService.GetLengths:output_type -> proto.api.component.gantry.v1.GetLengthsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_api_component_gantry_v1_gantry_proto_init() }
func file_proto_api_component_gantry_v1_gantry_proto_init() {
	if File_proto_api_component_gantry_v1_gantry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionRequest); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionResponse); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToPositionRequest); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToPositionResponse); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLengthsRequest); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLengthsResponse); i {
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
		file_proto_api_component_gantry_v1_gantry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
	file_proto_api_component_gantry_v1_gantry_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_component_gantry_v1_gantry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_gantry_v1_gantry_proto_goTypes,
		DependencyIndexes: file_proto_api_component_gantry_v1_gantry_proto_depIdxs,
		MessageInfos:      file_proto_api_component_gantry_v1_gantry_proto_msgTypes,
	}.Build()
	File_proto_api_component_gantry_v1_gantry_proto = out.File
	file_proto_api_component_gantry_v1_gantry_proto_rawDesc = nil
	file_proto_api_component_gantry_v1_gantry_proto_goTypes = nil
	file_proto_api_component_gantry_v1_gantry_proto_depIdxs = nil
}
