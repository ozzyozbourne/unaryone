// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: xlsxserver.proto

package pbout

import (
	schemas "github.com/ozzyozbourne/unaryone/pbout/schemas"
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

type XlsxValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string             `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	SheetName string             `protobuf:"bytes,2,opt,name=sheetName,proto3" json:"sheetName,omitempty"`
	Cols      []*schemas.Columns `protobuf:"bytes,3,rep,name=cols,proto3" json:"cols,omitempty"`
}

func (x *XlsxValues) Reset() {
	*x = XlsxValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlsxserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XlsxValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XlsxValues) ProtoMessage() {}

func (x *XlsxValues) ProtoReflect() protoreflect.Message {
	mi := &file_xlsxserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XlsxValues.ProtoReflect.Descriptor instead.
func (*XlsxValues) Descriptor() ([]byte, []int) {
	return file_xlsxserver_proto_rawDescGZIP(), []int{0}
}

func (x *XlsxValues) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *XlsxValues) GetSheetName() string {
	if x != nil {
		return x.SheetName
	}
	return ""
}

func (x *XlsxValues) GetCols() []*schemas.Columns {
	if x != nil {
		return x.Cols
	}
	return nil
}

type XlsxValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res uint32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *XlsxValuesResponse) Reset() {
	*x = XlsxValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlsxserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XlsxValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XlsxValuesResponse) ProtoMessage() {}

func (x *XlsxValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xlsxserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XlsxValuesResponse.ProtoReflect.Descriptor instead.
func (*XlsxValuesResponse) Descriptor() ([]byte, []int) {
	return file_xlsxserver_proto_rawDescGZIP(), []int{1}
}

func (x *XlsxValuesResponse) GetRes() uint32 {
	if x != nil {
		return x.Res
	}
	return 0
}

type GetValRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *GetValRequest) Reset() {
	*x = GetValRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlsxserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValRequest) ProtoMessage() {}

func (x *GetValRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xlsxserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValRequest.ProtoReflect.Descriptor instead.
func (*GetValRequest) Descriptor() ([]byte, []int) {
	return file_xlsxserver_proto_rawDescGZIP(), []int{2}
}

func (x *GetValRequest) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_xlsxserver_proto protoreflect.FileDescriptor

var file_xlsxserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x78, 0x6c, 0x73, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x6f, 0x75, 0x74, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6b, 0x0a, 0x0a, 0x58, 0x6c, 0x73, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x22, 0x26, 0x0a, 0x12,
	0x58, 0x6c, 0x73, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x72, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x85, 0x01, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65,
	0x49, 0x6e, 0x47, 0x52, 0x50, 0x43, 0x12, 0x3d, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x6f, 0x75, 0x74, 0x2e,
	0x58, 0x6c, 0x73, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x6f,
	0x75, 0x74, 0x2e, 0x58, 0x6c, 0x73, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x58, 0x6c, 0x73, 0x78,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x6f, 0x75, 0x74, 0x2e, 0x58, 0x6c, 0x73, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42,
	0x3a, 0x0a, 0x0e, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x62, 0x6f, 0x75,
	0x74, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x7a, 0x7a, 0x79, 0x6f, 0x7a, 0x62, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x2f, 0x75, 0x6e, 0x61,
	0x72, 0x79, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_xlsxserver_proto_rawDescOnce sync.Once
	file_xlsxserver_proto_rawDescData = file_xlsxserver_proto_rawDesc
)

func file_xlsxserver_proto_rawDescGZIP() []byte {
	file_xlsxserver_proto_rawDescOnce.Do(func() {
		file_xlsxserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_xlsxserver_proto_rawDescData)
	})
	return file_xlsxserver_proto_rawDescData
}

var file_xlsxserver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_xlsxserver_proto_goTypes = []interface{}{
	(*XlsxValues)(nil),         // 0: pbout.XlsxValues
	(*XlsxValuesResponse)(nil), // 1: pbout.XlsxValuesResponse
	(*GetValRequest)(nil),      // 2: pbout.GetValRequest
	(*schemas.Columns)(nil),    // 3: sample.columns
}
var file_xlsxserver_proto_depIdxs = []int32{
	3, // 0: pbout.XlsxValues.cols:type_name -> sample.columns
	0, // 1: pbout.SaveInGRPC.PersistValues:input_type -> pbout.XlsxValues
	2, // 2: pbout.SaveInGRPC.GetXlsxValues:input_type -> pbout.GetValRequest
	1, // 3: pbout.SaveInGRPC.PersistValues:output_type -> pbout.XlsxValuesResponse
	0, // 4: pbout.SaveInGRPC.GetXlsxValues:output_type -> pbout.XlsxValues
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xlsxserver_proto_init() }
func file_xlsxserver_proto_init() {
	if File_xlsxserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xlsxserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XlsxValues); i {
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
		file_xlsxserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XlsxValuesResponse); i {
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
		file_xlsxserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValRequest); i {
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
			RawDescriptor: file_xlsxserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xlsxserver_proto_goTypes,
		DependencyIndexes: file_xlsxserver_proto_depIdxs,
		MessageInfos:      file_xlsxserver_proto_msgTypes,
	}.Build()
	File_xlsxserver_proto = out.File
	file_xlsxserver_proto_rawDesc = nil
	file_xlsxserver_proto_goTypes = nil
	file_xlsxserver_proto_depIdxs = nil
}
