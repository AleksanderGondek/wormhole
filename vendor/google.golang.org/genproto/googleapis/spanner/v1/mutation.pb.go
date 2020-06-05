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
// source: google/spanner/v1/mutation.proto

package spanner

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// A modification to one or more Cloud Spanner rows.  Mutations can be
// applied to a Cloud Spanner database by sending them in a
// [Commit][google.spanner.v1.Spanner.Commit] call.
type Mutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The operation to perform.
	//
	// Types that are assignable to Operation:
	//	*Mutation_Insert
	//	*Mutation_Update
	//	*Mutation_InsertOrUpdate
	//	*Mutation_Replace
	//	*Mutation_Delete_
	Operation isMutation_Operation `protobuf_oneof:"operation"`
}

func (x *Mutation) Reset() {
	*x = Mutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_mutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation) ProtoMessage() {}

func (x *Mutation) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_mutation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation.ProtoReflect.Descriptor instead.
func (*Mutation) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_mutation_proto_rawDescGZIP(), []int{0}
}

func (m *Mutation) GetOperation() isMutation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *Mutation) GetInsert() *Mutation_Write {
	if x, ok := x.GetOperation().(*Mutation_Insert); ok {
		return x.Insert
	}
	return nil
}

func (x *Mutation) GetUpdate() *Mutation_Write {
	if x, ok := x.GetOperation().(*Mutation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *Mutation) GetInsertOrUpdate() *Mutation_Write {
	if x, ok := x.GetOperation().(*Mutation_InsertOrUpdate); ok {
		return x.InsertOrUpdate
	}
	return nil
}

func (x *Mutation) GetReplace() *Mutation_Write {
	if x, ok := x.GetOperation().(*Mutation_Replace); ok {
		return x.Replace
	}
	return nil
}

func (x *Mutation) GetDelete() *Mutation_Delete {
	if x, ok := x.GetOperation().(*Mutation_Delete_); ok {
		return x.Delete
	}
	return nil
}

type isMutation_Operation interface {
	isMutation_Operation()
}

type Mutation_Insert struct {
	// Insert new rows in a table. If any of the rows already exist,
	// the write or transaction fails with error `ALREADY_EXISTS`.
	Insert *Mutation_Write `protobuf:"bytes,1,opt,name=insert,proto3,oneof"`
}

type Mutation_Update struct {
	// Update existing rows in a table. If any of the rows does not
	// already exist, the transaction fails with error `NOT_FOUND`.
	Update *Mutation_Write `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type Mutation_InsertOrUpdate struct {
	// Like [insert][google.spanner.v1.Mutation.insert], except that if the row already exists, then
	// its column values are overwritten with the ones provided. Any
	// column values not explicitly written are preserved.
	//
	// When using [insert_or_update][google.spanner.v1.Mutation.insert_or_update], just as when using [insert][google.spanner.v1.Mutation.insert], all `NOT
	// NULL` columns in the table must be given a value. This holds true
	// even when the row already exists and will therefore actually be updated.
	InsertOrUpdate *Mutation_Write `protobuf:"bytes,3,opt,name=insert_or_update,json=insertOrUpdate,proto3,oneof"`
}

type Mutation_Replace struct {
	// Like [insert][google.spanner.v1.Mutation.insert], except that if the row already exists, it is
	// deleted, and the column values provided are inserted
	// instead. Unlike [insert_or_update][google.spanner.v1.Mutation.insert_or_update], this means any values not
	// explicitly written become `NULL`.
	//
	// In an interleaved table, if you create the child table with the
	// `ON DELETE CASCADE` annotation, then replacing a parent row
	// also deletes the child rows. Otherwise, you must delete the
	// child rows before you replace the parent row.
	Replace *Mutation_Write `protobuf:"bytes,4,opt,name=replace,proto3,oneof"`
}

type Mutation_Delete_ struct {
	// Delete rows from a table. Succeeds whether or not the named
	// rows were present.
	Delete *Mutation_Delete `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

func (*Mutation_Insert) isMutation_Operation() {}

func (*Mutation_Update) isMutation_Operation() {}

func (*Mutation_InsertOrUpdate) isMutation_Operation() {}

func (*Mutation_Replace) isMutation_Operation() {}

func (*Mutation_Delete_) isMutation_Operation() {}

// Arguments to [insert][google.spanner.v1.Mutation.insert], [update][google.spanner.v1.Mutation.update], [insert_or_update][google.spanner.v1.Mutation.insert_or_update], and
// [replace][google.spanner.v1.Mutation.replace] operations.
type Mutation_Write struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The table whose rows will be written.
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// The names of the columns in [table][google.spanner.v1.Mutation.Write.table] to be written.
	//
	// The list of columns must contain enough columns to allow
	// Cloud Spanner to derive values for all primary key columns in the
	// row(s) to be modified.
	Columns []string `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	// The values to be written. `values` can contain more than one
	// list of values. If it does, then multiple rows are written, one
	// for each entry in `values`. Each list in `values` must have
	// exactly as many entries as there are entries in [columns][google.spanner.v1.Mutation.Write.columns]
	// above. Sending multiple lists is equivalent to sending multiple
	// `Mutation`s, each containing one `values` entry and repeating
	// [table][google.spanner.v1.Mutation.Write.table] and [columns][google.spanner.v1.Mutation.Write.columns]. Individual values in each list are
	// encoded as described [here][google.spanner.v1.TypeCode].
	Values []*_struct.ListValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Mutation_Write) Reset() {
	*x = Mutation_Write{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_mutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutation_Write) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation_Write) ProtoMessage() {}

func (x *Mutation_Write) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_mutation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation_Write.ProtoReflect.Descriptor instead.
func (*Mutation_Write) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_mutation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Mutation_Write) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Mutation_Write) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Mutation_Write) GetValues() []*_struct.ListValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// Arguments to [delete][google.spanner.v1.Mutation.delete] operations.
type Mutation_Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The table whose rows will be deleted.
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// Required. The primary keys of the rows within [table][google.spanner.v1.Mutation.Delete.table] to delete.  The
	// primary keys must be specified in the order in which they appear in the
	// `PRIMARY KEY()` clause of the table's equivalent DDL statement (the DDL
	// statement used to create the table).
	// Delete is idempotent. The transaction will succeed even if some or all
	// rows do not exist.
	KeySet *KeySet `protobuf:"bytes,2,opt,name=key_set,json=keySet,proto3" json:"key_set,omitempty"`
}

