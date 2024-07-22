// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/bigquery/migration/v2/translation_details.proto

package migrationpb

import (
	reflect "reflect"
	sync "sync"

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

// The translation details to capture the necessary settings for a translation
// job.
type TranslationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mapping from source to target SQL.
	SourceTargetMapping []*SourceTargetMapping `protobuf:"bytes,1,rep,name=source_target_mapping,json=sourceTargetMapping,proto3" json:"source_target_mapping,omitempty"`
	// The base URI for all writes to persistent storage.
	TargetBaseUri string `protobuf:"bytes,2,opt,name=target_base_uri,json=targetBaseUri,proto3" json:"target_base_uri,omitempty"`
	// The default source environment values for the translation.
	SourceEnvironment *SourceEnvironment `protobuf:"bytes,3,opt,name=source_environment,json=sourceEnvironment,proto3" json:"source_environment,omitempty"`
	// The list of literal targets that will be directly returned to the response.
	// Each entry consists of the constructed path, EXCLUDING the base path. Not
	// providing a target_base_uri will prevent writing to persistent storage.
	TargetReturnLiterals []string `protobuf:"bytes,4,rep,name=target_return_literals,json=targetReturnLiterals,proto3" json:"target_return_literals,omitempty"`
	// The types of output to generate, e.g. sql, metadata,
	// lineage_from_sql_scripts, etc. If not specified, a default set of
	// targets will be generated. Some additional target types may be slower to
	// generate. See the documentation for the set of available target types.
	TargetTypes []string `protobuf:"bytes,5,rep,name=target_types,json=targetTypes,proto3" json:"target_types,omitempty"`
}

func (x *TranslationDetails) Reset() {
	*x = TranslationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationDetails) ProtoMessage() {}

func (x *TranslationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationDetails.ProtoReflect.Descriptor instead.
func (*TranslationDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationDetails) GetSourceTargetMapping() []*SourceTargetMapping {
	if x != nil {
		return x.SourceTargetMapping
	}
	return nil
}

func (x *TranslationDetails) GetTargetBaseUri() string {
	if x != nil {
		return x.TargetBaseUri
	}
	return ""
}

func (x *TranslationDetails) GetSourceEnvironment() *SourceEnvironment {
	if x != nil {
		return x.SourceEnvironment
	}
	return nil
}

func (x *TranslationDetails) GetTargetReturnLiterals() []string {
	if x != nil {
		return x.TargetReturnLiterals
	}
	return nil
}

func (x *TranslationDetails) GetTargetTypes() []string {
	if x != nil {
		return x.TargetTypes
	}
	return nil
}

// Represents one mapping from a source SQL to a target SQL.
type SourceTargetMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source SQL or the path to it.
	SourceSpec *SourceSpec `protobuf:"bytes,1,opt,name=source_spec,json=sourceSpec,proto3" json:"source_spec,omitempty"`
	// The target SQL or the path for it.
	TargetSpec *TargetSpec `protobuf:"bytes,2,opt,name=target_spec,json=targetSpec,proto3" json:"target_spec,omitempty"`
}

func (x *SourceTargetMapping) Reset() {
	*x = SourceTargetMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTargetMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTargetMapping) ProtoMessage() {}

func (x *SourceTargetMapping) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTargetMapping.ProtoReflect.Descriptor instead.
func (*SourceTargetMapping) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{1}
}

func (x *SourceTargetMapping) GetSourceSpec() *SourceSpec {
	if x != nil {
		return x.SourceSpec
	}
	return nil
}

func (x *SourceTargetMapping) GetTargetSpec() *TargetSpec {
	if x != nil {
		return x.TargetSpec
	}
	return nil
}

