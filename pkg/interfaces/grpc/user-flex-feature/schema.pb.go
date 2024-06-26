// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: docs/schema/user-flex-feature/v1/schema.proto

package user_flex_feature

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GeneralErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    string `protobuf:"bytes,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorDetails string `protobuf:"bytes,2,opt,name=error_details,json=errorDetails,proto3" json:"error_details,omitempty"`
}

func (x *GeneralErrorResponse) Reset() {
	*x = GeneralErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralErrorResponse) ProtoMessage() {}

func (x *GeneralErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralErrorResponse.ProtoReflect.Descriptor instead.
func (*GeneralErrorResponse) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralErrorResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *GeneralErrorResponse) GetErrorDetails() string {
	if x != nil {
		return x.ErrorDetails
	}
	return ""
}

type Percentage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]float64 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *Percentage) Reset() {
	*x = Percentage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Percentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Percentage) ProtoMessage() {}

func (x *Percentage) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Percentage.ProtoReflect.Descriptor instead.
func (*Percentage) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Percentage) GetValue() map[string]float64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type ProgressiveRolloutStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variation:
	//
	//	*ProgressiveRolloutStep_VariationValue
	Variation isProgressiveRolloutStep_Variation `protobuf_oneof:"variation"`
	// Types that are assignable to Percentage:
	//
	//	*ProgressiveRolloutStep_PercentageValue
	Percentage isProgressiveRolloutStep_Percentage `protobuf_oneof:"percentage"`
	// Types that are assignable to Date:
	//
	//	*ProgressiveRolloutStep_DateValue
	Date isProgressiveRolloutStep_Date `protobuf_oneof:"date"`
}

func (x *ProgressiveRolloutStep) Reset() {
	*x = ProgressiveRolloutStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressiveRolloutStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressiveRolloutStep) ProtoMessage() {}

func (x *ProgressiveRolloutStep) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressiveRolloutStep.ProtoReflect.Descriptor instead.
func (*ProgressiveRolloutStep) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (m *ProgressiveRolloutStep) GetVariation() isProgressiveRolloutStep_Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

func (x *ProgressiveRolloutStep) GetVariationValue() string {
	if x, ok := x.GetVariation().(*ProgressiveRolloutStep_VariationValue); ok {
		return x.VariationValue
	}
	return ""
}

func (m *ProgressiveRolloutStep) GetPercentage() isProgressiveRolloutStep_Percentage {
	if m != nil {
		return m.Percentage
	}
	return nil
}

func (x *ProgressiveRolloutStep) GetPercentageValue() float64 {
	if x, ok := x.GetPercentage().(*ProgressiveRolloutStep_PercentageValue); ok {
		return x.PercentageValue
	}
	return 0
}

func (m *ProgressiveRolloutStep) GetDate() isProgressiveRolloutStep_Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (x *ProgressiveRolloutStep) GetDateValue() string {
	if x, ok := x.GetDate().(*ProgressiveRolloutStep_DateValue); ok {
		return x.DateValue
	}
	return ""
}

type isProgressiveRolloutStep_Variation interface {
	isProgressiveRolloutStep_Variation()
}

type ProgressiveRolloutStep_VariationValue struct {
	VariationValue string `protobuf:"bytes,1,opt,name=variation_value,json=variationValue,proto3,oneof"`
}

func (*ProgressiveRolloutStep_VariationValue) isProgressiveRolloutStep_Variation() {}

type isProgressiveRolloutStep_Percentage interface {
	isProgressiveRolloutStep_Percentage()
}

type ProgressiveRolloutStep_PercentageValue struct {
	PercentageValue float64 `protobuf:"fixed64,2,opt,name=percentage_value,json=percentageValue,proto3,oneof"`
}

func (*ProgressiveRolloutStep_PercentageValue) isProgressiveRolloutStep_Percentage() {}

type isProgressiveRolloutStep_Date interface {
	isProgressiveRolloutStep_Date()
}

type ProgressiveRolloutStep_DateValue struct {
	DateValue string `protobuf:"bytes,3,opt,name=date_value,json=dateValue,proto3,oneof"`
}

func (*ProgressiveRolloutStep_DateValue) isProgressiveRolloutStep_Date() {}

type ProgressiveRollout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initial *ProgressiveRolloutStep `protobuf:"bytes,1,opt,name=initial,proto3" json:"initial,omitempty"`
	End     *ProgressiveRolloutStep `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ProgressiveRollout) Reset() {
	*x = ProgressiveRollout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressiveRollout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressiveRollout) ProtoMessage() {}

func (x *ProgressiveRollout) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressiveRollout.ProtoReflect.Descriptor instead.
func (*ProgressiveRollout) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ProgressiveRollout) GetInitial() *ProgressiveRolloutStep {
	if x != nil {
		return x.Initial
	}
	return nil
}