func (x *Mutation_Delete) Reset() {
	*x = Mutation_Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_mutation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutation_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation_Delete) ProtoMessage() {}

func (x *Mutation_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_mutation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation_Delete.ProtoReflect.Descriptor instead.
func (*Mutation_Delete) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_mutation_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Mutation_Delete) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Mutation_Delete) GetKeySet() *KeySet {
	if x != nil {
		return x.KeySet
	}
	return nil
}

var File_google_spanner_v1_mutation_proto protoreflect.FileDescriptor

var file_google_spanner_v1_mutation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x04, 0x0a, 0x08, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x06,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x5f, 0x6f, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x6b, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a,
	0x52, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6b, 0x65, 0x79,
	0x53, 0x65, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x96, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_spanner_v1_mutation_proto_rawDescOnce sync.Once
	file_google_spanner_v1_mutation_proto_rawDescData = file_google_spanner_v1_mutation_proto_rawDesc
)

func file_google_spanner_v1_mutation_proto_rawDescGZIP() []byte {
	file_google_spanner_v1_mutation_proto_rawDescOnce.Do(func() {
		file_google_spanner_v1_mutation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_v1_mutation_proto_rawDescData)
	})
	return file_google_spanner_v1_mutation_proto_rawDescData
}

var file_google_spanner_v1_mutation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_spanner_v1_mutation_proto_goTypes = []interface{}{
	(*Mutation)(nil),          // 0: google.spanner.v1.Mutation
	(*Mutation_Write)(nil),    // 1: google.spanner.v1.Mutation.Write
	(*Mutation_Delete)(nil),   // 2: google.spanner.v1.Mutation.Delete
	(*_struct.ListValue)(nil), // 3: google.protobuf.ListValue
	(*KeySet)(nil),            // 4: google.spanner.v1.KeySet
}
var file_google_spanner_v1_mutation_proto_depIdxs = []int32{
	1, // 0: google.spanner.v1.Mutation.insert:type_name -> google.spanner.v1.Mutation.Write
	1, // 1: google.spanner.v1.Mutation.update:type_name -> google.spanner.v1.Mutation.Write
	1, // 2: google.spanner.v1.Mutation.insert_or_update:type_name -> google.spanner.v1.Mutation.Write
	1, // 3: google.spanner.v1.Mutation.replace:type_name -> google.spanner.v1.Mutation.Write
	2, // 4: google.spanner.v1.Mutation.delete:type_name -> google.spanner.v1.Mutation.Delete
	3, // 5: google.spanner.v1.Mutation.Write.values:type_name -> google.protobuf.ListValue
	4, // 6: google.spanner.v1.Mutation.Delete.key_set:type_name -> google.spanner.v1.KeySet
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_spanner_v1_mutation_proto_init() }
func file_google_spanner_v1_mutation_proto_init() {
	if File_google_spanner_v1_mutation_proto != nil {
		return
	}
	file_google_spanner_v1_keys_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_spanner_v1_mutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutation); i {
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
		file_google_spanner_v1_mutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutation_Write); i {
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
		file_google_spanner_v1_mutation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutation_Delete); i {
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
	file_google_spanner_v1_mutation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Mutation_Insert)(nil),
		(*Mutation_Update)(nil),
		(*Mutation_InsertOrUpdate)(nil),
		(*Mutation_Replace)(nil),
		(*Mutation_Delete_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_spanner_v1_mutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_v1_mutation_proto_goTypes,
		DependencyIndexes: file_google_spanner_v1_mutation_proto_depIdxs,
		MessageInfos:      file_google_spanner_v1_mutation_proto_msgTypes,
	}.Build()
	File_google_spanner_v1_mutation_proto = out.File
	file_google_spanner_v1_mutation_proto_rawDesc = nil
	file_google_spanner_v1_mutation_proto_goTypes = nil
	file_google_spanner_v1_mutation_proto_depIdxs = nil
}
