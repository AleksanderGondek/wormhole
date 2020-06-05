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
// source: google/monitoring/v3/notification.proto

package monitoring

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	api "google.golang.org/genproto/googleapis/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	label "google.golang.org/genproto/googleapis/api/label"
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

// Indicates whether the channel has been verified or not. It is illegal
// to specify this field in a
// [`CreateNotificationChannel`][google.monitoring.v3.NotificationChannelService.CreateNotificationChannel]
// or an
// [`UpdateNotificationChannel`][google.monitoring.v3.NotificationChannelService.UpdateNotificationChannel]
// operation.
type NotificationChannel_VerificationStatus int32

const (
	// Sentinel value used to indicate that the state is unknown, omitted, or
	// is not applicable (as in the case of channels that neither support
	// nor require verification in order to function).
	NotificationChannel_VERIFICATION_STATUS_UNSPECIFIED NotificationChannel_VerificationStatus = 0
	// The channel has yet to be verified and requires verification to function.
	// Note that this state also applies to the case where the verification
	// process has been initiated by sending a verification code but where
	// the verification code has not been submitted to complete the process.
	NotificationChannel_UNVERIFIED NotificationChannel_VerificationStatus = 1
	// It has been proven that notifications can be received on this
	// notification channel and that someone on the project has access
	// to messages that are delivered to that channel.
	NotificationChannel_VERIFIED NotificationChannel_VerificationStatus = 2
)

// Enum value maps for NotificationChannel_VerificationStatus.
var (
	NotificationChannel_VerificationStatus_name = map[int32]string{
		0: "VERIFICATION_STATUS_UNSPECIFIED",
		1: "UNVERIFIED",
		2: "VERIFIED",
	}
	NotificationChannel_VerificationStatus_value = map[string]int32{
		"VERIFICATION_STATUS_UNSPECIFIED": 0,
		"UNVERIFIED":                      1,
		"VERIFIED":                        2,
	}
)

func (x NotificationChannel_VerificationStatus) Enum() *NotificationChannel_VerificationStatus {
	p := new(NotificationChannel_VerificationStatus)
	*p = x
	return p
}

func (x NotificationChannel_VerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationChannel_VerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_monitoring_v3_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationChannel_VerificationStatus) Type() protoreflect.EnumType {
	return &file_google_monitoring_v3_notification_proto_enumTypes[0]
}

func (x NotificationChannel_VerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationChannel_VerificationStatus.Descriptor instead.
func (NotificationChannel_VerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_monitoring_v3_notification_proto_rawDescGZIP(), []int{1, 0}
}

// A description of a notification channel. The descriptor includes
// the properties of the channel and the set of labels or fields that
// must be specified to configure channels of a given type.
type NotificationChannelDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full REST resource name for this descriptor. The format is:
	//
	//     projects/[PROJECT_ID_OR_NUMBER]/notificationChannelDescriptors/[TYPE]
	//
	// In the above, `[TYPE]` is the value of the `type` field.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// The type of notification channel, such as "email", "sms", etc.
	// Notification channel types are globally unique.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// A human-readable name for the notification channel type.  This
	// form of the name is suitable for a user interface.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A human-readable description of the notification channel
	// type. The description may include a description of the properties
	// of the channel and pointers to external documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The set of labels that must be defined to identify a particular
	// channel of the corresponding type. Each label includes a
	// description for how that field should be populated.
	Labels []*label.LabelDescriptor `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	// The tiers that support this notification channel; the project service tier
	// must be one of the supported_tiers.
	//
	// Deprecated: Do not use.
	SupportedTiers []ServiceTier `protobuf:"varint,5,rep,packed,name=supported_tiers,json=supportedTiers,proto3,enum=google.monitoring.v3.ServiceTier" json:"supported_tiers,omitempty"`
	// The product launch stage for channels of this type.
	LaunchStage api.LaunchStage `protobuf:"varint,7,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
}

func (x *NotificationChannelDescriptor) Reset() {
	*x = NotificationChannelDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationChannelDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationChannelDescriptor) ProtoMessage() {}

