// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: node/highway/v1/highway.proto

// Package Highway is used for defining a Highway node and its properties.

package highway

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_node_highway_v1_highway_proto protoreflect.FileDescriptor

var file_node_highway_v1_highway_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x1d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93, 0x18,
	0x0a, 0x0e, 0x48, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6e, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x76, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7a, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x72, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x7a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x7a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7c, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x76, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x76, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x76, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68,
	0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x76, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x70, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x62, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68,
	0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x30, 0x01, 0x12, 0x78, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68,
	0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x30,
	0x01, 0x12, 0x68, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x20, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x30, 0x01, 0x12, 0x6e, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68,
	0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x64, 0x2f, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69,
	0x64, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_node_highway_v1_highway_proto_goTypes = []interface{}{
	(*AccessNameRequest)(nil),       // 0: node.highway.v1.AccessNameRequest
	(*RegisterNameRequest)(nil),     // 1: node.highway.v1.RegisterNameRequest
	(*UpdateNameRequest)(nil),       // 2: node.highway.v1.UpdateNameRequest
	(*AccessServiceRequest)(nil),    // 3: node.highway.v1.AccessServiceRequest
	(*RegisterServiceRequest)(nil),  // 4: node.highway.v1.RegisterServiceRequest
	(*UpdateServiceRequest)(nil),    // 5: node.highway.v1.UpdateServiceRequest
	(*CreateChannelRequest)(nil),    // 6: node.highway.v1.CreateChannelRequest
	(*ReadChannelRequest)(nil),      // 7: node.highway.v1.ReadChannelRequest
	(*UpdateChannelRequest)(nil),    // 8: node.highway.v1.UpdateChannelRequest
	(*DeleteChannelRequest)(nil),    // 9: node.highway.v1.DeleteChannelRequest
	(*ListenChannelRequest)(nil),    // 10: node.highway.v1.ListenChannelRequest
	(*CreateBucketRequest)(nil),     // 11: node.highway.v1.CreateBucketRequest
	(*ReadBucketRequest)(nil),       // 12: node.highway.v1.ReadBucketRequest
	(*UpdateBucketRequest)(nil),     // 13: node.highway.v1.UpdateBucketRequest
	(*DeleteBucketRequest)(nil),     // 14: node.highway.v1.DeleteBucketRequest
	(*ListenBucketRequest)(nil),     // 15: node.highway.v1.ListenBucketRequest
	(*CreateObjectRequest)(nil),     // 16: node.highway.v1.CreateObjectRequest
	(*ReadObjectRequest)(nil),       // 17: node.highway.v1.ReadObjectRequest
	(*UpdateObjectRequest)(nil),     // 18: node.highway.v1.UpdateObjectRequest
	(*DeleteObjectRequest)(nil),     // 19: node.highway.v1.DeleteObjectRequest
	(*UploadBlobRequest)(nil),       // 20: node.highway.v1.UploadBlobRequest
	(*DownloadBlobRequest)(nil),     // 21: node.highway.v1.DownloadBlobRequest
	(*SyncBlobRequest)(nil),         // 22: node.highway.v1.SyncBlobRequest
	(*DeleteBlobRequest)(nil),       // 23: node.highway.v1.DeleteBlobRequest
	(*ParseDidRequest)(nil),         // 24: node.highway.v1.ParseDidRequest
	(*ResolveDidRequest)(nil),       // 25: node.highway.v1.ResolveDidRequest
	(*AccessNameResponse)(nil),      // 26: node.highway.v1.AccessNameResponse
	(*RegisterNameResponse)(nil),    // 27: node.highway.v1.RegisterNameResponse
	(*UpdateNameResponse)(nil),      // 28: node.highway.v1.UpdateNameResponse
	(*AccessServiceResponse)(nil),   // 29: node.highway.v1.AccessServiceResponse
	(*RegisterServiceResponse)(nil), // 30: node.highway.v1.RegisterServiceResponse
	(*UpdateServiceResponse)(nil),   // 31: node.highway.v1.UpdateServiceResponse
	(*CreateChannelResponse)(nil),   // 32: node.highway.v1.CreateChannelResponse
	(*ReadChannelResponse)(nil),     // 33: node.highway.v1.ReadChannelResponse
	(*UpdateChannelResponse)(nil),   // 34: node.highway.v1.UpdateChannelResponse
	(*DeleteChannelResponse)(nil),   // 35: node.highway.v1.DeleteChannelResponse
	(*ListenChannelResponse)(nil),   // 36: node.highway.v1.ListenChannelResponse
	(*CreateBucketResponse)(nil),    // 37: node.highway.v1.CreateBucketResponse
	(*ReadBucketResponse)(nil),      // 38: node.highway.v1.ReadBucketResponse
	(*UpdateBucketResponse)(nil),    // 39: node.highway.v1.UpdateBucketResponse
	(*DeleteBucketResponse)(nil),    // 40: node.highway.v1.DeleteBucketResponse
	(*ListenBucketResponse)(nil),    // 41: node.highway.v1.ListenBucketResponse
	(*CreateObjectResponse)(nil),    // 42: node.highway.v1.CreateObjectResponse
	(*ReadObjectResponse)(nil),      // 43: node.highway.v1.ReadObjectResponse
	(*UpdateObjectResponse)(nil),    // 44: node.highway.v1.UpdateObjectResponse
	(*DeleteObjectResponse)(nil),    // 45: node.highway.v1.DeleteObjectResponse
	(*UploadBlobResponse)(nil),      // 46: node.highway.v1.UploadBlobResponse
	(*DownloadBlobResponse)(nil),    // 47: node.highway.v1.DownloadBlobResponse
	(*SyncBlobResponse)(nil),        // 48: node.highway.v1.SyncBlobResponse
	(*DeleteBlobResponse)(nil),      // 49: node.highway.v1.DeleteBlobResponse
	(*ParseDidResponse)(nil),        // 50: node.highway.v1.ParseDidResponse
	(*ResolveDidResponse)(nil),      // 51: node.highway.v1.ResolveDidResponse
}
var file_node_highway_v1_highway_proto_depIdxs = []int32{
	0,  // 0: node.highway.v1.HighwayService.AccessName:input_type -> node.highway.v1.AccessNameRequest
	1,  // 1: node.highway.v1.HighwayService.RegisterName:input_type -> node.highway.v1.RegisterNameRequest
	2,  // 2: node.highway.v1.HighwayService.UpdateName:input_type -> node.highway.v1.UpdateNameRequest
	3,  // 3: node.highway.v1.HighwayService.AccessService:input_type -> node.highway.v1.AccessServiceRequest
	4,  // 4: node.highway.v1.HighwayService.RegisterService:input_type -> node.highway.v1.RegisterServiceRequest
	5,  // 5: node.highway.v1.HighwayService.UpdateService:input_type -> node.highway.v1.UpdateServiceRequest
	6,  // 6: node.highway.v1.HighwayService.CreateChannel:input_type -> node.highway.v1.CreateChannelRequest
	7,  // 7: node.highway.v1.HighwayService.ReadChannel:input_type -> node.highway.v1.ReadChannelRequest
	8,  // 8: node.highway.v1.HighwayService.UpdateChannel:input_type -> node.highway.v1.UpdateChannelRequest
	9,  // 9: node.highway.v1.HighwayService.DeleteChannel:input_type -> node.highway.v1.DeleteChannelRequest
	10, // 10: node.highway.v1.HighwayService.ListenChannel:input_type -> node.highway.v1.ListenChannelRequest
	11, // 11: node.highway.v1.HighwayService.CreateBucket:input_type -> node.highway.v1.CreateBucketRequest
	12, // 12: node.highway.v1.HighwayService.ReadBucket:input_type -> node.highway.v1.ReadBucketRequest
	13, // 13: node.highway.v1.HighwayService.UpdateBucket:input_type -> node.highway.v1.UpdateBucketRequest
	14, // 14: node.highway.v1.HighwayService.DeleteBucket:input_type -> node.highway.v1.DeleteBucketRequest
	15, // 15: node.highway.v1.HighwayService.ListenBucket:input_type -> node.highway.v1.ListenBucketRequest
	16, // 16: node.highway.v1.HighwayService.CreateObject:input_type -> node.highway.v1.CreateObjectRequest
	17, // 17: node.highway.v1.HighwayService.ReadObject:input_type -> node.highway.v1.ReadObjectRequest
	18, // 18: node.highway.v1.HighwayService.UpdateObject:input_type -> node.highway.v1.UpdateObjectRequest
	19, // 19: node.highway.v1.HighwayService.DeleteObject:input_type -> node.highway.v1.DeleteObjectRequest
	20, // 20: node.highway.v1.HighwayService.UploadBlob:input_type -> node.highway.v1.UploadBlobRequest
	21, // 21: node.highway.v1.HighwayService.DownloadBlob:input_type -> node.highway.v1.DownloadBlobRequest
	22, // 22: node.highway.v1.HighwayService.SyncBlob:input_type -> node.highway.v1.SyncBlobRequest
	23, // 23: node.highway.v1.HighwayService.DeleteBlob:input_type -> node.highway.v1.DeleteBlobRequest
	24, // 24: node.highway.v1.HighwayService.ParseDid:input_type -> node.highway.v1.ParseDidRequest
	25, // 25: node.highway.v1.HighwayService.ResolveDid:input_type -> node.highway.v1.ResolveDidRequest
	26, // 26: node.highway.v1.HighwayService.AccessName:output_type -> node.highway.v1.AccessNameResponse
	27, // 27: node.highway.v1.HighwayService.RegisterName:output_type -> node.highway.v1.RegisterNameResponse
	28, // 28: node.highway.v1.HighwayService.UpdateName:output_type -> node.highway.v1.UpdateNameResponse
	29, // 29: node.highway.v1.HighwayService.AccessService:output_type -> node.highway.v1.AccessServiceResponse
	30, // 30: node.highway.v1.HighwayService.RegisterService:output_type -> node.highway.v1.RegisterServiceResponse
	31, // 31: node.highway.v1.HighwayService.UpdateService:output_type -> node.highway.v1.UpdateServiceResponse
	32, // 32: node.highway.v1.HighwayService.CreateChannel:output_type -> node.highway.v1.CreateChannelResponse
	33, // 33: node.highway.v1.HighwayService.ReadChannel:output_type -> node.highway.v1.ReadChannelResponse
	34, // 34: node.highway.v1.HighwayService.UpdateChannel:output_type -> node.highway.v1.UpdateChannelResponse
	35, // 35: node.highway.v1.HighwayService.DeleteChannel:output_type -> node.highway.v1.DeleteChannelResponse
	36, // 36: node.highway.v1.HighwayService.ListenChannel:output_type -> node.highway.v1.ListenChannelResponse
	37, // 37: node.highway.v1.HighwayService.CreateBucket:output_type -> node.highway.v1.CreateBucketResponse
	38, // 38: node.highway.v1.HighwayService.ReadBucket:output_type -> node.highway.v1.ReadBucketResponse
	39, // 39: node.highway.v1.HighwayService.UpdateBucket:output_type -> node.highway.v1.UpdateBucketResponse
	40, // 40: node.highway.v1.HighwayService.DeleteBucket:output_type -> node.highway.v1.DeleteBucketResponse
	41, // 41: node.highway.v1.HighwayService.ListenBucket:output_type -> node.highway.v1.ListenBucketResponse
	42, // 42: node.highway.v1.HighwayService.CreateObject:output_type -> node.highway.v1.CreateObjectResponse
	43, // 43: node.highway.v1.HighwayService.ReadObject:output_type -> node.highway.v1.ReadObjectResponse
	44, // 44: node.highway.v1.HighwayService.UpdateObject:output_type -> node.highway.v1.UpdateObjectResponse
	45, // 45: node.highway.v1.HighwayService.DeleteObject:output_type -> node.highway.v1.DeleteObjectResponse
	46, // 46: node.highway.v1.HighwayService.UploadBlob:output_type -> node.highway.v1.UploadBlobResponse
	47, // 47: node.highway.v1.HighwayService.DownloadBlob:output_type -> node.highway.v1.DownloadBlobResponse
	48, // 48: node.highway.v1.HighwayService.SyncBlob:output_type -> node.highway.v1.SyncBlobResponse
	49, // 49: node.highway.v1.HighwayService.DeleteBlob:output_type -> node.highway.v1.DeleteBlobResponse
	50, // 50: node.highway.v1.HighwayService.ParseDid:output_type -> node.highway.v1.ParseDidResponse
	51, // 51: node.highway.v1.HighwayService.ResolveDid:output_type -> node.highway.v1.ResolveDidResponse
	26, // [26:52] is the sub-list for method output_type
	0,  // [0:26] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_node_highway_v1_highway_proto_init() }
func file_node_highway_v1_highway_proto_init() {
	if File_node_highway_v1_highway_proto != nil {
		return
	}
	file_node_highway_v1_request_proto_init()
	file_node_highway_v1_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_highway_v1_highway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_highway_v1_highway_proto_goTypes,
		DependencyIndexes: file_node_highway_v1_highway_proto_depIdxs,
	}.Build()
	File_node_highway_v1_highway_proto = out.File
	file_node_highway_v1_highway_proto_rawDesc = nil
	file_node_highway_v1_highway_proto_goTypes = nil
	file_node_highway_v1_highway_proto_depIdxs = nil
}