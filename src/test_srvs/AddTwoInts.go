package test_srvs

import (
    "bytes"
    "encoding/binary"
    "ros"
)

// Message type metadata
type type_AddTwoIntsRequest struct {
    definition string
    name       string
    md5sum     string
}

func (t *type_AddTwoIntsRequest) Definition() string      { return t.definition }
func (t *type_AddTwoIntsRequest) Name() string            { return t.name }
func (t *type_AddTwoIntsRequest) MD5Sum() string          { return t.md5sum }
func (t *type_AddTwoIntsRequest) NewMessage() ros.Message {
    return new(AddTwoIntsRequest)
}

var (
    TypeOfAddTwoIntsRequest = &type_AddTwoIntsRequest{
        "",
        "test_msgs/AddTwoIntsRequest",
        "992ce8a1687cec8c8bd883ec73ca41d1",
    }
)

type AddTwoIntsRequest struct {
    A int32
    B int32
}

func (m *AddTwoIntsRequest) Serialize() []byte {
    var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, m.A)
    binary.Write(&buf, binary.LittleEndian, m.B)
    return buf.Bytes()
}

func (m *AddTwoIntsRequest) Deserialize(buffer []byte) error {
    buf := bytes.NewBuffer(buffer)
    if err := binary.Read(buf, binary.LittleEndian, &m.A); err != nil {
        return err
    }
    if err := binary.Read(buf, binary.LittleEndian, &m.B); err != nil {
        return err
    }
    return nil
}


// Message type metadata
type type_AddTwoIntsResponse struct {
    definition string
    name       string
    md5sum     string
}

func (t *type_AddTwoIntsResponse) Definition() string      { return t.definition }
func (t *type_AddTwoIntsResponse) Name() string            { return t.name }
func (t *type_AddTwoIntsResponse) MD5Sum() string          { return t.md5sum }
func (t *type_AddTwoIntsResponse) NewMessage() ros.Message {
    return new(AddTwoIntsRequest)
}

var (
    TypeOfAddTwoIntsResponse = &type_AddTwoIntsResponse{
        "",
        "test_msgs/AddTwoIntsResponse",
        "992ce8a1687cec8c8bd883ec73ca41d1",
    }
)

type AddTwoIntsResponse struct {
    Result int32
}

func (m *AddTwoIntsResponse) Serialize() []byte {
    var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, m.Result)
    return buf.Bytes()
}

func (m *AddTwoIntsResponse) Deserialize(buffer []byte) error {
    buf := bytes.NewBuffer(buffer)
    if err := binary.Read(buf, binary.LittleEndian, &m.Result); err != nil {
        return err
    }
    return nil
}

// Service type metadata
type type_AddTwoInts struct {
    name string
    md5sum string
    reqType ros.MessageType
    resType ros.MessageType
}

func (t *type_AddTwoInts) Name() string { return t.name }
func (t *type_AddTwoInts) MD5Sum() string { return t.md5sum }
func (t *type_AddTwoInts) RequestType() ros.MessageType { return t.reqType }
func (t *type_AddTwoInts) ResponseType() ros.MessageType { return t.resType }

var (
    TypeOfAddTwoInts = &type_AddTwoInts {
        "test_srvs",
        "hoge",
        TypeOfAddTwoIntsRequest,
        TypeOfAddTwoIntsResponse,
    }
)


type AddTwoInts struct {
    Req AddTwoIntsRequest
    Res AddTwoIntsResponse
}

func (s *AddTwoInts) Request() ros.Message { return &s.Req }
func (s *AddTwoInts) Response() ros.Message { return &s.Res }

