// Code generated by protoc-gen-go. DO NOT EDIT.
// source: leaderboard/leaderboard.proto

package leaderboard

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Contest struct {
	ContestId            string   `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contest) Reset()         { *m = Contest{} }
func (m *Contest) String() string { return proto.CompactTextString(m) }
func (*Contest) ProtoMessage()    {}
func (*Contest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82336dd4a2bae2c7, []int{0}
}

func (m *Contest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contest.Unmarshal(m, b)
}
func (m *Contest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contest.Marshal(b, m, deterministic)
}
func (m *Contest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contest.Merge(m, src)
}
func (m *Contest) XXX_Size() int {
	return xxx_messageInfo_Contest.Size(m)
}
func (m *Contest) XXX_DiscardUnknown() {
	xxx_messageInfo_Contest.DiscardUnknown(m)
}

var xxx_messageInfo_Contest proto.InternalMessageInfo

func (m *Contest) GetContestId() string {
	if m != nil {
		return m.ContestId
	}
	return ""
}

type Participants struct {
	Participant          []*Participant `protobuf:"bytes,2,rep,name=participant,proto3" json:"participant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Participants) Reset()         { *m = Participants{} }
func (m *Participants) String() string { return proto.CompactTextString(m) }
func (*Participants) ProtoMessage()    {}
func (*Participants) Descriptor() ([]byte, []int) {
	return fileDescriptor_82336dd4a2bae2c7, []int{1}
}

func (m *Participants) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Participants.Unmarshal(m, b)
}
func (m *Participants) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Participants.Marshal(b, m, deterministic)
}
func (m *Participants) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participants.Merge(m, src)
}
func (m *Participants) XXX_Size() int {
	return xxx_messageInfo_Participants.Size(m)
}
func (m *Participants) XXX_DiscardUnknown() {
	xxx_messageInfo_Participants.DiscardUnknown(m)
}

var xxx_messageInfo_Participants proto.InternalMessageInfo

func (m *Participants) GetParticipant() []*Participant {
	if m != nil {
		return m.Participant
	}
	return nil
}

type Participant struct {
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Rating               int32    `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	NoOfQuestions        int32    `protobuf:"varint,6,opt,name=no_of_questions,json=noOfQuestions,proto3" json:"no_of_questions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Participant) Reset()         { *m = Participant{} }
func (m *Participant) String() string { return proto.CompactTextString(m) }
func (*Participant) ProtoMessage()    {}
func (*Participant) Descriptor() ([]byte, []int) {
	return fileDescriptor_82336dd4a2bae2c7, []int{2}
}

func (m *Participant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Participant.Unmarshal(m, b)
}
func (m *Participant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Participant.Marshal(b, m, deterministic)
}
func (m *Participant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participant.Merge(m, src)
}
func (m *Participant) XXX_Size() int {
	return xxx_messageInfo_Participant.Size(m)
}
func (m *Participant) XXX_DiscardUnknown() {
	xxx_messageInfo_Participant.DiscardUnknown(m)
}

var xxx_messageInfo_Participant proto.InternalMessageInfo

func (m *Participant) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Participant) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Participant) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Participant) GetNoOfQuestions() int32 {
	if m != nil {
		return m.NoOfQuestions
	}
	return 0
}

func init() {
	proto.RegisterType((*Contest)(nil), "Contest")
	proto.RegisterType((*Participants)(nil), "Participants")
	proto.RegisterType((*Participant)(nil), "Participant")
}

func init() { proto.RegisterFile("leaderboard/leaderboard.proto", fileDescriptor_82336dd4a2bae2c7) }

