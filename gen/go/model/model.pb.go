// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: model/model.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Тип доступа
type AccessType int32

const (
	AccessType_PRIVATE            AccessType = 0 // Только владелец (значение по умолчанию)
	AccessType_FRIENDS            AccessType = 1 // Все друзья
	AccessType_PUBLIC             AccessType = 2 // Доступно всем
	AccessType_CUSTOM             AccessType = 3 // Только явно указанные в allowed_user_uuids
	AccessType_FRIENDS_AND_CUSTOM AccessType = 4 // Друзья + дополнительные пользователи
)

// Enum value maps for AccessType.
var (
	AccessType_name = map[int32]string{
		0: "PRIVATE",
		1: "FRIENDS",
		2: "PUBLIC",
		3: "CUSTOM",
		4: "FRIENDS_AND_CUSTOM",
	}
	AccessType_value = map[string]int32{
		"PRIVATE":            0,
		"FRIENDS":            1,
		"PUBLIC":             2,
		"CUSTOM":             3,
		"FRIENDS_AND_CUSTOM": 4,
	}
)

func (x AccessType) Enum() *AccessType {
	p := new(AccessType)
	*p = x
	return p
}

func (x AccessType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_model_proto_enumTypes[0].Descriptor()
}

func (AccessType) Type() protoreflect.EnumType {
	return &file_model_model_proto_enumTypes[0]
}

func (x AccessType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessType.Descriptor instead.
func (AccessType) EnumDescriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{0}
}

type File struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Uuid             string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                                            // Уникальный ID файла
	OwnerUuid        string                 `protobuf:"bytes,2,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`                                 // Владелец
	Name             *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`                                                      // Имя файла
	Path             *string                `protobuf:"bytes,4,opt,name=path,proto3,oneof" json:"path,omitempty"`                                                      // Путь в хранилище
	Size             *uint64                `protobuf:"varint,5,opt,name=size,proto3,oneof" json:"size,omitempty"`                                                     // Размер в байтах
	MimeType         *string                `protobuf:"bytes,6,opt,name=mime_type,json=mimeType,proto3,oneof" json:"mime_type,omitempty"`                              // MIME-тип (например, "image/png")
	UploadedAt       *int64                 `protobuf:"varint,7,opt,name=uploaded_at,json=uploadedAt,proto3,oneof" json:"uploaded_at,omitempty"`                       // Дата загрузки (Unix timestamp)
	IsPublic         *bool                  `protobuf:"varint,8,opt,name=is_public,json=isPublic,proto3,oneof" json:"is_public,omitempty"`                             // Доступен ли файл публично
	AccessType       *AccessType            `protobuf:"varint,9,opt,name=access_type,json=accessType,proto3,enum=model.AccessType,oneof" json:"access_type,omitempty"` // Тип доступа
	AllowedUserUuids []string               `protobuf:"bytes,10,rep,name=allowed_user_uuids,json=allowedUserUuids,proto3" json:"allowed_user_uuids,omitempty"`         // Список UUID пользователей, которым разрешен доступ
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_model_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *File) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *File) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *File) GetMimeType() string {
	if x != nil && x.MimeType != nil {
		return *x.MimeType
	}
	return ""
}

func (x *File) GetUploadedAt() int64 {
	if x != nil && x.UploadedAt != nil {
		return *x.UploadedAt
	}
	return 0
}

func (x *File) GetIsPublic() bool {
	if x != nil && x.IsPublic != nil {
		return *x.IsPublic
	}
	return false
}

func (x *File) GetAccessType() AccessType {
	if x != nil && x.AccessType != nil {
		return *x.AccessType
	}
	return AccessType_PRIVATE
}

func (x *File) GetAllowedUserUuids() []string {
	if x != nil {
		return x.AllowedUserUuids
	}
	return nil
}

// Статус операции
type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         *string                `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_model_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Status) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_model_model_proto protoreflect.FileDescriptor

var file_model_model_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xac, 0x03, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x06, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69,
	0x64, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2a, 0x56, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x68, 0x69, 0x43, 0x75, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_model_proto_rawDescOnce sync.Once
	file_model_model_proto_rawDescData []byte
)

func file_model_model_proto_rawDescGZIP() []byte {
	file_model_model_proto_rawDescOnce.Do(func() {
		file_model_model_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_model_proto_rawDesc), len(file_model_model_proto_rawDesc)))
	})
	return file_model_model_proto_rawDescData
}

var file_model_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_model_proto_goTypes = []any{
	(AccessType)(0), // 0: model.AccessType
	(*File)(nil),    // 1: model.File
	(*Status)(nil),  // 2: model.Status
}
var file_model_model_proto_depIdxs = []int32{
	0, // 0: model.File.access_type:type_name -> model.AccessType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_model_proto_init() }
func file_model_model_proto_init() {
	if File_model_model_proto != nil {
		return
	}
	file_model_model_proto_msgTypes[0].OneofWrappers = []any{}
	file_model_model_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_model_proto_rawDesc), len(file_model_model_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_proto_goTypes,
		DependencyIndexes: file_model_model_proto_depIdxs,
		EnumInfos:         file_model_model_proto_enumTypes,
		MessageInfos:      file_model_model_proto_msgTypes,
	}.Build()
	File_model_model_proto = out.File
	file_model_model_proto_goTypes = nil
	file_model_model_proto_depIdxs = nil
}