// Represents one path to the location that holds source data.
type SourceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The specific source SQL.
	//
	// Types that are assignable to Source:
	//
	//	*SourceSpec_BaseUri
	//	*SourceSpec_Literal
	Source isSourceSpec_Source `protobuf_oneof:"source"`
	// Optional. The optional field to specify the encoding of the sql bytes.
	Encoding string `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
}

func (x *SourceSpec) Reset() {
	*x = SourceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceSpec) ProtoMessage() {}

func (x *SourceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceSpec.ProtoReflect.Descriptor instead.
func (*SourceSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{2}
}

func (m *SourceSpec) GetSource() isSourceSpec_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *SourceSpec) GetBaseUri() string {
	if x, ok := x.GetSource().(*SourceSpec_BaseUri); ok {
		return x.BaseUri
	}
	return ""
}

func (x *SourceSpec) GetLiteral() *Literal {
	if x, ok := x.GetSource().(*SourceSpec_Literal); ok {
		return x.Literal
	}
	return nil
}

func (x *SourceSpec) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

type isSourceSpec_Source interface {
	isSourceSpec_Source()
}

type SourceSpec_BaseUri struct {
	// The base URI for all files to be read in as sources for translation.
	BaseUri string `protobuf:"bytes,1,opt,name=base_uri,json=baseUri,proto3,oneof"`
}

type SourceSpec_Literal struct {
	// Source literal.
	Literal *Literal `protobuf:"bytes,2,opt,name=literal,proto3,oneof"`
}

func (*SourceSpec_BaseUri) isSourceSpec_Source() {}

func (*SourceSpec_Literal) isSourceSpec_Source() {}

// Represents one path to the location that holds target data.
type TargetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative path for the target data. Given source file
	// `base_uri/input/sql`, the output would be
	// `target_base_uri/sql/relative_path/input.sql`.
	RelativePath string `protobuf:"bytes,1,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
}

func (x *TargetSpec) Reset() {
	*x = TargetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetSpec) ProtoMessage() {}

