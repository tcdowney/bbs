// Code generated by protoc-gen-gogo.
// source: check_definition.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CheckDefinition struct {
	Checks []*Check `protobuf:"bytes,1,rep,name=checks" json:"checks,omitempty"`
}

func (m *CheckDefinition) Reset()                    { *m = CheckDefinition{} }
func (*CheckDefinition) ProtoMessage()               {}
func (*CheckDefinition) Descriptor() ([]byte, []int) { return fileDescriptorCheckDefinition, []int{0} }

func (m *CheckDefinition) GetChecks() []*Check {
	if m != nil {
		return m.Checks
	}
	return nil
}

type Check struct {
	// oneof is hard to use right now, instead we can do this check in validation
	// oneof check {
	TcpCheck  *TCPCheck  `protobuf:"bytes,1,opt,name=tcp_check,json=tcpCheck" json:"tcp_check,omitempty"`
	HttpCheck *HTTPCheck `protobuf:"bytes,2,opt,name=http_check,json=httpCheck" json:"http_check,omitempty"`
}

func (m *Check) Reset()                    { *m = Check{} }
func (*Check) ProtoMessage()               {}
func (*Check) Descriptor() ([]byte, []int) { return fileDescriptorCheckDefinition, []int{1} }

func (m *Check) GetTcpCheck() *TCPCheck {
	if m != nil {
		return m.TcpCheck
	}
	return nil
}

func (m *Check) GetHttpCheck() *HTTPCheck {
	if m != nil {
		return m.HttpCheck
	}
	return nil
}

type TCPCheck struct {
	Port             uint32 `protobuf:"varint,1,opt,name=port" json:"port"`
	ConnectTimeoutMs uint64 `protobuf:"varint,2,opt,name=connect_timeout_ms,json=connectTimeoutMs" json:"connect_timeout_ms,omitempty"`
}

func (m *TCPCheck) Reset()                    { *m = TCPCheck{} }
func (*TCPCheck) ProtoMessage()               {}
func (*TCPCheck) Descriptor() ([]byte, []int) { return fileDescriptorCheckDefinition, []int{2} }

func (m *TCPCheck) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TCPCheck) GetConnectTimeoutMs() uint64 {
	if m != nil {
		return m.ConnectTimeoutMs
	}
	return 0
}

type HTTPCheck struct {
	Port             uint32 `protobuf:"varint,1,opt,name=port" json:"port"`
	RequestTimeoutMs uint64 `protobuf:"varint,2,opt,name=request_timeout_ms,json=requestTimeoutMs" json:"request_timeout_ms,omitempty"`
	Path             string `protobuf:"bytes,3,opt,name=path" json:"path"`
}

func (m *HTTPCheck) Reset()                    { *m = HTTPCheck{} }
func (*HTTPCheck) ProtoMessage()               {}
func (*HTTPCheck) Descriptor() ([]byte, []int) { return fileDescriptorCheckDefinition, []int{3} }

func (m *HTTPCheck) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HTTPCheck) GetRequestTimeoutMs() uint64 {
	if m != nil {
		return m.RequestTimeoutMs
	}
	return 0
}

