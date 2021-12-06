// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: bert.proto

package grpcapi

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

type EncodeRequest_PoolingStrategy int32

const (
	// get the encoding state corresponding to [CLS], i.e. the first token (default)
	EncodeRequest_CLS_TOKEN EncodeRequest_PoolingStrategy = 0
	// take the average of the encoding states
	EncodeRequest_REDUCE_MEAN EncodeRequest_PoolingStrategy = 1
	//  take the maximum of the encoding states
	EncodeRequest_REDUCE_MAX EncodeRequest_PoolingStrategy = 2
	// do REDUCE_MEAN and REDUCE_MAX separately and then concat them together
	EncodeRequest_REDUCE_MEAN_MAX EncodeRequest_PoolingStrategy = 3
)

// Enum value maps for EncodeRequest_PoolingStrategy.
var (
	EncodeRequest_PoolingStrategy_name = map[int32]string{
		0: "CLS_TOKEN",
		1: "REDUCE_MEAN",
		2: "REDUCE_MAX",
		3: "REDUCE_MEAN_MAX",
	}
	EncodeRequest_PoolingStrategy_value = map[string]int32{
		"CLS_TOKEN":       0,
		"REDUCE_MEAN":     1,
		"REDUCE_MAX":      2,
		"REDUCE_MEAN_MAX": 3,
	}
)

func (x EncodeRequest_PoolingStrategy) Enum() *EncodeRequest_PoolingStrategy {
	p := new(EncodeRequest_PoolingStrategy)
	*p = x
	return p
}

func (x EncodeRequest_PoolingStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodeRequest_PoolingStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_bert_proto_enumTypes[0].Descriptor()
}

func (EncodeRequest_PoolingStrategy) Type() protoreflect.EnumType {
	return &file_bert_proto_enumTypes[0]
}

func (x EncodeRequest_PoolingStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodeRequest_PoolingStrategy.Descriptor instead.
func (EncodeRequest_PoolingStrategy) EnumDescriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{8, 0}
}

// The answer request message containing the passage and question to answer.
type AnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passage  string `protobuf:"bytes,1,opt,name=passage,proto3" json:"passage,omitempty"`
	Question string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *AnswerRequest) Reset() {
	*x = AnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerRequest) ProtoMessage() {}

func (x *AnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerRequest.ProtoReflect.Descriptor instead.
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{0}
}

func (x *AnswerRequest) GetPassage() string {
	if x != nil {
		return x.Passage
	}
	return ""
}

func (x *AnswerRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

// The response message containing the answers.
type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string  `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Start      int32   `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End        int32   `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Confidence float64 `protobuf:"fixed64,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{1}
}

func (x *Answer) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Answer) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Answer) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Answer) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type AnswerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*Answer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
	// Took is the number of milliseconds it took the server to execute the request.
	Took int64 `protobuf:"varint,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *AnswerReply) Reset() {
	*x = AnswerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerReply) ProtoMessage() {}

func (x *AnswerReply) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerReply.ProtoReflect.Descriptor instead.
func (*AnswerReply) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{2}
}

func (x *AnswerReply) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *AnswerReply) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// The discriminate request message containing the text.
type DiscriminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *DiscriminateRequest) Reset() {
	*x = DiscriminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscriminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscriminateRequest) ProtoMessage() {}

func (x *DiscriminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscriminateRequest.ProtoReflect.Descriptor instead.
func (*DiscriminateRequest) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{3}
}

func (x *DiscriminateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// The response message containing the tokens from discriminate analysis.
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Start int32  `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End   int32  `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{4}
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Token) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Token) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type DiscriminateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	// Took is the number of milliseconds it took the server to execute the request.
	Took int64 `protobuf:"varint,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *DiscriminateReply) Reset() {
	*x = DiscriminateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscriminateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscriminateReply) ProtoMessage() {}

func (x *DiscriminateReply) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscriminateReply.ProtoReflect.Descriptor instead.
func (*DiscriminateReply) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{5}
}

func (x *DiscriminateReply) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *DiscriminateReply) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// The predict request message containing the text.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{6}
}

func (x *PredictRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// The response message containing the tokens from BERT prediction.
type PredictReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	// Took is the number of milliseconds it took the server to execute the request.
	Took int64 `protobuf:"varint,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *PredictReply) Reset() {
	*x = PredictReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictReply) ProtoMessage() {}

func (x *PredictReply) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictReply.ProtoReflect.Descriptor instead.
func (*PredictReply) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{7}
}

func (x *PredictReply) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *PredictReply) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// The encode request message containing the text.
type EncodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text            string                        `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	PoolingStrategy EncodeRequest_PoolingStrategy `protobuf:"varint,2,opt,name=pooling_strategy,json=poolingStrategy,proto3,enum=bert.grpcapi.EncodeRequest_PoolingStrategy" json:"pooling_strategy,omitempty"`
}

