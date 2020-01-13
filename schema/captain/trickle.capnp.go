// Code generated by capnpc-go. DO NOT EDIT.

package captain

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Auth struct{ capnp.Struct }

// Auth_TypeID is the unique identifier for the type Auth.
const Auth_TypeID = 0xd8cdf7e707cbaf27

func NewAuth(s *capnp.Segment) (Auth, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Auth{st}, err
}

func NewRootAuth(s *capnp.Segment) (Auth, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Auth{st}, err
}

func ReadRootAuth(msg *capnp.Message) (Auth, error) {
	root, err := msg.RootPtr()
	return Auth{root.Struct()}, err
}

func (s Auth) String() string {
	str, _ := text.Marshal(0xd8cdf7e707cbaf27, s.Struct)
	return str
}

func (s Auth) NodeId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Auth) HasNodeId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Auth) SetNodeId(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Auth_List is a list of Auth.
type Auth_List struct{ capnp.List }

// NewAuth creates a new list of Auth.
func NewAuth_List(s *capnp.Segment, sz int32) (Auth_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Auth_List{l}, err
}

func (s Auth_List) At(i int) Auth { return Auth{s.List.Struct(i)} }

func (s Auth_List) Set(i int, v Auth) error { return s.List.SetStruct(i, v.Struct) }

func (s Auth_List) String() string {
	str, _ := text.MarshalList(0xd8cdf7e707cbaf27, s.List)
	return str
}

// Auth_Promise is a wrapper for a Auth promised by a client call.
type Auth_Promise struct{ *capnp.Pipeline }

func (p Auth_Promise) Struct() (Auth, error) {
	s, err := p.Pipeline.Struct()
	return Auth{s}, err
}

type Ping struct{ capnp.Struct }

// Ping_TypeID is the unique identifier for the type Ping.
const Ping_TypeID = 0x8b3b5f400419ca82

func NewPing(s *capnp.Segment) (Ping, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ping{st}, err
}

func NewRootPing(s *capnp.Segment) (Ping, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ping{st}, err
}

func ReadRootPing(msg *capnp.Message) (Ping, error) {
	root, err := msg.RootPtr()
	return Ping{root.Struct()}, err
}

func (s Ping) String() string {
	str, _ := text.Marshal(0x8b3b5f400419ca82, s.Struct)
	return str
}

func (s Ping) Nonce() uint32 {
	return s.Struct.Uint32(0)
}

func (s Ping) SetNonce(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Ping_List is a list of Ping.
type Ping_List struct{ capnp.List }

// NewPing creates a new list of Ping.
func NewPing_List(s *capnp.Segment, sz int32) (Ping_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Ping_List{l}, err
}

func (s Ping_List) At(i int) Ping { return Ping{s.List.Struct(i)} }

func (s Ping_List) Set(i int, v Ping) error { return s.List.SetStruct(i, v.Struct) }

func (s Ping_List) String() string {
	str, _ := text.MarshalList(0x8b3b5f400419ca82, s.List)
	return str
}

// Ping_Promise is a wrapper for a Ping promised by a client call.
type Ping_Promise struct{ *capnp.Pipeline }

func (p Ping_Promise) Struct() (Ping, error) {
	s, err := p.Pipeline.Struct()
	return Ping{s}, err
}

type Pong struct{ capnp.Struct }

// Pong_TypeID is the unique identifier for the type Pong.
const Pong_TypeID = 0xc10601ff09f32611

func NewPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func NewRootPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func ReadRootPong(msg *capnp.Message) (Pong, error) {
	root, err := msg.RootPtr()
	return Pong{root.Struct()}, err
}

func (s Pong) String() string {
	str, _ := text.Marshal(0xc10601ff09f32611, s.Struct)
	return str
}

func (s Pong) Nonce() uint32 {
	return s.Struct.Uint32(0)
}

func (s Pong) SetNonce(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Pong_List is a list of Pong.
type Pong_List struct{ capnp.List }

// NewPong creates a new list of Pong.
func NewPong_List(s *capnp.Segment, sz int32) (Pong_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Pong_List{l}, err
}

func (s Pong_List) At(i int) Pong { return Pong{s.List.Struct(i)} }

func (s Pong_List) Set(i int, v Pong) error { return s.List.SetStruct(i, v.Struct) }

func (s Pong_List) String() string {
	str, _ := text.MarshalList(0xc10601ff09f32611, s.List)
	return str
}

// Pong_Promise is a wrapper for a Pong promised by a client call.
type Pong_Promise struct{ *capnp.Pipeline }

func (p Pong_Promise) Struct() (Pong, error) {
	s, err := p.Pipeline.Struct()
	return Pong{s}, err
}

type Announce struct{ capnp.Struct }

// Announce_TypeID is the unique identifier for the type Announce.
const Announce_TypeID = 0xda6b658bc18baf1d

func NewAnnounce(s *capnp.Segment) (Announce, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Announce{st}, err
}

func NewRootAnnounce(s *capnp.Segment) (Announce, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Announce{st}, err
}

func ReadRootAnnounce(msg *capnp.Message) (Announce, error) {
	root, err := msg.RootPtr()
	return Announce{root.Struct()}, err
}

func (s Announce) String() string {
	str, _ := text.Marshal(0xda6b658bc18baf1d, s.Struct)
	return str
}

func (s Announce) ChannelId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Announce) SetChannelId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Announce) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Announce) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Announce) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Announce_List is a list of Announce.
type Announce_List struct{ capnp.List }