func (x *NotificationChannelDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationChannelDescriptor.ProtoReflect.Descriptor instead.
func (*NotificationChannelDescriptor) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationChannelDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationChannelDescriptor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotificationChannelDescriptor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *NotificationChannelDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationChannelDescriptor) GetLabels() []*label.LabelDescriptor {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Deprecated: Do not use.
func (x *NotificationChannelDescriptor) GetSupportedTiers() []ServiceTier {
	if x != nil {
		return x.SupportedTiers
	}
	return nil
}

func (x *NotificationChannelDescriptor) GetLaunchStage() api.LaunchStage {
	if x != nil {
		return x.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

// A `NotificationChannel` is a medium through which an alert is
// delivered when a policy violation is detected. Examples of channels
// include email, SMS, and third-party messaging applications. Fields
// containing sensitive information like authentication tokens or
// contact info are only partially populated on retrieval.
type NotificationChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the notification channel. This field matches the
	// value of the [NotificationChannelDescriptor.type][google.monitoring.v3.NotificationChannelDescriptor.type] field.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The full REST resource name for this channel. The format is:
	//
	//     projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	//
	// The `[CHANNEL_ID]` is automatically assigned by the server on creation.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// An optional human-readable name for this notification channel. It is
	// recommended that you specify a non-empty and unique name in order to
	// make it easier to identify the channels in your project, though this is
	// not enforced. The display name is limited to 512 Unicode characters.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// An optional human-readable description of this notification channel. This
	// description may provide additional details, beyond the display
	// name, for the channel. This may not exceed 1024 Unicode characters.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// [NotificationChannelDescriptor.labels][google.monitoring.v3.NotificationChannelDescriptor.labels] of the
	// `NotificationChannelDescriptor` corresponding to the `type` field.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User-supplied key/value data that does not need to conform to
	// the corresponding `NotificationChannelDescriptor`'s schema, unlike
	// the `labels` field. This field is intended to be used for organizing
	// and identifying the `NotificationChannel` objects.
	//
	// The field can contain up to 64 entries. Each key and value is limited to
	// 63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	// values can contain only lowercase letters, numerals, underscores, and
	// dashes. Keys must begin with a letter.
	UserLabels map[string]string `protobuf:"bytes,8,rep,name=user_labels,json=userLabels,proto3" json:"user_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Indicates whether this channel has been verified or not. On a
	// [`ListNotificationChannels`][google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	// or
	// [`GetNotificationChannel`][google.monitoring.v3.NotificationChannelService.GetNotificationChannel]
	// operation, this field is expected to be populated.
	//
	// If the value is `UNVERIFIED`, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification);
	// otherwise, it is assumed that the channel works.
	//
	// If the channel is neither `VERIFIED` nor `UNVERIFIED`, it implies that
	// the channel is of a type that does not require verification or that
	// this specific channel has been exempted from verification because it was
	// created prior to verification being required for channels of this type.
	//
	// This field cannot be modified using a standard
	// [`UpdateNotificationChannel`][google.monitoring.v3.NotificationChannelService.UpdateNotificationChannel]
	// operation. To change the value of this field, you must call
	// [`VerifyNotificationChannel`][google.monitoring.v3.NotificationChannelService.VerifyNotificationChannel].
	VerificationStatus NotificationChannel_VerificationStatus `protobuf:"varint,9,opt,name=verification_status,json=verificationStatus,proto3,enum=google.monitoring.v3.NotificationChannel_VerificationStatus" json:"verification_status,omitempty"`
	// Whether notifications are forwarded to the described channel. This makes
	// it possible to disable delivery of notifications to a particular channel
	// without removing the channel from all alerting policies that reference
	// the channel. This is a more convenient approach when the change is
	// temporary and you want to receive notifications from the same set
	// of alerting policies on the channel at some point in the future.
	Enabled *wrappers.BoolValue `protobuf:"bytes,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *NotificationChannel) Reset() {
	*x = NotificationChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationChannel) ProtoMessage() {}

func (x *NotificationChannel) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationChannel.ProtoReflect.Descriptor instead.
func (*NotificationChannel) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationChannel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotificationChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationChannel) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *NotificationChannel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationChannel) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *NotificationChannel) GetUserLabels() map[string]string {
	if x != nil {
		return x.UserLabels
	}
	return nil
}

func (x *NotificationChannel) GetVerificationStatus() NotificationChannel_VerificationStatus {
	if x != nil {
		return x.VerificationStatus
	}
	return NotificationChannel_VERIFICATION_STATUS_UNSPECIFIED
}

func (x *NotificationChannel) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

var File_google_monitoring_v3_notification_proto protoreflect.FileDescriptor

