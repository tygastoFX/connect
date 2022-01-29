// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: minekube/connect/v1alpha1/watch_service.proto

package connectv1alpha1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sending this endpoint adds the server to the list
	// of multiplexed servers for this endpoint name.
	//
	// It is allowed to send other messages before this.
	//
	// If this is message is not sent the stream will
	// be closed serverside after a timeout.
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Sending this message rejects a session proposed by
	// the WatchService. This message should be sent
	// to inform the WatchService that the server will not try
	// to make a take the proposed session. The only purpose of
	// this message is to provide quicker feedback to the player
	// that he will not be connected with an optional localized
	// reason. See https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto.
	// If the session is not rejected the watcher should establish
	// the connection for the proposed session.
	// If neither of these actions happen the proposal times out
	// out and the player receives a connection timeout error
	// indicating that the endpoint is currently unavailable.
	SessionRejection *SessionRejection `protobuf:"bytes,2,opt,name=session_rejection,json=sessionRejection,proto3" json:"session_rejection,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRequest) GetEndpoint() *Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *WatchRequest) GetSessionRejection() *SessionRejection {
	if x != nil {
		return x.SessionRejection
	}
	return nil
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The proposed session that intents to connect to the endpoint.
	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{1}
}

func (x *WatchResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// It is the server name that players connect by.
	// If there are multiple equally named endpoints
	// player connections are multiplexed.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{2}
}

func (x *Endpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SessionRejection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the proposed session.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The optional reason why the proposed session was rejected.
	// To specify a user facing localized message refer to
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	Reason *status.Status `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *SessionRejection) Reset() {
	*x = SessionRejection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRejection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRejection) ProtoMessage() {}

func (x *SessionRejection) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRejection.ProtoReflect.Descriptor instead.
func (*SessionRejection) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{3}
}

func (x *SessionRejection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SessionRejection) GetReason() *status.Status {
	if x != nil {
		return x.Reason
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique session id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The address of the tunnel service to establish the connection.
	TunnelServiceAddr string `protobuf:"bytes,2,opt,name=tunnel_service_addr,json=tunnelServiceAddr,proto3" json:"tunnel_service_addr,omitempty"`
	// The player that will be connected.
	Player *Player `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{4}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetTunnelServiceAddr() string {
	if x != nil {
		return x.TunnelServiceAddr
	}
	return ""
}

func (x *Session) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The profile of the player.
	Profile *GameProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	// The optional IP address of the player.
	// This field may be empty or is a fake address
	// generated for this particular player.
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetProfile() *GameProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Player) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type GameProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Minecraft UUID of the player.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The Minecraft name of the player.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The profile properties that may contain skin data and more.
	Properties []*GameProfileProperty `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *GameProfile) Reset() {
	*x = GameProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameProfile) ProtoMessage() {}

func (x *GameProfile) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameProfile.ProtoReflect.Descriptor instead.
func (*GameProfile) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{6}
}

func (x *GameProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GameProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameProfile) GetProperties() []*GameProfileProperty {
	if x != nil {
		return x.Properties
	}
	return nil
}

type GameProfileProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value of this property
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// The signature of this property.
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *GameProfileProperty) Reset() {
	*x = GameProfileProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameProfileProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameProfileProperty) ProtoMessage() {}

func (x *GameProfileProperty) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameProfileProperty.ProtoReflect.Descriptor instead.
func (*GameProfileProperty) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP(), []int{7}
}

func (x *GameProfileProperty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameProfileProperty) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GameProfileProperty) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_minekube_connect_v1alpha1_watch_service_proto protoreflect.FileDescriptor

var file_minekube_connect_v1alpha1_watch_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x4d, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1e,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e,
	0x0a, 0x10, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x84,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x13, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x6e, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x27, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0xfd, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x11, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x45, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x43, 0x58, 0xaa, 0x02, 0x19, 0x4d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x4d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x25, 0x4d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x5c, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x4d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_minekube_connect_v1alpha1_watch_service_proto_rawDescOnce sync.Once
	file_minekube_connect_v1alpha1_watch_service_proto_rawDescData = file_minekube_connect_v1alpha1_watch_service_proto_rawDesc
)