// NewAnnounce creates a new list of Announce.
func NewAnnounce_List(s *capnp.Segment, sz int32) (Announce_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Announce_List{l}, err
}

func (s Announce_List) At(i int) Announce { return Announce{s.List.Struct(i)} }

func (s Announce_List) Set(i int, v Announce) error { return s.List.SetStruct(i, v.Struct) }

func (s Announce_List) String() string {
	str, _ := text.MarshalList(0xda6b658bc18baf1d, s.List)
	return str
}

// Announce_Promise is a wrapper for a Announce promised by a client call.
type Announce_Promise struct{ *capnp.Pipeline }

func (p Announce_Promise) Struct() (Announce, error) {
	s, err := p.Pipeline.Struct()
	return Announce{s}, err
}

type Request struct{ capnp.Struct }

// Request_TypeID is the unique identifier for the type Request.
const Request_TypeID = 0xe526bff04d4ce1e9

func NewRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Request{st}, err
}

func NewRootRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Request{st}, err
}

func ReadRootRequest(msg *capnp.Message) (Request, error) {
	root, err := msg.RootPtr()
	return Request{root.Struct()}, err
}

func (s Request) String() string {
	str, _ := text.Marshal(0xe526bff04d4ce1e9, s.Struct)
	return str
}

func (s Request) ChannelId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Request) SetChannelId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Request) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Request) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Request) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Request_List is a list of Request.
type Request_List struct{ capnp.List }

// NewRequest creates a new list of Request.
func NewRequest_List(s *capnp.Segment, sz int32) (Request_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Request_List{l}, err
}

func (s Request_List) At(i int) Request { return Request{s.List.Struct(i)} }

func (s Request_List) Set(i int, v Request) error { return s.List.SetStruct(i, v.Struct) }

func (s Request_List) String() string {
	str, _ := text.MarshalList(0xe526bff04d4ce1e9, s.List)
	return str
}

// Request_Promise is a wrapper for a Request promised by a client call.
type Request_Promise struct{ *capnp.Pipeline }

func (p Request_Promise) Struct() (Request, error) {
	s, err := p.Pipeline.Struct()
	return Request{s}, err
}

type Response struct{ capnp.Struct }

// Response_TypeID is the unique identifier for the type Response.
const Response_TypeID = 0xe06f2535c7dfe01b

func NewResponse(s *capnp.Segment) (Response, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Response{st}, err
}

func NewRootResponse(s *capnp.Segment) (Response, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Response{st}, err
}

func ReadRootResponse(msg *capnp.Message) (Response, error) {
	root, err := msg.RootPtr()
	return Response{root.Struct()}, err
}

func (s Response) String() string {
	str, _ := text.Marshal(0xe06f2535c7dfe01b, s.Struct)
	return str
}

func (s Response) ChannelId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Response) SetChannelId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Response) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Response) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Response) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Response) OriginId() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Response) HasOriginId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Response) SetOriginId(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Response) TargetIds() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.DataList{List: p.List()}, err
}

func (s Response) HasTargetIds() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Response) SetTargetIds(v capnp.DataList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewTargetIds sets the targetIds field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Response) NewTargetIds(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Response) Payload() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Response) HasPayload() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Response) SetPayload(v []byte) error {
	return s.Struct.SetData(3, v)
}

// Response_List is a list of Response.
type Response_List struct{ capnp.List }

// NewResponse creates a new list of Response.
func NewResponse_List(s *capnp.Segment, sz int32) (Response_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return Response_List{l}, err
}

func (s Response_List) At(i int) Response { return Response{s.List.Struct(i)} }