func (x *EncodeRequest) Reset() {
	*x = EncodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeRequest) ProtoMessage() {}

func (x *EncodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeRequest.ProtoReflect.Descriptor instead.
func (*EncodeRequest) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{8}
}

func (x *EncodeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EncodeRequest) GetPoolingStrategy() EncodeRequest_PoolingStrategy {
	if x != nil {
		return x.PoolingStrategy
	}
	return EncodeRequest_CLS_TOKEN
}

// The response message containing the tokens from BERT prediction.
type EncodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector []float32 `protobuf:"fixed32,1,rep,packed,name=vector,proto3" json:"vector,omitempty"`
	// Took is the number of milliseconds it took the server to execute the request.
	Took int64 `protobuf:"varint,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *EncodeReply) Reset() {
	*x = EncodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeReply) ProtoMessage() {}

func (x *EncodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeReply.ProtoReflect.Descriptor instead.
func (*EncodeReply) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{9}
}

func (x *EncodeReply) GetVector() []float32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *EncodeReply) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// The classify request message containing the text to classify
type ClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasText2 bool   `protobuf:"varint,1,opt,name=has_text2,json=hasText2,proto3" json:"has_text2,omitempty"` // always set this to "true" when using text2
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Text2    string `protobuf:"bytes,3,opt,name=text2,proto3" json:"text2,omitempty"`
}

func (x *ClassifyRequest) Reset() {
	*x = ClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyRequest) ProtoMessage() {}

func (x *ClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyRequest.ProtoReflect.Descriptor instead.
func (*ClassifyRequest) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{10}
}

func (x *ClassifyRequest) GetHasText2() bool {
	if x != nil {
		return x.HasText2
	}
	return false
}

func (x *ClassifyRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ClassifyRequest) GetText2() string {
	if x != nil {
		return x.Text2
	}
	return ""
}

// The response message containing the classification.
type ClassConfidencePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class      string  `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *ClassConfidencePair) Reset() {
	*x = ClassConfidencePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassConfidencePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassConfidencePair) ProtoMessage() {}

func (x *ClassConfidencePair) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassConfidencePair.ProtoReflect.Descriptor instead.
func (*ClassConfidencePair) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{11}
}

