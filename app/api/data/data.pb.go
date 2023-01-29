// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: app/api/proto/data.proto

package data

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

type Bird_Type int32

const (
	Bird_Type_SHILL   Bird_Type = 0
	Bird_Type_COMPANY Bird_Type = 1
)

// Enum value maps for Bird_Type.
var (
	Bird_Type_name = map[int32]string{
		0: "SHILL",
		1: "COMPANY",
	}
	Bird_Type_value = map[string]int32{
		"SHILL":   0,
		"COMPANY": 1,
	}
)

func (x Bird_Type) Enum() *Bird_Type {
	p := new(Bird_Type)
	*p = x
	return p
}

func (x Bird_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bird_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_app_api_proto_data_proto_enumTypes[0].Descriptor()
}

func (Bird_Type) Type() protoreflect.EnumType {
	return &file_app_api_proto_data_proto_enumTypes[0]
}

func (x Bird_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bird_Type.Descriptor instead.
func (Bird_Type) EnumDescriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{0}
}

type Bird struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      Bird_Type `protobuf:"varint,3,opt,name=type,proto3,enum=data.Bird_Type" json:"type,omitempty"`
	BirdFk    int32     `protobuf:"varint,4,opt,name=bird_fk,json=birdFk,proto3" json:"bird_fk,omitempty"`
	Score     float32   `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	NPositive int32     `protobuf:"varint,6,opt,name=n_positive,json=nPositive,proto3" json:"n_positive,omitempty"`
	NNegative int32     `protobuf:"varint,7,opt,name=n_negative,json=nNegative,proto3" json:"n_negative,omitempty"`
	ImgUrl    string    `protobuf:"bytes,8,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
}

func (x *Bird) Reset() {
	*x = Bird{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_api_proto_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bird) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bird) ProtoMessage() {}

func (x *Bird) ProtoReflect() protoreflect.Message {
	mi := &file_app_api_proto_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bird.ProtoReflect.Descriptor instead.
func (*Bird) Descriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{0}
}

func (x *Bird) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bird) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bird) GetType() Bird_Type {
	if x != nil {
		return x.Type
	}
	return Bird_Type_SHILL
}

func (x *Bird) GetBirdFk() int32 {
	if x != nil {
		return x.BirdFk
	}
	return 0
}

func (x *Bird) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Bird) GetNPositive() int32 {
	if x != nil {
		return x.NPositive
	}
	return 0
}

func (x *Bird) GetNNegative() int32 {
	if x != nil {
		return x.NNegative
	}
	return 0
}

func (x *Bird) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

type Birds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Birds []*Bird `protobuf:"bytes,1,rep,name=birds,proto3" json:"birds,omitempty"`
}

func (x *Birds) Reset() {
	*x = Birds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_api_proto_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Birds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Birds) ProtoMessage() {}

func (x *Birds) ProtoReflect() protoreflect.Message {
	mi := &file_app_api_proto_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Birds.ProtoReflect.Descriptor instead.
func (*Birds) Descriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{1}
}

func (x *Birds) GetBirds() []*Bird {
	if x != nil {
		return x.Birds
	}
	return nil
}

type TweetStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TweetStreamRequest) Reset() {
	*x = TweetStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_api_proto_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetStreamRequest) ProtoMessage() {}

func (x *TweetStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_api_proto_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetStreamRequest.ProtoReflect.Descriptor instead.
func (*TweetStreamRequest) Descriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{2}
}

type BirdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BirdsRequest) Reset() {
	*x = BirdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_api_proto_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BirdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BirdsRequest) ProtoMessage() {}

func (x *BirdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_api_proto_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BirdsRequest.ProtoReflect.Descriptor instead.
func (*BirdsRequest) Descriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{3}
}

type BirdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BirdId int32 `protobuf:"varint,1,opt,name=bird_id,json=birdId,proto3" json:"bird_id,omitempty"`
}

func (x *BirdRequest) Reset() {
	*x = BirdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_api_proto_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BirdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BirdRequest) ProtoMessage() {}