var fileDescriptor_82336dd4a2bae2c7 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0xb5, 0xdb, 0x76, 0xb6, 0x55, 0xc8, 0x41, 0x43, 0xa1, 0xb0, 0xec, 0x41, 0x72,
	0xda, 0x4a, 0xbd, 0xf7, 0xe2, 0x41, 0x0a, 0x82, 0xba, 0x7f, 0x60, 0x49, 0x9b, 0x54, 0x02, 0x9a,
	0x59, 0x93, 0xe9, 0x1f, 0xf0, 0x97, 0xcb, 0xc6, 0xb5, 0x9b, 0xdb, 0xfb, 0xde, 0x3c, 0x78, 0xcc,
	0x83, 0xd5, 0xa7, 0x51, 0xda, 0xf8, 0x3d, 0x2a, 0xaf, 0xd7, 0x89, 0xae, 0x5a, 0x8f, 0x84, 0xa5,
	0x84, 0xc9, 0x13, 0x3a, 0x32, 0x81, 0xf8, 0x0a, 0xe0, 0xf0, 0x27, 0x1b, 0xab, 0x05, 0x2b, 0x98,
	0x9c, 0xd5, 0xb3, 0xde, 0xd9, 0xe9, 0x72, 0x0b, 0xf3, 0x37, 0xe5, 0xc9, 0x1e, 0x6c, 0xab, 0x1c,
	0x05, 0x5e, 0x41, 0xde, 0x0e, 0x2c, 0x2e, 0x8b, 0x91, 0xcc, 0x37, 0xf3, 0x2a, 0xc9, 0xd4, 0x69,
	0xa0, 0xfc, 0x61, 0x90, 0x27, 0x47, 0x7e, 0x07, 0x93, 0x53, 0x30, 0xbe, 0xeb, 0x1a, 0xc5, 0xae,
	0xac, 0xc3, 0x9d, 0xe6, 0x4b, 0x98, 0x76, 0xca, 0xa9, 0x2f, 0x23, 0xae, 0xe2, 0xe5, 0xcc, 0xfc,
	0x16, 0x32, 0xaf, 0xc8, 0xba, 0x0f, 0x31, 0x2e, 0x98, 0x1c, 0xd7, 0x3d, 0xf1, 0x7b, 0xb8, 0x71,
	0xd8, 0xe0, 0xb1, 0xf9, 0x3e, 0x99, 0x40, 0x16, 0x5d, 0x10, 0x59, 0x0c, 0x2c, 0x1c, 0xbe, 0x1e,
	0xdf, 0xff, 0xcd, 0xcd, 0x16, 0xf2, 0x97, 0x61, 0x03, 0xbe, 0x86, 0xeb, 0x67, 0x43, 0xa9, 0x33,
	0xad, 0xfa, 0x39, 0x96, 0x8b, 0xf4, 0x95, 0x50, 0x5e, 0x48, 0xf6, 0xc0, 0xf6, 0x59, 0x5c, 0xed,
	0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xc9, 0x58, 0x6b, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LeaderboardClient is the client API for Leaderboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeaderboardClient interface {
	// Returns Participants in sorted order for each Contest
	GetLeaderboard(ctx context.Context, opts ...grpc.CallOption) (Leaderboard_GetLeaderboardClient, error)
}

type leaderboardClient struct {
	cc *grpc.ClientConn
}

func NewLeaderboardClient(cc *grpc.ClientConn) LeaderboardClient {
	return &leaderboardClient{cc}
}

func (c *leaderboardClient) GetLeaderboard(ctx context.Context, opts ...grpc.CallOption) (Leaderboard_GetLeaderboardClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Leaderboard_serviceDesc.Streams[0], "/Leaderboard/GetLeaderboard", opts...)
	if err != nil {
		return nil, err
	}
	x := &leaderboardGetLeaderboardClient{stream}
	return x, nil
}

type Leaderboard_GetLeaderboardClient interface {
	Send(*Contest) error
	Recv() (*Participants, error)
	grpc.ClientStream
}

type leaderboardGetLeaderboardClient struct {
	grpc.ClientStream
}

func (x *leaderboardGetLeaderboardClient) Send(m *Contest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *leaderboardGetLeaderboardClient) Recv() (*Participants, error) {
	m := new(Participants)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LeaderboardServer is the server API for Leaderboard service.
type LeaderboardServer interface {
	// Returns Participants in sorted order for each Contest
	GetLeaderboard(Leaderboard_GetLeaderboardServer) error
}

func RegisterLeaderboardServer(s *grpc.Server, srv LeaderboardServer) {
	s.RegisterService(&_Leaderboard_serviceDesc, srv)
}

func _Leaderboard_GetLeaderboard_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LeaderboardServer).GetLeaderboard(&leaderboardGetLeaderboardServer{stream})
}

type Leaderboard_GetLeaderboardServer interface {
	Send(*Participants) error
	Recv() (*Contest, error)
	grpc.ServerStream
}

type leaderboardGetLeaderboardServer struct {
	grpc.ServerStream
}

func (x *leaderboardGetLeaderboardServer) Send(m *Participants) error {
	return x.ServerStream.SendMsg(m)
}

func (x *leaderboardGetLeaderboardServer) Recv() (*Contest, error) {
	m := new(Contest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Leaderboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Leaderboard",
	HandlerType: (*LeaderboardServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLeaderboard",
			Handler:       _Leaderboard_GetLeaderboard_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "leaderboard/leaderboard.proto",
}