// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defs.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	defs.proto

It has these top-level messages:
	Reservation
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Reservation struct {
	Email        string  `protobuf:"bytes,1,opt,name=Email" json:"Email,omitempty"`
	Venue        string  `protobuf:"bytes,2,opt,name=Venue" json:"Venue,omitempty"`
	Code         string  `protobuf:"bytes,3,opt,name=Code" json:"Code,omitempty"`
	Instructions string  `protobuf:"bytes,4,opt,name=Instructions" json:"Instructions,omitempty"`
	Time         float64 `protobuf:"fixed64,5,opt,name=Time" json:"Time,omitempty"`
}

func (m *Reservation) Reset()                    { *m = Reservation{} }
func (m *Reservation) String() string            { return proto.CompactTextString(m) }
func (*Reservation) ProtoMessage()               {}
func (*Reservation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Reservation) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Reservation) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *Reservation) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Reservation) GetInstructions() string {
	if m != nil {
		return m.Instructions
	}
	return ""
}

func (m *Reservation) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*Reservation)(nil), "main.Reservation")
}

func init() { proto.RegisterFile("defs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x4d, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x6a, 0x65, 0xe4,
	0xe2, 0x0e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x13, 0x12, 0xe1, 0x62,
	0x75, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x40, 0xa2,
	0x61, 0xa9, 0x79, 0xa5, 0xa9, 0x12, 0x4c, 0x10, 0x51, 0x30, 0x47, 0x48, 0x88, 0x8b, 0xc5, 0x39,
	0x3f, 0x25, 0x55, 0x82, 0x19, 0x2c, 0x08, 0x66, 0x0b, 0x29, 0x71, 0xf1, 0x78, 0xe6, 0x15, 0x97,
	0x14, 0x95, 0x26, 0x83, 0x8c, 0x2b, 0x96, 0x60, 0x01, 0xcb, 0xa1, 0x88, 0x81, 0xf4, 0x85, 0x64,
	0xe6, 0xa6, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0x30, 0x06, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0x47, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27, 0x19, 0x55, 0xa3, 0xa2, 0x00, 0x00, 0x00,
}