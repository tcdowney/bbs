package models

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/gogo/protobuf/proto"
)

type SerializationFormat byte

const JSON SerializationFormat = 0
const PROTO SerializationFormat = 1

type Version byte

const V0 Version = 0

type Envelope struct {
	SerializationFormat SerializationFormat
	Version             Version
	Payload             []byte
}

func Open(data []byte) *Envelope {
	e := &Envelope{}
	if !isEncoded(data) {
		e.SerializationFormat = JSON
		e.Version = V0
		e.Payload = data
		return e
	}

	e.SerializationFormat = SerializationFormat(data[0])
	e.Version = Version(V0)
	e.Payload = data[2:]
	return e
}

func isEncoded(data []byte) bool {
	if len(data) < 2 {
		return false
	}

	switch SerializationFormat(data[0]) {
	case JSON, PROTO:
	default:
		return false
	}

	switch Version(data[1]) {
	case V0:
	default:
		return false
	}

	return true
}

func Marshal(version Version, v ProtoValidator) ([]byte, *Error) {
	payload, err := ToProto(v)
	if err != nil {
		return nil, err
	}
	return append([]byte{byte(PROTO), byte(version)}, payload...), nil
}

func (e *Envelope) Unmarshal(model Validator) error {
	switch e.SerializationFormat {
	case JSON:
		return FromJSON(e.Payload, model)
	case PROTO:
		if pv, ok := model.(ProtoValidator); ok {
			return FromProto(e.Payload, pv)
		}
	}
	return errors.New("Unexpected serialization format")
}

func FromJSON(payload []byte, v Validator) error {
	err := json.Unmarshal(payload, v)
	if err != nil {
		return err
	}
	return v.Validate()
}

func ToJSON(v Validator) ([]byte, *Error) {
	if !isNil(v) {
		if err := v.Validate(); err != nil {
			return nil, NewError(InvalidRecord, err.Error())
		}
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, NewError(InvalidJSON, err.Error())
	}

	return bytes, nil
}

func FromProto(payload []byte, v ProtoValidator) error {
	err := proto.Unmarshal(payload, v)
	if err != nil {
		return err
	}
	return v.Validate()
}

func ToProto(v ProtoValidator) ([]byte, *Error) {
	if !isNil(v) {
		if err := v.Validate(); err != nil {
			return nil, NewError(InvalidRecord, err.Error())
		}
	}

	bytes, err := proto.Marshal(v)
	if err != nil {
		return nil, NewError(InvalidProtobufMessage, err.Error())
	}

	return bytes, nil
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}

	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}

	return false
}
