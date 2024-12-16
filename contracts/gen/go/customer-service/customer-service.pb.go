// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: customer-service/customer-service.proto

package customer_contract

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

type CustomerProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid                string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserId              string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PreferredCategories []string `protobuf:"bytes,3,rep,name=preferred_categories,json=preferredCategories,proto3" json:"preferred_categories,omitempty"`
}

func (x *CustomerProfile) Reset() {
	*x = CustomerProfile{}
	mi := &file_customer_service_customer_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerProfile) ProtoMessage() {}

func (x *CustomerProfile) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_customer_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerProfile.ProtoReflect.Descriptor instead.
func (*CustomerProfile) Descriptor() ([]byte, []int) {
	return file_customer_service_customer_service_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerProfile) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CustomerProfile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CustomerProfile) GetPreferredCategories() []string {
	if x != nil {
		return x.PreferredCategories
	}
	return nil
}

// GET
type GetCustomerProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerUuid string `protobuf:"bytes,1,opt,name=customer_uuid,json=customerUuid,proto3" json:"customer_uuid,omitempty"`
}

func (x *GetCustomerProfileRequest) Reset() {
	*x = GetCustomerProfileRequest{}
	mi := &file_customer_service_customer_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerProfileRequest) ProtoMessage() {}

func (x *GetCustomerProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_customer_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerProfileRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerProfileRequest) Descriptor() ([]byte, []int) {
	return file_customer_service_customer_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerProfileRequest) GetCustomerUuid() string {
	if x != nil {
		return x.CustomerUuid
	}
	return ""
}

type GetCustomerProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *CustomerProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *GetCustomerProfileResponse) Reset() {
	*x = GetCustomerProfileResponse{}
	mi := &file_customer_service_customer_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerProfileResponse) ProtoMessage() {}

func (x *GetCustomerProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_customer_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerProfileResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerProfileResponse) Descriptor() ([]byte, []int) {
	return file_customer_service_customer_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCustomerProfileResponse) GetProfile() *CustomerProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// UPDATE
type UpdateCustomerProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerUuid    string           `protobuf:"bytes,1,opt,name=customer_uuid,json=customerUuid,proto3" json:"customer_uuid,omitempty"`
	CustomerProfile *CustomerProfile `protobuf:"bytes,2,opt,name=customer_profile,json=customerProfile,proto3" json:"customer_profile,omitempty"`
}

func (x *UpdateCustomerProfileRequest) Reset() {
	*x = UpdateCustomerProfileRequest{}
	mi := &file_customer_service_customer_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerProfileRequest) ProtoMessage() {}

func (x *UpdateCustomerProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_customer_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerProfileRequest) Descriptor() ([]byte, []int) {
	return file_customer_service_customer_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCustomerProfileRequest) GetCustomerUuid() string {
	if x != nil {
		return x.CustomerUuid
	}
	return ""
}

func (x *UpdateCustomerProfileRequest) GetCustomerProfile() *CustomerProfile {
	if x != nil {
		return x.CustomerProfile
	}
	return nil
}

type UpdateCustomerProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *CustomerProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateCustomerProfileResponse) Reset() {
	*x = UpdateCustomerProfileResponse{}
	mi := &file_customer_service_customer_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerProfileResponse) ProtoMessage() {}

func (x *UpdateCustomerProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_customer_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateCustomerProfileResponse) Descriptor() ([]byte, []int) {
	return file_customer_service_customer_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCustomerProfileResponse) GetProfile() *CustomerProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_customer_service_customer_service_proto protoreflect.FileDescriptor

var file_customer_service_customer_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x40,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x22, 0x59, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x4c, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x5c, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xfc, 0x01,
	0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x78, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37,
	0x77, 0x69, 0x6c, 0x64, 0x62, 0x65, 0x72, 0x72, 0x69, 0x65, 0x73, 0x32, 0x2e, 0x30, 0x2d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_service_customer_service_proto_rawDescOnce sync.Once
	file_customer_service_customer_service_proto_rawDescData = file_customer_service_customer_service_proto_rawDesc
)

func file_customer_service_customer_service_proto_rawDescGZIP() []byte {
	file_customer_service_customer_service_proto_rawDescOnce.Do(func() {
		file_customer_service_customer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_service_customer_service_proto_rawDescData)
	})
	return file_customer_service_customer_service_proto_rawDescData
}

var file_customer_service_customer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_customer_service_customer_service_proto_goTypes = []any{
	(*CustomerProfile)(nil),               // 0: customer_service.CustomerProfile
	(*GetCustomerProfileRequest)(nil),     // 1: customer_service.GetCustomerProfileRequest
	(*GetCustomerProfileResponse)(nil),    // 2: customer_service.GetCustomerProfileResponse
	(*UpdateCustomerProfileRequest)(nil),  // 3: customer_service.UpdateCustomerProfileRequest
	(*UpdateCustomerProfileResponse)(nil), // 4: customer_service.UpdateCustomerProfileResponse
}
var file_customer_service_customer_service_proto_depIdxs = []int32{
	0, // 0: customer_service.GetCustomerProfileResponse.profile:type_name -> customer_service.CustomerProfile
	0, // 1: customer_service.UpdateCustomerProfileRequest.customer_profile:type_name -> customer_service.CustomerProfile
	0, // 2: customer_service.UpdateCustomerProfileResponse.profile:type_name -> customer_service.CustomerProfile
	1, // 3: customer_service.CustomerService.GetCustomerProfile:input_type -> customer_service.GetCustomerProfileRequest
	3, // 4: customer_service.CustomerService.UpdateCustomerProfile:input_type -> customer_service.UpdateCustomerProfileRequest
	2, // 5: customer_service.CustomerService.GetCustomerProfile:output_type -> customer_service.GetCustomerProfileResponse
	4, // 6: customer_service.CustomerService.UpdateCustomerProfile:output_type -> customer_service.UpdateCustomerProfileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customer_service_customer_service_proto_init() }
func file_customer_service_customer_service_proto_init() {
	if File_customer_service_customer_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customer_service_customer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_service_customer_service_proto_goTypes,
		DependencyIndexes: file_customer_service_customer_service_proto_depIdxs,
		MessageInfos:      file_customer_service_customer_service_proto_msgTypes,
	}.Build()
	File_customer_service_customer_service_proto = out.File
	file_customer_service_customer_service_proto_rawDesc = nil
	file_customer_service_customer_service_proto_goTypes = nil
	file_customer_service_customer_service_proto_depIdxs = nil
}