func (x *ClassConfidencePair) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *ClassConfidencePair) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type ClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class        string                 `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Confidence   float64                `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Distribution []*ClassConfidencePair `protobuf:"bytes,3,rep,name=distribution,proto3" json:"distribution,omitempty"`
	// Took is the number of milliseconds it took the server to execute the request.
	Took int64 `protobuf:"varint,4,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *ClassifyReply) Reset() {
	*x = ClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bert_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyReply) ProtoMessage() {}

func (x *ClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_bert_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyReply.ProtoReflect.Descriptor instead.
func (*ClassifyReply) Descriptor() ([]byte, []int) {
	return file_bert_proto_rawDescGZIP(), []int{12}
}

func (x *ClassifyReply) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *ClassifyReply) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *ClassifyReply) GetDistribution() []*ClassConfidencePair {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *ClassifyReply) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

var File_bert_proto protoreflect.FileDescriptor

var file_bert_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65,
	0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x45, 0x0a, 0x0d, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x64, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x69,
	0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x59, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0x54, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4f, 0x0a, 0x0c,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x22, 0xd3, 0x01,
	0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x56, 0x0a, 0x10, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0f, 0x70, 0x6f, 0x6f, 0x6c,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x56, 0x0a, 0x0f, 0x50,
	0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4c, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x02, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x41, 0x4e, 0x5f, 0x4d, 0x41,
	0x58, 0x10, 0x03, 0x22, 0x39, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f,
	0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x22, 0x58,
	0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x54, 0x65, 0x78, 0x74, 0x32, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x78, 0x74, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x65, 0x78, 0x74, 0x32, 0x22, 0x4b, 0x0a, 0x13, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x32, 0xf5, 0x02, 0x0a, 0x04, 0x42, 0x45, 0x52,
	0x54, 0x12, 0x42, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x62, 0x65,
	0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x07, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x62,
	0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x65, 0x72, 0x74,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x79, 0x12, 0x1d, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x6c, 0x70, 0x6f, 0x64, 0x79, 0x73, 0x73, 0x65, 0x79, 0x2f, 0x73, 0x70, 0x61, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6c, 0x70, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x65, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bert_proto_rawDescOnce sync.Once
	file_bert_proto_rawDescData = file_bert_proto_rawDesc
)

func file_bert_proto_rawDescGZIP() []byte {
	file_bert_proto_rawDescOnce.Do(func() {
		file_bert_proto_rawDescData = protoimpl.X.CompressGZIP(file_bert_proto_rawDescData)
	})
	return file_bert_proto_rawDescData
}

var file_bert_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bert_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_bert_proto_goTypes = []interface{}{
	(EncodeRequest_PoolingStrategy)(0), // 0: bert.grpcapi.EncodeRequest.PoolingStrategy
	(*AnswerRequest)(nil),              // 1: bert.grpcapi.AnswerRequest
	(*Answer)(nil),                     // 2: bert.grpcapi.Answer
	(*AnswerReply)(nil),                // 3: bert.grpcapi.AnswerReply
	(*DiscriminateRequest)(nil),        // 4: bert.grpcapi.DiscriminateRequest
	(*Token)(nil),                      // 5: bert.grpcapi.Token
	(*DiscriminateReply)(nil),          // 6: bert.grpcapi.DiscriminateReply
	(*PredictRequest)(nil),             // 7: bert.grpcapi.PredictRequest
	(*PredictReply)(nil),               // 8: bert.grpcapi.PredictReply
	(*EncodeRequest)(nil),              // 9: bert.grpcapi.EncodeRequest
	(*EncodeReply)(nil),                // 10: bert.grpcapi.EncodeReply
	(*ClassifyRequest)(nil),            // 11: bert.grpcapi.ClassifyRequest
	(*ClassConfidencePair)(nil),        // 12: bert.grpcapi.ClassConfidencePair
	(*ClassifyReply)(nil),              // 13: bert.grpcapi.ClassifyReply
}
var file_bert_proto_depIdxs = []int32{
	2,  // 0: bert.grpcapi.AnswerReply.answers:type_name -> bert.grpcapi.Answer
	5,  // 1: bert.grpcapi.DiscriminateReply.tokens:type_name -> bert.grpcapi.Token
	5,  // 2: bert.grpcapi.PredictReply.tokens:type_name -> bert.grpcapi.Token
	0,  // 3: bert.grpcapi.EncodeRequest.pooling_strategy:type_name -> bert.grpcapi.EncodeRequest.PoolingStrategy
	12, // 4: bert.grpcapi.ClassifyReply.distribution:type_name -> bert.grpcapi.ClassConfidencePair
	1,  // 5: bert.grpcapi.BERT.Answer:input_type -> bert.grpcapi.AnswerRequest
	4,  // 6: bert.grpcapi.BERT.Discriminate:input_type -> bert.grpcapi.DiscriminateRequest
	7,  // 7: bert.grpcapi.BERT.Predict:input_type -> bert.grpcapi.PredictRequest
	9,  // 8: bert.grpcapi.BERT.Encode:input_type -> bert.grpcapi.EncodeRequest
	11, // 9: bert.grpcapi.BERT.Classify:input_type -> bert.grpcapi.ClassifyRequest
	3,  // 10: bert.grpcapi.BERT.Answer:output_type -> bert.grpcapi.AnswerReply
	6,  // 11: bert.grpcapi.BERT.Discriminate:output_type -> bert.grpcapi.DiscriminateReply
	8,  // 12: bert.grpcapi.BERT.Predict:output_type -> bert.grpcapi.PredictReply
	10, // 13: bert.grpcapi.BERT.Encode:output_type -> bert.grpcapi.EncodeReply
	13, // 14: bert.grpcapi.BERT.Classify:output_type -> bert.grpcapi.ClassifyReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_bert_proto_init() }
func file_bert_proto_init() {
	if File_bert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerRequest); i {
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
		file_bert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_bert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerReply); i {
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
		file_bert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscriminateRequest); i {
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
		file_bert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_bert_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscriminateReply); i {
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
		file_bert_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_bert_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictReply); i {
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
		file_bert_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeRequest); i {
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
		file_bert_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeReply); i {
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
		file_bert_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyRequest); i {
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
		file_bert_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassConfidencePair); i {
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
		file_bert_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyReply); i {
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
			RawDescriptor: file_bert_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bert_proto_goTypes,
		DependencyIndexes: file_bert_proto_depIdxs,
		EnumInfos:         file_bert_proto_enumTypes,
		MessageInfos:      file_bert_proto_msgTypes,
	}.Build()
	File_bert_proto = out.File
	file_bert_proto_rawDesc = nil
	file_bert_proto_goTypes = nil
	file_bert_proto_depIdxs = nil
}