var file_google_monitoring_v3_notification_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a,
	0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x04, 0x0a, 0x1d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x69, 0x65, 0x72, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x3a, 0xa0, 0x02, 0xea, 0x41, 0x9c, 0x02, 0x0a, 0x37, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x46, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x7d, 0x12, 0x50, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x7d, 0x12, 0x44,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x7d, 0x12, 0x01, 0x2a, 0x22, 0xa6, 0x07, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5a, 0x0a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x6d, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x3a,
	0xfe, 0x01, 0xea, 0x41, 0xfa, 0x01, 0x0a, 0x2d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x7b,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x12, 0x48, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x12,
	0x3c, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x12, 0x01, 0x2a,
	0x42, 0xc9, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x33, 0xca,
	0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x33, 0xea, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_v3_notification_proto_rawDescOnce sync.Once
	file_google_monitoring_v3_notification_proto_rawDescData = file_google_monitoring_v3_notification_proto_rawDesc
)

func file_google_monitoring_v3_notification_proto_rawDescGZIP() []byte {
	file_google_monitoring_v3_notification_proto_rawDescOnce.Do(func() {
		file_google_monitoring_v3_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_v3_notification_proto_rawDescData)
	})
	return file_google_monitoring_v3_notification_proto_rawDescData
}

var file_google_monitoring_v3_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_monitoring_v3_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_monitoring_v3_notification_proto_goTypes = []interface{}{
	(NotificationChannel_VerificationStatus)(0), // 0: google.monitoring.v3.NotificationChannel.VerificationStatus
	(*NotificationChannelDescriptor)(nil),       // 1: google.monitoring.v3.NotificationChannelDescriptor
	(*NotificationChannel)(nil),                 // 2: google.monitoring.v3.NotificationChannel
	nil,                                         // 3: google.monitoring.v3.NotificationChannel.LabelsEntry
	nil,                                         // 4: google.monitoring.v3.NotificationChannel.UserLabelsEntry
	(*label.LabelDescriptor)(nil),               // 5: google.api.LabelDescriptor
	(ServiceTier)(0),                            // 6: google.monitoring.v3.ServiceTier
	(api.LaunchStage)(0),                        // 7: google.api.LaunchStage
	(*wrappers.BoolValue)(nil),                  // 8: google.protobuf.BoolValue
}
var file_google_monitoring_v3_notification_proto_depIdxs = []int32{
	5, // 0: google.monitoring.v3.NotificationChannelDescriptor.labels:type_name -> google.api.LabelDescriptor
	6, // 1: google.monitoring.v3.NotificationChannelDescriptor.supported_tiers:type_name -> google.monitoring.v3.ServiceTier
	7, // 2: google.monitoring.v3.NotificationChannelDescriptor.launch_stage:type_name -> google.api.LaunchStage
	3, // 3: google.monitoring.v3.NotificationChannel.labels:type_name -> google.monitoring.v3.NotificationChannel.LabelsEntry
	4, // 4: google.monitoring.v3.NotificationChannel.user_labels:type_name -> google.monitoring.v3.NotificationChannel.UserLabelsEntry
	0, // 5: google.monitoring.v3.NotificationChannel.verification_status:type_name -> google.monitoring.v3.NotificationChannel.VerificationStatus
	8, // 6: google.monitoring.v3.NotificationChannel.enabled:type_name -> google.protobuf.BoolValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_monitoring_v3_notification_proto_init() }
func file_google_monitoring_v3_notification_proto_init() {
	if File_google_monitoring_v3_notification_proto != nil {
		return
	}
	file_google_monitoring_v3_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_v3_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationChannelDescriptor); i {
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
		file_google_monitoring_v3_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationChannel); i {
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
			RawDescriptor: file_google_monitoring_v3_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_v3_notification_proto_goTypes,
		DependencyIndexes: file_google_monitoring_v3_notification_proto_depIdxs,
		EnumInfos:         file_google_monitoring_v3_notification_proto_enumTypes,
		MessageInfos:      file_google_monitoring_v3_notification_proto_msgTypes,
	}.Build()
	File_google_monitoring_v3_notification_proto = out.File
	file_google_monitoring_v3_notification_proto_rawDesc = nil
	file_google_monitoring_v3_notification_proto_goTypes = nil
	file_google_monitoring_v3_notification_proto_depIdxs = nil
}
