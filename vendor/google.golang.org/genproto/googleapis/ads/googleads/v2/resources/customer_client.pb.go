// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/ads/googleads/v2/resources/customer_client.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// Output only. Specifies whether this is a
	// [hidden account](https://support.google.com/google-ads/answer/7519830).
	// Read only.
	Hidden *wrappers.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Output only. Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level *wrappers.Int64Value `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	// Output only. Common Locale Data Repository (CLDR) string representation of the
	// time zone of the client, e.g. America/Los_Angeles. Read only.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Output only. Identifies if the client is a test account. Read only.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,7,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Output only. Identifies if the client is a manager. Read only.
	Manager *wrappers.BoolValue `protobuf:"bytes,8,opt,name=manager,proto3" json:"manager,omitempty"`
	// Output only. Descriptive name for the client. Read only.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,9,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// Output only. Currency code (e.g. 'USD', 'EUR') for the client. Read only.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Output only. The ID of the client customer. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerClient) Reset() {
	*x = CustomerClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_customer_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerClient) ProtoMessage() {}

func (x *CustomerClient) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_customer_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerClient.ProtoReflect.Descriptor instead.
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_customer_client_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerClient) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerClient) GetClientCustomer() *wrappers.StringValue {
	if x != nil {
		return x.ClientCustomer
	}
	return nil
}

func (x *CustomerClient) GetHidden() *wrappers.BoolValue {
	if x != nil {
		return x.Hidden
	}
	return nil
}

func (x *CustomerClient) GetLevel() *wrappers.Int64Value {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *CustomerClient) GetTimeZone() *wrappers.StringValue {
	if x != nil {
		return x.TimeZone
	}
	return nil
}

func (x *CustomerClient) GetTestAccount() *wrappers.BoolValue {
	if x != nil {
		return x.TestAccount
	}
	return nil
}

func (x *CustomerClient) GetManager() *wrappers.BoolValue {
	if x != nil {
		return x.Manager
	}
	return nil
}

func (x *CustomerClient) GetDescriptiveName() *wrappers.StringValue {
	if x != nil {
		return x.DescriptiveName
	}
	return nil
}

func (x *CustomerClient) GetCurrencyCode() *wrappers.StringValue {
	if x != nil {
		return x.CurrencyCode
	}
	return nil
}

func (x *CustomerClient) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_google_ads_googleads_v2_resources_customer_client_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_customer_client_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x06, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x4a, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x68,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3e, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x42, 0x0a, 0x0c,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x39, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x3a, 0x64, 0xea, 0x41, 0x61, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0x80, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_customer_client_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_customer_client_proto_rawDescData = file_google_ads_googleads_v2_resources_customer_client_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_customer_client_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_customer_client_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_customer_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_customer_client_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_customer_client_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_customer_client_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_resources_customer_client_proto_goTypes = []interface{}{
	(*CustomerClient)(nil),       // 0: google.ads.googleads.v2.resources.CustomerClient
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrappers.BoolValue)(nil),   // 2: google.protobuf.BoolValue
	(*wrappers.Int64Value)(nil),  // 3: google.protobuf.Int64Value
}
var file_google_ads_googleads_v2_resources_customer_client_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v2.resources.CustomerClient.client_customer:type_name -> google.protobuf.StringValue
	2, // 1: google.ads.googleads.v2.resources.CustomerClient.hidden:type_name -> google.protobuf.BoolValue
	3, // 2: google.ads.googleads.v2.resources.CustomerClient.level:type_name -> google.protobuf.Int64Value
	1, // 3: google.ads.googleads.v2.resources.CustomerClient.time_zone:type_name -> google.protobuf.StringValue
	2, // 4: google.ads.googleads.v2.resources.CustomerClient.test_account:type_name -> google.protobuf.BoolValue
	2, // 5: google.ads.googleads.v2.resources.CustomerClient.manager:type_name -> google.protobuf.BoolValue
	1, // 6: google.ads.googleads.v2.resources.CustomerClient.descriptive_name:type_name -> google.protobuf.StringValue
	1, // 7: google.ads.googleads.v2.resources.CustomerClient.currency_code:type_name -> google.protobuf.StringValue
	3, // 8: google.ads.googleads.v2.resources.CustomerClient.id:type_name -> google.protobuf.Int64Value
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_customer_client_proto_init() }
func file_google_ads_googleads_v2_resources_customer_client_proto_init() {
	if File_google_ads_googleads_v2_resources_customer_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_customer_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerClient); i {
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
			RawDescriptor: file_google_ads_googleads_v2_resources_customer_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_customer_client_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_customer_client_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_customer_client_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_customer_client_proto = out.File
	file_google_ads_googleads_v2_resources_customer_client_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_customer_client_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_customer_client_proto_depIdxs = nil
}