func (x *BirdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_api_proto_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BirdRequest.ProtoReflect.Descriptor instead.
func (*BirdRequest) Descriptor() ([]byte, []int) {
	return file_app_api_proto_data_proto_rawDescGZIP(), []int{4}
}

func (x *BirdRequest) GetBirdId() int32 {
	if x != nil {
		return x.BirdId
	}
	return 0
}

var File_app_api_proto_data_proto protoreflect.FileDescriptor

var file_app_api_proto_data_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xd5, 0x01, 0x0a, 0x04, 0x42, 0x69, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x42, 0x69, 0x72, 0x64, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x69, 0x72, 0x64, 0x5f, 0x66, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x69, 0x72, 0x64, 0x46, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x05, 0x42, 0x69, 0x72, 0x64,
	0x73, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x72, 0x64, 0x52, 0x05, 0x62, 0x69,
	0x72, 0x64, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x77, 0x65, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x42, 0x69, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0b, 0x42, 0x69, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x69, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x69, 0x72, 0x64, 0x49,
	0x64, 0x2a, 0x23, 0x0a, 0x09, 0x42, 0x69, 0x72, 0x64, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x48, 0x49, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x37, 0x0a, 0x0b, 0x54, 0x77, 0x65, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x42, 0x69, 0x72, 0x64, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42,
	0x69, 0x72, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x72, 0x64,
	0x73, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x72,
	0x64, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_api_proto_data_proto_rawDescOnce sync.Once
	file_app_api_proto_data_proto_rawDescData = file_app_api_proto_data_proto_rawDesc
)

func file_app_api_proto_data_proto_rawDescGZIP() []byte {
	file_app_api_proto_data_proto_rawDescOnce.Do(func() {
		file_app_api_proto_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_api_proto_data_proto_rawDescData)
	})
	return file_app_api_proto_data_proto_rawDescData
}

var file_app_api_proto_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_api_proto_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_api_proto_data_proto_goTypes = []interface{}{
	(Bird_Type)(0),             // 0: data.Bird_Type
	(*Bird)(nil),               // 1: data.Bird
	(*Birds)(nil),              // 2: data.Birds
	(*TweetStreamRequest)(nil), // 3: data.TweetStreamRequest
	(*BirdsRequest)(nil),       // 4: data.BirdsRequest
	(*BirdRequest)(nil),        // 5: data.BirdRequest
}
var file_app_api_proto_data_proto_depIdxs = []int32{
	0, // 0: data.Bird.type:type_name -> data.Bird_Type
	1, // 1: data.Birds.birds:type_name -> data.Bird
	3, // 2: data.Data.TweetStream:input_type -> data.TweetStreamRequest
	5, // 3: data.Data.BirdData:input_type -> data.BirdRequest
	4, // 4: data.Data.AllBirds:input_type -> data.BirdsRequest
	1, // 5: data.Data.TweetStream:output_type -> data.Bird
	1, // 6: data.Data.BirdData:output_type -> data.Bird
	2, // 7: data.Data.AllBirds:output_type -> data.Birds
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_api_proto_data_proto_init() }
func file_app_api_proto_data_proto_init() {
	if File_app_api_proto_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_api_proto_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bird); i {
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
		file_app_api_proto_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Birds); i {
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
		file_app_api_proto_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetStreamRequest); i {
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
		file_app_api_proto_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BirdsRequest); i {
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
		file_app_api_proto_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BirdRequest); i {
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
			RawDescriptor: file_app_api_proto_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_api_proto_data_proto_goTypes,
		DependencyIndexes: file_app_api_proto_data_proto_depIdxs,
		EnumInfos:         file_app_api_proto_data_proto_enumTypes,
		MessageInfos:      file_app_api_proto_data_proto_msgTypes,
	}.Build()
	File_app_api_proto_data_proto = out.File
	file_app_api_proto_data_proto_rawDesc = nil
	file_app_api_proto_data_proto_goTypes = nil
	file_app_api_proto_data_proto_depIdxs = nil
}