func (x *ProgressiveRollout) GetEnd() *ProgressiveRolloutStep {
	if x != nil {
		return x.End
	}
	return nil
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VariationResult string `protobuf:"bytes,2,opt,name=variation_result,json=variationResult,proto3" json:"variation_result,omitempty"`
	Query           string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Types that are assignable to Percentage:
	//
	//	*Rule_PercentageValue
	Percentage isRule_Percentage `protobuf_oneof:"percentage"`
	// Types that are assignable to ProgressiveRollout:
	//
	//	*Rule_ProgressiveRolloutValue
	ProgressiveRollout isRule_ProgressiveRollout `protobuf_oneof:"progressive_rollout"`
	// Types that are assignable to Disable:
	//
	//	*Rule_DisableValue
	Disable isRule_Disable `protobuf_oneof:"disable"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{4}
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetVariationResult() string {
	if x != nil {
		return x.VariationResult
	}
	return ""
}

func (x *Rule) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (m *Rule) GetPercentage() isRule_Percentage {
	if m != nil {
		return m.Percentage
	}
	return nil
}

func (x *Rule) GetPercentageValue() *Percentage {
	if x, ok := x.GetPercentage().(*Rule_PercentageValue); ok {
		return x.PercentageValue
	}
	return nil
}

func (m *Rule) GetProgressiveRollout() isRule_ProgressiveRollout {
	if m != nil {
		return m.ProgressiveRollout
	}
	return nil
}

func (x *Rule) GetProgressiveRolloutValue() *ProgressiveRollout {
	if x, ok := x.GetProgressiveRollout().(*Rule_ProgressiveRolloutValue); ok {
		return x.ProgressiveRolloutValue
	}
	return nil
}

func (m *Rule) GetDisable() isRule_Disable {
	if m != nil {
		return m.Disable
	}
	return nil
}

func (x *Rule) GetDisableValue() bool {
	if x, ok := x.GetDisable().(*Rule_DisableValue); ok {
		return x.DisableValue
	}
	return false
}

type isRule_Percentage interface {
	isRule_Percentage()
}

type Rule_PercentageValue struct {
	PercentageValue *Percentage `protobuf:"bytes,4,opt,name=percentage_value,json=percentageValue,proto3,oneof"`
}

func (*Rule_PercentageValue) isRule_Percentage() {}

type isRule_ProgressiveRollout interface {
	isRule_ProgressiveRollout()
}

type Rule_ProgressiveRolloutValue struct {
	ProgressiveRolloutValue *ProgressiveRollout `protobuf:"bytes,5,opt,name=progressive_rollout_value,json=progressiveRolloutValue,proto3,oneof"`
}

func (*Rule_ProgressiveRolloutValue) isRule_ProgressiveRollout() {}

type isRule_Disable interface {
	isRule_Disable()
}

type Rule_DisableValue struct {
	DisableValue bool `protobuf:"varint,6,opt,name=disable_value,json=disableValue,proto3,oneof"`
}

func (*Rule_DisableValue) isRule_Disable() {}

type RuleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Rule *Rule  `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *RuleUpdateRequest) Reset() {
	*x = RuleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleUpdateRequest) ProtoMessage() {}

func (x *RuleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RuleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{5}
}

func (x *RuleUpdateRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RuleUpdateRequest) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type RuleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RuleUpdateResponse) Reset() {
	*x = RuleUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleUpdateResponse) ProtoMessage() {}

func (x *RuleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleUpdateResponse.ProtoReflect.Descriptor instead.
func (*RuleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP(), []int{6}
}

func (x *RuleUpdateResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_docs_schema_user_flex_feature_v1_schema_proto protoreflect.FileDescriptor

var file_docs_schema_user_flex_feature_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x13, 0x92, 0x41, 0x10, 0xf2, 0x02, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x3a, 0x22, 0x92, 0x41, 0x1f, 0x0a, 0x1d, 0xd2, 0x01, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0xd2, 0x01, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x0a, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66,
	0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb4, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x65, 0x70,
	0x12, 0x29, 0x0a, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x10, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x9c, 0x01, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c,
	0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78,
	0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x8a, 0x03, 0x0a, 0x04,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4d, 0x0a, 0x10, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x66, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f,
	0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x48, 0x01, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x25, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x1f, 0x92, 0x41, 0x1c, 0x0a, 0x1a, 0xd2, 0x01,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x10, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x69, 0x0a, 0x11, 0x52, 0x75, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2e, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x3a,
	0x12, 0x92, 0x41, 0x0f, 0x0a, 0x0d, 0xd2, 0x01, 0x03, 0x6b, 0x65, 0x79, 0xd2, 0x01, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x52, 0x75, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xae, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6c, 0x65, 0x78, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a,
	0x0a, 0x52, 0x75, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78,
	0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x6b, 0x65,
	0x79, 0x7d, 0x42, 0x6b, 0x92, 0x41, 0x53, 0x12, 0x18, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x20,
	0x66, 0x6c, 0x65, 0x78, 0x20, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x52, 0x37, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x30, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5a, 0x13, 0x2e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docs_schema_user_flex_feature_v1_schema_proto_rawDescOnce sync.Once
	file_docs_schema_user_flex_feature_v1_schema_proto_rawDescData = file_docs_schema_user_flex_feature_v1_schema_proto_rawDesc
)

func file_docs_schema_user_flex_feature_v1_schema_proto_rawDescGZIP() []byte {
	file_docs_schema_user_flex_feature_v1_schema_proto_rawDescOnce.Do(func() {
		file_docs_schema_user_flex_feature_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_docs_schema_user_flex_feature_v1_schema_proto_rawDescData)
	})
	return file_docs_schema_user_flex_feature_v1_schema_proto_rawDescData
}