func (s Response_List) Set(i int, v Response) error { return s.List.SetStruct(i, v.Struct) }

func (s Response_List) String() string {
	str, _ := text.MarshalList(0xe06f2535c7dfe01b, s.List)
	return str
}

// Response_Promise is a wrapper for a Response promised by a client call.
type Response_Promise struct{ *capnp.Pipeline }

func (p Response_Promise) Struct() (Response, error) {
	s, err := p.Pipeline.Struct()
	return Response{s}, err
}

const schema_f71cc0af2f870b3c = "x\xda\xac\x90\xcfk\x13A\x18\x86\xbfwf6I\x91" +
	"T\xb6\xbb\x1e\x8a\x86\x9c\xac?\x10c\xab\x05\xa9B[" +
	"\xc1CD!\xd3\x937Y6C\x12\x1agcvc" +
	"\x11\x94\x8a V\xe2AA\xf0,x\x11\x94\x98?@" +
	"4'A\xf1\xaeB\xb1\x1e\x8a\xe4 \x08\x1ez\\\x99" +
	"\x15R\x09\x85\xe6\xd0\xdb\xf0\xcd\xf3>\xef7sj\x1d" +
	"\x0bl\xdaz\x96&\x92W\xadT|\xef\xe3\xa4X\xb8" +
	"v\xaeM2\x07\xc4\xe7\xf7=(t\xde\x1f\xda\"\x91" +
	"&:\xdd\xe5\x13pz<M\xe4\xbc\xe5\xf3\x84\xd8\x9e" +
	"\xfa3\x16#\xd5\xdb\x09^7\xf0\xaf\x04\xee'\xf0\x91" +
	"\xce\xa7\xf4\xcf\xad\xcf_\xc8\xce\xfd\xc7Z0\xf0\x98\x98" +
	"\x803ir\xce\x01a\xe0\\\xa7\xddk\xab\xe5oC" +
	"\xe6\x7f\xf4\xac\x98\x81s1\xa1\x17\xc5\x0a!>\xb8\xf1" +
	"\xfd\xc3\xec\xe1`c\x98N\x16yn\xe8nr|-" +
	"\xf2 \xc4\xfd\x1f\x97\xaf\xfc~7\xb5\xb9\xa3\xfc\xabu" +
	"\x1cN\xdf2\xf2Mk\x85\xbaq\xe8W\xd5u\xaf\xe0" +
	"s\xaf\x11y5]\x88\x9a5\x7f\xb9\xaeN\xfa^C" +
	"7\xe6J5\x8dJ\x09\x90\x82\x0b\"\x01\";;C" +
	"$3\x1c\xd2e\xc8\xeb@\xfb\x0a\x19b\xc8\x10vs" +
	"\x05{\xe7ZlE\xa8\x0e\xb9\xe6\xb6]\xf3:(\xab" +
	"b\x19Yb\xc8\xee.\xd3:\x1f\xb4\xb4\xaf\x8c03" +
	"\x10\x1e[\"\x92G9\xe4\x19\x06\xc0\x85\x99M_ " +
	"\x92'8\xe4Y\x86\xd8\xafzZ\xabz\x91PF\x8a" +
	"\x18R\x84UuS\xe9h\xe4\xe6%\x15\xe6\x1b\x81\x0e" +
	"\x93fw\xd0|\xc74\xdf\xe6\x90k\xdb\xcd\xf7M\xf3" +
	"]\x0e\xf9\x88\xc1fp\xc1\x88\xec\x87\x97\x88\xe4\x1a\x87" +
	"|\xca`s\xe6\x82\x13\xd9OL\xfa1\x87|\xc5`" +
	"\x0b\xeeB\x10\xd9/M\xfc\x05\x87|3\xe2\xe2A\xb3" +
	"V\xa9\xe9b\x99\x88\x06\xb3\xc8kVTT,\x13B" +
	"\x8c\x13J\x1c\xc9\xd58a\xb5\xe1\xdd\xaa\x07\xde\xe8\xef" +
	"\xbe\xb1\xbf\xa5\xc2h\xef?\xfco\x00\x00\x00\xff\xffR" +
	"\xe8\xef\x1a"

func init() {
	schemas.Register(schema_f71cc0af2f870b3c,
		0x8b3b5f400419ca82,
		0xc10601ff09f32611,
		0xd8cdf7e707cbaf27,
		0xda6b658bc18baf1d,
		0xe06f2535c7dfe01b,
		0xe526bff04d4ce1e9)
}
