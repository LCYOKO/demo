// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: book.proto

package book

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILED  Status = 1
	Status_DEAD    Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
		2: "DEAD",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
		"DEAD":    2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_book_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_book_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

type BookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books   []*Book          `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	Status  Status           `protobuf:"varint,2,opt,name=status,proto3,enum=book.Status" json:"status,omitempty"`
	BookMap map[string]*Book `protobuf:"bytes,3,rep,name=bookMap,proto3" json:"bookMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *BookResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *BookResponse) GetBookMap() map[string]*Book {
	if x != nil {
		return x.BookMap
	}
	return nil
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d,
	0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd9, 0x01,
	0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x4d, 0x61,
	0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x4d, 0x61,
	0x70, 0x1a, 0x46, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x32, 0xe5, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x31, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x34, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x33, 0x12, 0x11, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x36, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x34, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_book_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: book.Status
	(*Book)(nil),                   // 1: book.Book
	(*BookRequest)(nil),            // 2: book.BookRequest
	(*BookResponse)(nil),           // 3: book.BookResponse
	nil,                            // 4: book.BookResponse.BookMapEntry
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_book_proto_depIdxs = []int32{
	5, // 0: book.BookRequest.name:type_name -> google.protobuf.StringValue
	1, // 1: book.BookResponse.books:type_name -> book.Book
	0, // 2: book.BookResponse.status:type_name -> book.Status
	4, // 3: book.BookResponse.bookMap:type_name -> book.BookResponse.BookMapEntry
	1, // 4: book.BookResponse.BookMapEntry.value:type_name -> book.Book
	2, // 5: book.BookService.getBooks1:input_type -> book.BookRequest
	2, // 6: book.BookService.getBooks2:input_type -> book.BookRequest
	2, // 7: book.BookService.getBooks3:input_type -> book.BookRequest
	2, // 8: book.BookService.getBooks4:input_type -> book.BookRequest
	3, // 9: book.BookService.getBooks1:output_type -> book.BookResponse
	3, // 10: book.BookService.getBooks2:output_type -> book.BookResponse
	3, // 11: book.BookService.getBooks3:output_type -> book.BookResponse
	3, // 12: book.BookService.getBooks4:output_type -> book.BookResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookResponse); i {
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
	file_book_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		EnumInfos:         file_book_proto_enumTypes,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}