func file_minekube_connect_v1alpha1_watch_service_proto_rawDescGZIP() []byte {
	file_minekube_connect_v1alpha1_watch_service_proto_rawDescOnce.Do(func() {
		file_minekube_connect_v1alpha1_watch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_minekube_connect_v1alpha1_watch_service_proto_rawDescData)
	})
	return file_minekube_connect_v1alpha1_watch_service_proto_rawDescData
}

var file_minekube_connect_v1alpha1_watch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_minekube_connect_v1alpha1_watch_service_proto_goTypes = []interface{}{
	(*WatchRequest)(nil),        // 0: minekube.connect.v1alpha1.WatchRequest
	(*WatchResponse)(nil),       // 1: minekube.connect.v1alpha1.WatchResponse
	(*Endpoint)(nil),            // 2: minekube.connect.v1alpha1.Endpoint
	(*SessionRejection)(nil),    // 3: minekube.connect.v1alpha1.SessionRejection
	(*Session)(nil),             // 4: minekube.connect.v1alpha1.Session
	(*Player)(nil),              // 5: minekube.connect.v1alpha1.Player
	(*GameProfile)(nil),         // 6: minekube.connect.v1alpha1.GameProfile
	(*GameProfileProperty)(nil), // 7: minekube.connect.v1alpha1.GameProfileProperty
	(*status.Status)(nil),       // 8: google.rpc.Status
}
var file_minekube_connect_v1alpha1_watch_service_proto_depIdxs = []int32{
	2, // 0: minekube.connect.v1alpha1.WatchRequest.endpoint:type_name -> minekube.connect.v1alpha1.Endpoint
	3, // 1: minekube.connect.v1alpha1.WatchRequest.session_rejection:type_name -> minekube.connect.v1alpha1.SessionRejection
	4, // 2: minekube.connect.v1alpha1.WatchResponse.session:type_name -> minekube.connect.v1alpha1.Session
	8, // 3: minekube.connect.v1alpha1.SessionRejection.reason:type_name -> google.rpc.Status
	5, // 4: minekube.connect.v1alpha1.Session.player:type_name -> minekube.connect.v1alpha1.Player
	6, // 5: minekube.connect.v1alpha1.Player.profile:type_name -> minekube.connect.v1alpha1.GameProfile
	7, // 6: minekube.connect.v1alpha1.GameProfile.properties:type_name -> minekube.connect.v1alpha1.GameProfileProperty
	0, // 7: minekube.connect.v1alpha1.WatchService.Watch:input_type -> minekube.connect.v1alpha1.WatchRequest
	1, // 8: minekube.connect.v1alpha1.WatchService.Watch:output_type -> minekube.connect.v1alpha1.WatchResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_minekube_connect_v1alpha1_watch_service_proto_init() }
func file_minekube_connect_v1alpha1_watch_service_proto_init() {
	if File_minekube_connect_v1alpha1_watch_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRejection); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameProfile); i {
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
		file_minekube_connect_v1alpha1_watch_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameProfileProperty); i {
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
			RawDescriptor: file_minekube_connect_v1alpha1_watch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minekube_connect_v1alpha1_watch_service_proto_goTypes,
		DependencyIndexes: file_minekube_connect_v1alpha1_watch_service_proto_depIdxs,
		MessageInfos:      file_minekube_connect_v1alpha1_watch_service_proto_msgTypes,
	}.Build()
	File_minekube_connect_v1alpha1_watch_service_proto = out.File
	file_minekube_connect_v1alpha1_watch_service_proto_rawDesc = nil
	file_minekube_connect_v1alpha1_watch_service_proto_goTypes = nil
	file_minekube_connect_v1alpha1_watch_service_proto_depIdxs = nil
}