func (m *HTTPCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckDefinition)(nil), "models.CheckDefinition")
	proto.RegisterType((*Check)(nil), "models.Check")
	proto.RegisterType((*TCPCheck)(nil), "models.TCPCheck")
	proto.RegisterType((*HTTPCheck)(nil), "models.HTTPCheck")
}
func (this *CheckDefinition) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CheckDefinition)
	if !ok {
		that2, ok := that.(CheckDefinition)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Checks) != len(that1.Checks) {
		return false
	}
	for i := range this.Checks {
		if !this.Checks[i].Equal(that1.Checks[i]) {
			return false
		}
	}
	return true
}
func (this *Check) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Check)
	if !ok {
		that2, ok := that.(Check)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.TcpCheck.Equal(that1.TcpCheck) {
		return false
	}
	if !this.HttpCheck.Equal(that1.HttpCheck) {
		return false
	}
	return true
}
func (this *TCPCheck) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TCPCheck)
	if !ok {
		that2, ok := that.(TCPCheck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.ConnectTimeoutMs != that1.ConnectTimeoutMs {
		return false
	}
	return true
}
func (this *HTTPCheck) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HTTPCheck)
	if !ok {
		that2, ok := that.(HTTPCheck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.RequestTimeoutMs != that1.RequestTimeoutMs {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	return true
}
func (this *CheckDefinition) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&models.CheckDefinition{")
	if this.Checks != nil {
		s = append(s, "Checks: "+fmt.Sprintf("%#v", this.Checks)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Check) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.Check{")
	if this.TcpCheck != nil {
		s = append(s, "TcpCheck: "+fmt.Sprintf("%#v", this.TcpCheck)+",\n")
	}
	if this.HttpCheck != nil {
		s = append(s, "HttpCheck: "+fmt.Sprintf("%#v", this.HttpCheck)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TCPCheck) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.TCPCheck{")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "ConnectTimeoutMs: "+fmt.Sprintf("%#v", this.ConnectTimeoutMs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HTTPCheck) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&models.HTTPCheck{")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "RequestTimeoutMs: "+fmt.Sprintf("%#v", this.RequestTimeoutMs)+",\n")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCheckDefinition(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CheckDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckDefinition) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Checks) > 0 {
		for _, msg := range m.Checks {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCheckDefinition(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Check) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Check) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TcpCheck != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCheckDefinition(dAtA, i, uint64(m.TcpCheck.Size()))
		n1, err := m.TcpCheck.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.HttpCheck != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCheckDefinition(dAtA, i, uint64(m.HttpCheck.Size()))
		n2, err := m.HttpCheck.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *TCPCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TCPCheck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintCheckDefinition(dAtA, i, uint64(m.Port))
	dAtA[i] = 0x10
	i++
	i = encodeVarintCheckDefinition(dAtA, i, uint64(m.ConnectTimeoutMs))
	return i, nil
}

func (m *HTTPCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPCheck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintCheckDefinition(dAtA, i, uint64(m.Port))
	dAtA[i] = 0x10
	i++
	i = encodeVarintCheckDefinition(dAtA, i, uint64(m.RequestTimeoutMs))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintCheckDefinition(dAtA, i, uint64(len(m.Path)))
	i += copy(dAtA[i:], m.Path)
	return i, nil
}

func encodeFixed64CheckDefinition(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32CheckDefinition(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCheckDefinition(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CheckDefinition) Size() (n int) {
	var l int
	_ = l
	if len(m.Checks) > 0 {
		for _, e := range m.Checks {
			l = e.Size()
			n += 1 + l + sovCheckDefinition(uint64(l))
		}
	}
	return n
}

func (m *Check) Size() (n int) {
	var l int
	_ = l
	if m.TcpCheck != nil {
		l = m.TcpCheck.Size()
		n += 1 + l + sovCheckDefinition(uint64(l))
	}
	if m.HttpCheck != nil {
		l = m.HttpCheck.Size()
		n += 1 + l + sovCheckDefinition(uint64(l))
	}
	return n
}

func (m *TCPCheck) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovCheckDefinition(uint64(m.Port))
	n += 1 + sovCheckDefinition(uint64(m.ConnectTimeoutMs))
	return n
}

func (m *HTTPCheck) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovCheckDefinition(uint64(m.Port))
	n += 1 + sovCheckDefinition(uint64(m.RequestTimeoutMs))
	l = len(m.Path)
	n += 1 + l + sovCheckDefinition(uint64(l))
	return n
}

func sovCheckDefinition(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCheckDefinition(x uint64) (n int) {
	return sovCheckDefinition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CheckDefinition) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CheckDefinition{`,
		`Checks:` + strings.Replace(fmt.Sprintf("%v", this.Checks), "Check", "Check", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Check) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Check{`,
		`TcpCheck:` + strings.Replace(fmt.Sprintf("%v", this.TcpCheck), "TCPCheck", "TCPCheck", 1) + `,`,
		`HttpCheck:` + strings.Replace(fmt.Sprintf("%v", this.HttpCheck), "HTTPCheck", "HTTPCheck", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TCPCheck) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TCPCheck{`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`ConnectTimeoutMs:` + fmt.Sprintf("%v", this.ConnectTimeoutMs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPCheck) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPCheck{`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`RequestTimeoutMs:` + fmt.Sprintf("%v", this.RequestTimeoutMs) + `,`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCheckDefinition(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CheckDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckDefinition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checks = append(m.Checks, &Check{})
			if err := m.Checks[len(m.Checks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Check) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckDefinition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Check: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Check: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpCheck", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TcpCheck == nil {
				m.TcpCheck = &TCPCheck{}
			}
			if err := m.TcpCheck.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpCheck", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpCheck == nil {
				m.HttpCheck = &HTTPCheck{}
			}
			if err := m.HttpCheck.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TCPCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckDefinition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TCPCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectTimeoutMs", wireType)
			}
			m.ConnectTimeoutMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectTimeoutMs |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCheckDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HTTPCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckDefinition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeoutMs", wireType)
			}
			m.RequestTimeoutMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestTimeoutMs |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckDefinition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCheckDefinition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheckDefinition
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCheckDefinition
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCheckDefinition
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCheckDefinition
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCheckDefinition(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCheckDefinition = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheckDefinition   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("check_definition.proto", fileDescriptorCheckDefinition) }

var fileDescriptorCheckDefinition = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0x3f, 0x4b, 0xc3, 0x40,
	0x1c, 0x86, 0x73, 0xb6, 0x96, 0xe6, 0x4a, 0xb1, 0x46, 0xd1, 0x20, 0x72, 0x29, 0x41, 0xa1, 0x43,
	0x4d, 0xa1, 0x93, 0x73, 0x2a, 0x22, 0x82, 0x20, 0x21, 0x7b, 0x68, 0xd3, 0x6b, 0x13, 0x34, 0xb9,
	0xd8, 0x5c, 0x40, 0x37, 0x3f, 0x82, 0xe0, 0xe8, 0x17, 0xf0, 0xa3, 0x74, 0xec, 0xe8, 0x14, 0x6c,
	0x5c, 0xa4, 0x53, 0x3f, 0x82, 0xf4, 0x97, 0x3f, 0x16, 0x0a, 0xdd, 0xee, 0xde, 0xf7, 0xbd, 0xe7,
	0x39, 0x7c, 0x64, 0x3b, 0xd4, 0x7e, 0xb0, 0x86, 0x74, 0xe4, 0xfa, 0x2e, 0x77, 0x99, 0xaf, 0x05,
	0x13, 0xc6, 0x99, 0x54, 0xf1, 0xd8, 0x90, 0x3e, 0x86, 0x27, 0x17, 0x63, 0x97, 0x3b, 0xd1, 0x40,
	0xb3, 0x99, 0xd7, 0x19, 0xb3, 0x31, 0xeb, 0x40, 0x3d, 0x88, 0x46, 0x70, 0x83, 0x0b, 0x9c, 0xd2,
	0x67, 0xea, 0x25, 0xde, 0xeb, 0xad, 0x80, 0x57, 0x05, 0x4f, 0x3a, 0xc7, 0x15, 0x70, 0x84, 0x32,
	0x6a, 0x96, 0x5a, 0xb5, 0x6e, 0x5d, 0x4b, 0xd1, 0x1a, 0x0c, 0x8d, 0xac, 0x54, 0x3f, 0x10, 0xde,
	0x85, 0x44, 0xba, 0xc6, 0x22, 0xb7, 0x03, 0x0b, 0x72, 0x19, 0x35, 0x51, 0xab, 0xd6, 0x6d, 0xe4,
	0x6f, 0xcc, 0xde, 0x3d, 0x8c, 0xf4, 0xe3, 0x45, 0xac, 0x1c, 0x14, 0xb3, 0x36, 0xf3, 0x5c, 0x4e,
	0xbd, 0x80, 0xbf, 0x18, 0x55, 0x6e, 0x07, 0x29, 0xe7, 0x16, 0x63, 0x87, 0xf3, 0x1c, 0xb4, 0x03,
	0xa0, 0xfd, 0x1c, 0x74, 0x63, 0x9a, 0x19, 0x49, 0x5e, 0xc4, 0xca, 0xe1, 0xff, 0x70, 0x0d, 0x25,
	0xae, 0x52, 0x18, 0xa9, 0xcf, 0xb8, 0x9a, 0xab, 0x25, 0x19, 0x97, 0x03, 0x36, 0xe1, 0xf0, 0xb5,
	0xba, 0x5e, 0x9e, 0xc6, 0x8a, 0x60, 0x40, 0x22, 0x19, 0x58, 0xb2, 0x99, 0xef, 0x53, 0x9b, 0x5b,
	0xdc, 0xf5, 0x28, 0x8b, 0xb8, 0xe5, 0x85, 0x60, 0x2e, 0xeb, 0x67, 0xab, 0xdd, 0x22, 0x56, 0x4e,
	0x37, 0x17, 0x6b, 0xca, 0x46, 0xd6, 0x9a, 0x69, 0x79, 0x17, 0xaa, 0xef, 0x08, 0x8b, 0xc5, 0x67,
	0xb7, 0xbb, 0x27, 0xf4, 0x29, 0xa2, 0xe1, 0x36, 0xf7, 0xe6, 0x62, 0xdd, 0x9d, 0xb5, 0x85, 0x1b,
	0x6c, 0x7d, 0xee, 0xc8, 0xa5, 0x26, 0x6a, 0x89, 0x85, 0xad, 0xcf, 0x1d, 0xbd, 0x3d, 0x9b, 0x13,
	0xe1, 0x6b, 0x4e, 0x84, 0xe5, 0x9c, 0xa0, 0xd7, 0x84, 0xa0, 0xcf, 0x84, 0xa0, 0x69, 0x42, 0xd0,
	0x2c, 0x21, 0xe8, 0x3b, 0x21, 0xe8, 0x37, 0x21, 0xc2, 0x32, 0x21, 0xe8, 0xed, 0x87, 0x08, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xef, 0x53, 0x4d, 0x65, 0x02, 0x00, 0x00,
}