var file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_docs_schema_user_flex_feature_v1_schema_proto_goTypes = []interface{}{
	(*GeneralErrorResponse)(nil),   // 0: user_flex_feature.v1.GeneralErrorResponse
	(*Percentage)(nil),             // 1: user_flex_feature.v1.Percentage
	(*ProgressiveRolloutStep)(nil), // 2: user_flex_feature.v1.ProgressiveRolloutStep
	(*ProgressiveRollout)(nil),     // 3: user_flex_feature.v1.ProgressiveRollout
	(*Rule)(nil),                   // 4: user_flex_feature.v1.Rule
	(*RuleUpdateRequest)(nil),      // 5: user_flex_feature.v1.RuleUpdateRequest
	(*RuleUpdateResponse)(nil),     // 6: user_flex_feature.v1.RuleUpdateResponse
	nil,                            // 7: user_flex_feature.v1.Percentage.ValueEntry
}
var file_docs_schema_user_flex_feature_v1_schema_proto_depIdxs = []int32{
	7, // 0: user_flex_feature.v1.Percentage.value:type_name -> user_flex_feature.v1.Percentage.ValueEntry
	2, // 1: user_flex_feature.v1.ProgressiveRollout.initial:type_name -> user_flex_feature.v1.ProgressiveRolloutStep
	2, // 2: user_flex_feature.v1.ProgressiveRollout.end:type_name -> user_flex_feature.v1.ProgressiveRolloutStep
	1, // 3: user_flex_feature.v1.Rule.percentage_value:type_name -> user_flex_feature.v1.Percentage
	3, // 4: user_flex_feature.v1.Rule.progressive_rollout_value:type_name -> user_flex_feature.v1.ProgressiveRollout
	4, // 5: user_flex_feature.v1.RuleUpdateRequest.rule:type_name -> user_flex_feature.v1.Rule
	5, // 6: user_flex_feature.v1.UserFlexFeatureService.RuleUpdate:input_type -> user_flex_feature.v1.RuleUpdateRequest
	6, // 7: user_flex_feature.v1.UserFlexFeatureService.RuleUpdate:output_type -> user_flex_feature.v1.RuleUpdateResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_docs_schema_user_flex_feature_v1_schema_proto_init() }
func file_docs_schema_user_flex_feature_v1_schema_proto_init() {
	if File_docs_schema_user_flex_feature_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralErrorResponse); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Percentage); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressiveRolloutStep); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressiveRollout); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleUpdateRequest); i {
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
		file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleUpdateResponse); i {
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
	file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ProgressiveRolloutStep_VariationValue)(nil),
		(*ProgressiveRolloutStep_PercentageValue)(nil),
		(*ProgressiveRolloutStep_DateValue)(nil),
	}
	file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Rule_PercentageValue)(nil),
		(*Rule_ProgressiveRolloutValue)(nil),
		(*Rule_DisableValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_docs_schema_user_flex_feature_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_docs_schema_user_flex_feature_v1_schema_proto_goTypes,
		DependencyIndexes: file_docs_schema_user_flex_feature_v1_schema_proto_depIdxs,
		MessageInfos:      file_docs_schema_user_flex_feature_v1_schema_proto_msgTypes,
	}.Build()
	File_docs_schema_user_flex_feature_v1_schema_proto = out.File
	file_docs_schema_user_flex_feature_v1_schema_proto_rawDesc = nil
	file_docs_schema_user_flex_feature_v1_schema_proto_goTypes = nil
	file_docs_schema_user_flex_feature_v1_schema_proto_depIdxs = nil
}