func (x *TargetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetSpec.ProtoReflect.Descriptor instead.
func (*TargetSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{3}
}

func (x *TargetSpec) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

// Literal data.
type Literal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The literal SQL contents.
	//
	// Types that are assignable to LiteralData:
	//
	//	*Literal_LiteralString
	//	*Literal_LiteralBytes
	LiteralData isLiteral_LiteralData `protobuf_oneof:"literal_data"`
	// Required. The identifier of the literal entry.
	RelativePath string `protobuf:"bytes,1,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
}

func (x *Literal) Reset() {
	*x = Literal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Literal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Literal) ProtoMessage() {}

func (x *Literal) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Literal.ProtoReflect.Descriptor instead.
func (*Literal) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{4}
}

func (m *Literal) GetLiteralData() isLiteral_LiteralData {
	if m != nil {
		return m.LiteralData
	}
	return nil
}

func (x *Literal) GetLiteralString() string {
	if x, ok := x.GetLiteralData().(*Literal_LiteralString); ok {
		return x.LiteralString
	}
	return ""
}

func (x *Literal) GetLiteralBytes() []byte {
	if x, ok := x.GetLiteralData().(*Literal_LiteralBytes); ok {
		return x.LiteralBytes
	}
	return nil
}

func (x *Literal) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

type isLiteral_LiteralData interface {
	isLiteral_LiteralData()
}

type Literal_LiteralString struct {
	// Literal string data.
	LiteralString string `protobuf:"bytes,2,opt,name=literal_string,json=literalString,proto3,oneof"`
}

type Literal_LiteralBytes struct {
	// Literal byte data.
	LiteralBytes []byte `protobuf:"bytes,3,opt,name=literal_bytes,json=literalBytes,proto3,oneof"`
}

func (*Literal_LiteralString) isLiteral_LiteralData() {}

func (*Literal_LiteralBytes) isLiteral_LiteralData() {}

// Represents the default source environment values for the translation.
type SourceEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The default database name to fully qualify SQL objects when their database
	// name is missing.
	DefaultDatabase string `protobuf:"bytes,1,opt,name=default_database,json=defaultDatabase,proto3" json:"default_database,omitempty"`
	// The schema search path. When SQL objects are missing schema name,
	// translation engine will search through this list to find the value.
	SchemaSearchPath []string `protobuf:"bytes,2,rep,name=schema_search_path,json=schemaSearchPath,proto3" json:"schema_search_path,omitempty"`
	// Optional. Expects a validQ BigQuery dataset ID that exists, e.g.,
	// project-123.metadata_store_123.  If specified, translation will search and
	// read the required schema information from a metadata store in this dataset.
	// If metadata store doesn't exist, translation will parse the metadata file
	// and upload the schema info to a temp table in the dataset to speed up
	// future translation jobs.
	MetadataStoreDataset string `protobuf:"bytes,3,opt,name=metadata_store_dataset,json=metadataStoreDataset,proto3" json:"metadata_store_dataset,omitempty"`
}

func (x *SourceEnvironment) Reset() {
	*x = SourceEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceEnvironment) ProtoMessage() {}

func (x *SourceEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceEnvironment.ProtoReflect.Descriptor instead.
func (*SourceEnvironment) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP(), []int{5}
}

func (x *SourceEnvironment) GetDefaultDatabase() string {
	if x != nil {
		return x.DefaultDatabase
	}
	return ""
}

func (x *SourceEnvironment) GetSchemaSearchPath() []string {
	if x != nil {
		return x.SchemaSearchPath
	}
	return nil
}

func (x *SourceEnvironment) GetMetadataStoreDataset() string {
	if x != nil {
		return x.MetadataStoreDataset
	}
	return ""
}

var File_google_cloud_bigquery_migration_v2_translation_details_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6b, 0x0a, 0x15, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x69, 0x12,
	0x64, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xb7,
	0x01, 0x0a, 0x13, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x4f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x9d, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x69, 0x12, 0x47, 0x0a, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x1f, 0x0a,
	0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x93, 0x01, 0x0a, 0x07,
	0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0e, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x25, 0x0a, 0x0d, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x42, 0x0e, 0x0a, 0x0c, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x39, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x14, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0xd3, 0x01, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescData = file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_bigquery_migration_v2_translation_details_proto_goTypes = []any{
	(*TranslationDetails)(nil),  // 0: google.cloud.bigquery.migration.v2.TranslationDetails
	(*SourceTargetMapping)(nil), // 1: google.cloud.bigquery.migration.v2.SourceTargetMapping
	(*SourceSpec)(nil),          // 2: google.cloud.bigquery.migration.v2.SourceSpec
	(*TargetSpec)(nil),          // 3: google.cloud.bigquery.migration.v2.TargetSpec
	(*Literal)(nil),             // 4: google.cloud.bigquery.migration.v2.Literal
	(*SourceEnvironment)(nil),   // 5: google.cloud.bigquery.migration.v2.SourceEnvironment
}
var file_google_cloud_bigquery_migration_v2_translation_details_proto_depIdxs = []int32{
	1, // 0: google.cloud.bigquery.migration.v2.TranslationDetails.source_target_mapping:type_name -> google.cloud.bigquery.migration.v2.SourceTargetMapping
	5, // 1: google.cloud.bigquery.migration.v2.TranslationDetails.source_environment:type_name -> google.cloud.bigquery.migration.v2.SourceEnvironment
	2, // 2: google.cloud.bigquery.migration.v2.SourceTargetMapping.source_spec:type_name -> google.cloud.bigquery.migration.v2.SourceSpec
	3, // 3: google.cloud.bigquery.migration.v2.SourceTargetMapping.target_spec:type_name -> google.cloud.bigquery.migration.v2.TargetSpec
	4, // 4: google.cloud.bigquery.migration.v2.SourceSpec.literal:type_name -> google.cloud.bigquery.migration.v2.Literal
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2_translation_details_proto_init() }
func file_google_cloud_bigquery_migration_v2_translation_details_proto_init() {
	if File_google_cloud_bigquery_migration_v2_translation_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TranslationDetails); i {
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
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SourceTargetMapping); i {
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
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SourceSpec); i {
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
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TargetSpec); i {
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
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Literal); i {
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
		file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SourceEnvironment); i {
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
	file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[2].OneofWrappers = []any{
		(*SourceSpec_BaseUri)(nil),
		(*SourceSpec_Literal)(nil),
	}
	file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes[4].OneofWrappers = []any{
		(*Literal_LiteralString)(nil),
		(*Literal_LiteralBytes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2_translation_details_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2_translation_details_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_migration_v2_translation_details_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2_translation_details_proto = out.File
	file_google_cloud_bigquery_migration_v2_translation_details_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2_translation_details_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2_translation_details_proto_depIdxs = nil
}